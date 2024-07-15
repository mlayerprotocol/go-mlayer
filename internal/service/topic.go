package service

import (
	"context"
	"encoding/hex"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

/*
Validate an agent authorization
*/
func ValidateTopicData(topic *entities.Topic, authState *models.AuthorizationState) (currentTopicState *models.TopicState, err error) {
	subnet := models.SubnetState{}

	// TODO state might have changed befor receiving event, so we need to find state that is relevant to this event.
	err = query.GetOne(models.SubnetState{Subnet: entities.Subnet{ID: topic.Subnet}}, &subnet)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperror.Forbidden("Invalid subnet id")
		}
		return nil, apperror.Internal(err.Error())
	}
	if  *authState.Priviledge < constants.StandardPriviledge {
		return nil, apperror.Forbidden("Agent does not have enough permission to create topics")
	}
	
	if len(topic.Ref) > 40 {
		return nil, apperror.BadRequest("Topic handle can not be more than 40 characters")
	}
	if !utils.IsAlphaNumericDot(topic.Ref) {
		return nil, apperror.BadRequest("Handle must be alphanumeric, _ and . but cannot start with a number")
	}
	return nil, nil
}

func HandleNewPubSubTopicEvent(event *entities.Event, ctx *context.Context) {
	logger.WithFields(logrus.Fields{"event": event}).Debug("New topic event from pubsub channel")
	markAsSynced := false
	updateState := false
	var eventError string
	// hash, _ := event.GetHash()

	logger.Infof("Event is a valid event %s", event.PayloadHash)
	cfg, _ := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)

	// Extract and validate the Data of the paylaod which is an Events Payload Data,
	data := event.Payload.Data.(*entities.Topic)
	logger.Infof("NEWTOPICEVENT: %s", event.Hash)

	err := ValidateEvent(*event)

	if err != nil {
		logger.Error(err)
		return
	}
	hash, _ := data.GetHash()
	data.Hash = hex.EncodeToString(hash)
	authEventHash := event.AuthEventHash
	authState, authError := query.GetOneAuthorizationState(entities.Authorization{Event: authEventHash})

	currentState, err := ValidateTopicData(data, authState)
	if err != nil {
		// penalize node for broadcasting invalid data
		logger.Infof("Invalid topic data %v. Node should be penalized", err)
		return
	}

	// check if we are upto date on this event

	prevEventUpToDate := query.EventExist(&event.PreviousEventHash) || (currentState == nil && event.PreviousEventHash.Hash == "") || (currentState != nil && currentState.Event.Hash == event.PreviousEventHash.Hash)
	authEventUpToDate := query.EventExist(&event.AuthEventHash) || (authState == nil && event.AuthEventHash.Hash == "") || (authState != nil && authState.Event == authEventHash)

	// Confirm if this is an older event coming after a newer event.
	// If it is, then we only have to update our event history, else we need to also update our current state
	isMoreRecent := false
	if currentState != nil && currentState.Hash != data.Hash {
		var currentStateEvent = &models.TopicEvent{}
		query.GetOne(entities.Event{Hash: currentState.Event.Hash}, currentStateEvent)
		isMoreRecent, markAsSynced = IsMoreRecent(
			currentStateEvent.ID,
			currentState.Event.Hash,
			currentStateEvent.Payload.Timestamp,
			event.Hash,
			event.Payload.Timestamp,
			markAsSynced,
		)
		// if uint64(currentStateEvent.Payload.Timestamp) < uint64(event.Payload.Timestamp) {
		// 	isMoreRecent = true
		// }
		// if uint64(currentStateEvent.Payload.Timestamp) > uint64(event.Payload.Timestamp) {
		// 	isMoreRecent = false
		// }
		// // if the authorization was created at exactly the same time but their hash is different
		// // use the last 4 digits of their event hash
		// if uint64(currentStateEvent.Payload.Timestamp) == uint64(event.Payload.Timestamp) {
		// 	// get the event payload of the current state

		// 	if err != nil && err != gorm.ErrRecordNotFound {
		// 		logger.Error("DB error", err)
		// 	}
		// 	if currentStateEvent.ID == "" {
		// 		markAsSynced = false
		// 	} else {
		// 		if currentStateEvent.Payload.Timestamp < event.Payload.Timestamp {
		// 			isMoreRecent = true
		// 		}
		// 		if currentStateEvent.Payload.Timestamp == event.Payload.Timestamp {
		// 			// logger.Infof("Current state %v", currentStateEvent.Payload)
		// 			csN := new(big.Int)
		// 			csN.SetString(currentState.Event.Hash[56:], 16)
		// 			nsN := new(big.Int)
		// 			nsN.SetString(event.Hash[56:], 16)

		// 			if csN.Cmp(nsN) < 1 {
		// 				isMoreRecent = true
		// 			}
		// 		}
		// 	}
		// }
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
				if strings.HasPrefix(authError.Error(), constants.ErrorForbidden) || strings.HasPrefix(authError.Error(), constants.ErrorUnauthorized) {
					markAsSynced = false
				} else {
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
			if currentState == nil || isMoreRecent {
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
	logger.Info(":::::updateState: Db Error", updateState, currentState == nil)

	// If the event was not signed by your node
	if string(event.Validator) != (*cfg).OperatorPublicKey  {
		// save the event
		event.Error = eventError
		event.IsValid = markAsSynced && len(eventError) == 0.
		event.Synced = markAsSynced
		event.Broadcasted = true
		_, _, err := query.SaveRecord(models.TopicEvent{
			Event: entities.Event{
				PayloadHash: event.PayloadHash,
			},
		}, models.TopicEvent{
			Event: *event,
		}, false, tx)
		if err != nil {
			tx.Rollback()
			logger.Error("1000: Db Error", err)
			return
		}
	} else {
		if markAsSynced {
			_, _, err := query.SaveRecord(models.TopicEvent{
				Event: entities.Event{PayloadHash: event.PayloadHash},
			}, models.TopicEvent{
				Event: entities.Event{Synced: true, Broadcasted: true, Error: eventError, IsValid: len(eventError) == 0},
			}, true, tx)
			if err != nil {
				logger.Error("DB error", err)
			}
		} else {
			// mark as broadcasted
			_, _, err := query.SaveRecord(models.TopicEvent{
				Event: entities.Event{PayloadHash: event.PayloadHash, Broadcasted: false},
			},
				models.TopicEvent{
					Event: entities.Event{Broadcasted: true},
				}, true, tx)
			if err != nil {
				logger.Error("DB error", err)
			}
		}
	}

	d, err := event.Payload.EncodeBytes()
	if err != nil {
		logger.Errorf("Invalid event payload")
	}
	agent, err := crypto.GetSignerECC(&d, &event.Payload.Signature)
	if err != nil {
		logger.Errorf("Invalid event payload")
	}
	data.Event = *entities.NewEventPath(event.Validator, entities.TopicEventModel, event.Hash)
	data.Agent = entities.AddressFromString(agent).ToDeviceString()
	data.Account = event.Payload.Account
	// logger.Error("data.Public ", data.Public)
	var newState *models.TopicState
	if updateState {
		newState, _, err = query.SaveRecord(models.TopicState{
			Topic: entities.Topic{ID: data.ID},
		}, models.TopicState{
			Topic: *data,
		}, event.EventType == uint16(constants.UpdateTopicEvent), tx)
		if err != nil {
			tx.Rollback()
			logger.Error("7000: Db Error", err)
			return
		}
	}
	tx.Commit()
	if markAsSynced {
		go OnFinishProcessingEvent(ctx, &data.Event, utils.IfThenElse(newState!=nil, &newState.ID, nil), utils.IfThenElse(event.Error!="", apperror.Internal(event.Error), nil))
	}

	if string(event.Validator) != (*cfg).OperatorPublicKey  {
		dependent, err := query.GetDependentEvents(*event)
		if err != nil {
			logger.Info("Unable to get dependent events", err)
		}
		for _, dep := range *dependent {
			go HandleNewPubSubTopicEvent(&dep, ctx)
		}
	}

	// TODO Broadcast the updated state
}
