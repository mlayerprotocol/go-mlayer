package service

import (
	"context"
	"encoding/hex"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
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
func ValidateSubscriptionData(subscription *entities.Subscription, payload *entities.ClientPayload) (currentSubscriptionState *models.SubscriptionState, err error) {
	// check fields of subscription
	var currentState *models.SubscriptionState

	err = query.GetOne(models.SubscriptionState{
		Subscription: entities.Subscription{Account: subscription.Account, Subnet: subscription.Subnet, Topic: subscription.Topic},
	}, &currentState)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			//return nil, nil, apperror.Unauthorized("Not a subscriber")
			// } else {
			return nil, err
		}
	}
	return currentState, err
}

func HandleNewPubSubSubscriptionEvent(event *entities.Event, ctx *context.Context) {
	logger.WithFields(logrus.Fields{"event": event}).Debug("New subscription event from pubsub channel")
	markAsSynced := false
	updateState := false
	var eventError string
	// hash, _ := event.GetHash()
	err := ValidateEvent(*event)

	if err != nil {
		logger.Error(err)
		return
	}
	d, err := event.Payload.EncodeBytes()
	if err != nil {
		logger.Errorf("Invalid event payload")
	}
	agent, err := crypto.GetSignerECC(&d, &event.Payload.Signature)
	if err != nil {
		logger.Errorf("Invalid event payload")
	}

	logger.Infof("Event is a valid event %s", event.PayloadHash)
	cfg, _ := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)

	// Extract and validate the Data of the paylaod which is an Events Payload Data,
	data := event.Payload.Data.(*entities.Subscription)
	hash, _ := data.GetHash()
	data.Hash = hex.EncodeToString(hash)

	var topicData *models.TopicState

	query.GetOne(models.TopicState{
		Topic: entities.Topic{ID: data.Topic, Subnet: event.Payload.Subnet},
	}, &topicData)

	data.Subnet = event.Payload.Subnet
	currentState, authError := ValidateSubscriptionData(data, &event.Payload)
	prevEventUpToDate := false
	authEventUpToDate := false

	// check if we are upto date on this event
	prevEventUpToDate = query.EventExist(&event.PreviousEventHash) || (currentState == nil && event.PreviousEventHash.Hash == "") || (currentState != nil && currentState.Event.Hash == event.PreviousEventHash.Hash)

	authState, _ := query.GetOneAuthorizationState(entities.Authorization{Event: event.AuthEventHash})

	authEventUpToDate = query.EventExist(&event.AuthEventHash) || (authState.ID == "" && event.AuthEventHash.Hash == "") || (authState.ID != "" && authState.Event.Hash == event.AuthEventHash.Hash)

	// Confirm if this is an older event coming after a newer event.
	// If it is, then we only have to update our event history, else we need to also update our current state
	isMoreRecent := false
	if currentState != nil && currentState.Hash != data.Hash {
		var currentStateEvent *models.SubscriptionEvent
		query.GetOne(entities.Event{Hash: currentState.Event.Hash}, &currentStateEvent)
		isMoreRecent, markAsSynced = IsMoreRecent(
			currentStateEvent.ID,
			currentState.Event.Hash,
			currentStateEvent.Payload.Timestamp,
			event.Hash,
			event.Payload.Timestamp,
			markAsSynced,
		 )
	}

	if currentState == nil || (currentState != nil && isMoreRecent) { // it is a morer ecent event
		markAsSynced = true
	}

	if event.Payload.EventType != uint16(constants.SubscribeTopicEvent) {
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
	}

	// Save stuff permanently
	tx := sql.Db.Begin()

	// If the event was not signed by your node
	if string(event.Validator) != (*cfg).NetworkPublicKey {
		// save the event
		event.Error = eventError
		event.IsValid = markAsSynced && len(eventError) == 0.
		event.Synced = markAsSynced
		event.Broadcasted = true
		_, _, err := query.SaveRecord(models.SubscriptionEvent{
			Event: entities.Event{
				PayloadHash: event.PayloadHash,
			},
		}, models.SubscriptionEvent{
			Event: *event,
		}, false, tx)
		if err != nil {
			tx.Rollback()
			logger.Fatal("5000: Db Error", err)
			return
		}
	} else {
		if markAsSynced {
			_, _, err := query.SaveRecord(models.SubscriptionEvent{
				Event: entities.Event{PayloadHash: event.PayloadHash},
			}, models.SubscriptionEvent{
				Event: entities.Event{Synced: true, Broadcasted: true, Error: eventError, IsValid: len(eventError) == 0},
			}, true, tx)
			if err != nil {
				logger.Fatal("DB error", err)
			}
		} else {
			// mark as broadcasted
			_, _, err := query.SaveRecord(models.SubscriptionEvent{
				Event: entities.Event{PayloadHash: event.PayloadHash, Broadcasted: false},
			},
				models.SubscriptionEvent{
					Event: entities.Event{Broadcasted: true},
				}, true, tx)
			if err != nil {
				logger.Fatal("DB error", err)
			}
		}
	}

	//Update subscription status based on the event type
	switch event.Payload.EventType {
	case uint16(constants.SubscribeTopicEvent):
		if *topicData.Public {
			data.Status = constants.SubscribedSubscriptionStatus
		} else {
			data.Status = constants.PendingSubscriptionStatus
		}
	case uint16(constants.LeaveEvent):
		data.Status = constants.UnsubscribedSubscriptionStatus
	case uint16(constants.ApprovedEvent):
		data.Status = constants.SubscribedSubscriptionStatus
	case uint16(constants.BanMemberEvent):
		data.Status = constants.BannedSubscriptionStatus
	case uint16(constants.UnbanMemberEvent):
		data.Status = constants.SubscribedSubscriptionStatus
	default:

	}

	data.Event = *entities.NewEventPath(event.Validator, entities.SubscriptionEventModel, event.Hash)
	data.Agent = entities.AddressFromString(agent).ToDeviceString()

	if markAsSynced && eventError == "" {
		updateState = true
	}

	data.Subnet = event.Payload.Subnet

	if updateState {
		_, _, err := query.SaveRecord(models.SubscriptionState{
			Subscription: entities.Subscription{ID: data.ID, Subnet: data.Subnet},
		}, models.SubscriptionState{
			Subscription: *data,
		}, true, tx)
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
			go HandleNewPubSubSubscriptionEvent(&dep, ctx)
		}
	}

	// TODO Broadcast the updated state

}
