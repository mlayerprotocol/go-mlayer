package service

import (
	"context"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/channelpool"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

/*
	Validate an agent authorization
*/
func ValidateTopicData(topic *entities.Topic) (prevTopicState *models.TopicState, grantorAuthState *models.AuthorizationState, err error) {
	// check fields of topic
	if len(topic.Handle) > 40 {
		return nil, nil, apperror.BadRequest("Topic handle cannont be more than 40 characters")
	}
	if !utils.IsAlphaNumericDot(topic.Handle) {
		return nil, nil, apperror.BadRequest("Handle must be alphanumeric, _ and . but cannot start with a number")
	}
	return nil, nil, nil
}




func HandleNewPubSubTopicEvent (event *entities.Event, ctx context.Context) {
		logger.WithFields(logrus.Fields{"event": event}).Debug("New auth event from pubsub channel")
		markAsSynced := false
		updateState := false
		var eventError string
		// hash, _ := event.GetHash()
		err := ValidateEvent(*event)
		
		if err != nil {
			logger.Error(err)
			return
		}

		logger.Infof("Event is a valid event %s",  event.PayloadHash)
		cfg, _ := ctx.Value(constants.ConfigKey).(*configs.MainConfiguration)

		// Extract and validate the Data of the paylaod which is an Events Payload Data,
		authRequest := event.Payload.Data.(*entities.Authorization)
		hash, _ := authRequest.GetHash()
		authRequest.Hash = hex.EncodeToString(hash)
		currentState, authState, authError := ValidateAuthData(authRequest)

		
	
		// check if we are upto date on this event
		prevEventUpToDate :=  (currentState == nil && event.PreviousEventHash == "") ||  (currentState != nil && currentState.EventHash == event.PreviousEventHash) 
		authEventUpToDate := (authState == nil && event.AuthEventHash == "") || (authState != nil && authState.EventHash == event.AuthEventHash) 
		
		// Confirm if this is an older event coming after a newer event.
		// If it is, then we only have to update our event history, else we need to also update our current state
		isMoreRecent := false
		if (currentState != nil && currentState.Hash != authRequest.Hash) {
			if currentState.Timestamp < authRequest.Timestamp {
				isMoreRecent = true
			}
			if currentState.Timestamp > authRequest.Timestamp {
				isMoreRecent = false
			}
			// if the authorization was created at exactly the same time but their hash is different
			// use the last 4 digits of their event hash
			if currentState.Timestamp == authRequest.Timestamp {
				// get the event payload of the current state
				 currentStateEvent, err := query.GetOneAuthorizationEvent( entities.Event{Hash: currentState.EventHash})
				if err != nil && err != gorm.ErrRecordNotFound {
					logger.Fatal("DB error", err)
				}
				if currentStateEvent == nil {
					markAsSynced = false
				} else {
					if currentStateEvent.Payload.Timestamp < event.Payload.Timestamp {
						isMoreRecent = true
					}
					if currentStateEvent.Payload.Timestamp == event.Payload.Timestamp {
						// logger.Infof("Current state %v", currentStateEvent.Payload)
						csN := new(big.Int)
						csN.SetString(currentState.EventHash[56:], 16)
						nsN := new(big.Int)
						nsN.SetString(event.Hash[56:], 16)
						
						if csN.Cmp(nsN) < 1 {
							isMoreRecent = true
						}	
					}
				}		
			}	
		}	
		if authError != nil {
			// check if we are upto date. If we are, then the error is an actual one
			// the error should be attached when saving the event	
			// But if we are not upto date, then we might need to wait for more info from the network
			
			if prevEventUpToDate && authEventUpToDate {
				// we are upto date. This is an actual error. No need to expect an update from the network
				eventError = authError.Error()
				markAsSynced = true
			} else {
				if currentState == nil || (currentState != nil && isMoreRecent) { // it is a morer ecent event
					if strings.HasPrefix(authError.Error() , constants.ErrorForbidden) || strings.HasPrefix(authError.Error() , constants.ErrorUnauthorized) {
						markAsSynced = false
					}  else {
						// entire event can be considered bad since the payload data is bad
						// this should have been sorted out before broadcasting to the network
						// TODO penalize the node that broadcasted this
						eventError = authError.Error()
						markAsSynced = true
					}
					
				} else { 
					// we are upto date. We just need to store this event as well.
					// No need to update state
					markAsSynced = true
					eventError = authError.Error()
				}
			}
				
		}
		
		// If no error, then we should act accordingly as well
		// If are upto date, then we should update the state based on if its a recent or old event
		if len(eventError) == 0 {
			if prevEventUpToDate && authEventUpToDate { // we are upto date
				if currentState == nil || (currentState != nil && isMoreRecent) {
					updateState = true
					markAsSynced = true
				} else {
					// Its an old event
					markAsSynced = true
					updateState = false
				}
			} else {
				updateState = false
				markAsSynced = false				
			}
			
		}
		
		
		// Save stuff permanently
		tx := sql.Db.Begin()

		// If the event was not signed by your node
		if string(event.Validator) != (*cfg).NetworkPublicKey {
			// save the event
			event.Error = eventError
			event.IsValid =  markAsSynced && len(eventError) == 0.
			event.Synced = markAsSynced
			event.Broadcasted = true
			_, _, err := query.SaveAuthorizationEvent(event, true, tx)
			if err != nil {
				tx.Rollback()
				logger.Fatal("5000: Db Error", err)
				return
			}	
		} else {
			if markAsSynced {
				_, err := query.UpdateAuthorizationEvent(entities.Event{PayloadHash: event.PayloadHash},
					entities.Event{Synced: true, Broadcasted: true, Error: eventError, IsValid: len(eventError) == 0}, tx)

					if err != nil {
						logger.Fatal("DB error", err)
					}
			} else {
				// mark as broadcasted
				_, err := query.UpdateAuthorizationEvent(entities.Event{PayloadHash: event.PayloadHash, Broadcasted: false},
					entities.Event{Broadcasted: true}, tx)
					if err != nil {
					logger.Fatal("DB error", err)
				}
			}
		}
		authRequest.EventHash = event.Hash
		if updateState {
			_, err := query.SaveAuthorizationState(authRequest, tx)
			if err != nil {
				tx.Rollback()
				logger.Fatal("5000: Db Error", err)
				return
			}
		}
		tx.Commit()
		
		if string(event.Validator) != (*cfg).NetworkPublicKey {
			dependent, err := query.GetDependentEvents(*event)
			if err != nil {
				logger.Info("Unable to get dependent events", err)
			}
			for _, dep := range *dependent {
				channelpool.AuthorizationEventPublishC <- &dep.Event
			}
		}

		// TODO Broadcast the updated state
	
}		

