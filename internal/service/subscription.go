package service

import (
	"context"
	"encoding/hex"
	"slices"
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
func ValidateSubscriptionData(subscription *entities.Subscription, payload *entities.ClientPayload, topic *entities.Topic) (currentSubscriptionState *models.SubscriptionState, err error) {
	// check fields of subscription
	var currentState *models.SubscriptionState
	

	err = query.GetOne(models.SubscriptionState{
		Subscription: entities.Subscription{Subscriber: subscription.Subscriber, Subnet: subscription.Subnet, Topic: subscription.Topic},
	}, &currentState)
	if err != nil {
		logger.Errorf("Invalid event payload %e ", err)

		if err != gorm.ErrRecordNotFound {
			//return nil, nil, apperror.Unauthorized("Not a subscriber")
			// } else {
			return nil, err
		} else {
			logger.Errorf("gorm.ErrRecordNotFound %e ", gorm.ErrRecordNotFound)
			return nil, nil
		}
	}
	if payload.EventType == uint16(constants.SubscribeTopicEvent) { 
		// someone inviting someone else
		if subscription.Subscriber != payload.Account && subscription.Agent != payload.Agent {
			if !slices.Contains([]constants.SubscriptionStatus{constants.InvitedSubscriptionStatus, constants.BannedSubscriptionStatus}, *subscription.Status) {
				return nil, apperror.Forbidden("Subscription status must be Invited or Banned")
			}
		}
	} else {
		// subscribing oneself
		// if the topic is not public, you have to have been invited
		if !(*topic.Public) && currentState == nil {
			return nil, apperror.Forbidden("Must be invited first")
		}
		if  currentState != nil && *currentState.Status == constants.BannedSubscriptionStatus {
			return nil, apperror.Forbidden("Banned subscriber")
		}
		if (currentState != nil && *currentState.Role != *subscription.Role) || (currentState == nil && *subscription.Role > *topic.DefaultSubscriberRole) {
			return nil, apperror.Forbidden("Invalid role selected")
		}
		if !slices.Contains([]constants.SubscriptionStatus{constants.UnsubscribedSubscriptionStatus, constants.SubscribedSubscriptionStatus}, *subscription.Status) {
			return nil, apperror.Forbidden("Subscription status must be Subscribed or Unsubscribed")
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
	logger.Infof("ValidateSubscriptionData %v", data)
	currentState, authError := ValidateSubscriptionData(data, &event.Payload, &topicData.Topic)
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

	if currentState == nil || isMoreRecent { // it is a morer ecent event
		markAsSynced = true
	}

	if currentState != nil && data.Subscriber == topicData.Account && event.EventType == uint16(constants.SubscribeTopicEvent) {
		authError =  apperror.BadRequest("Topic already owned by account")
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
				if currentState == nil || isMoreRecent { // it is a morer ecent event
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
				if currentState == nil ||  isMoreRecent {
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
	tx := sql.SqlDb.Begin()

	// If the event was not signed by your node
	if string(event.Validator) != (*cfg).PublicKey  {
		// save the event
		event.Error = eventError
		event.IsValid = markAsSynced && len(eventError) == 0.
		event.Synced = markAsSynced
		event.Broadcasted = true
		_, _, err := query.SaveRecord(models.SubscriptionEvent{
			Event: entities.Event{
				PayloadHash: event.PayloadHash,
			},
		}, &models.SubscriptionEvent{
			Event: *event,
		}, nil, tx)
		if err != nil {
			tx.Rollback()
			logger.Error("5000: Db Error", err)
			return
		}
	} else {
		if markAsSynced {
			_, _, err := query.SaveRecord(models.SubscriptionEvent{
				Event: entities.Event{PayloadHash: event.PayloadHash},
			}, 
			&models.SubscriptionEvent{
				Event: *event,
			},
			&models.SubscriptionEvent{
				Event: entities.Event{Synced: true, Broadcasted: true, Error: eventError, IsValid: len(eventError) == 0},
			}, tx)
			if err != nil {
				logger.Error("DB error", err)
			}
		} else {
			// mark as broadcasted
			_, _, err := query.SaveRecord(models.SubscriptionEvent{
				Event: entities.Event{PayloadHash: event.PayloadHash, Broadcasted: false},
			},
			&models.SubscriptionEvent{
				Event: *event,
			},
				&models.SubscriptionEvent{
					Event: entities.Event{Broadcasted: true},
				},
			tx)
			if err != nil {
				logger.Error("DB error", err)
			}
		}
	}

	//Update subscription status based on the event type
	switch event.Payload.EventType {
	// case uint16(constants.SubscribeTopicEvent):
	// 	if *topicData.Public {
	// 		data.Status = &constants.SubscribedSubscriptionStatus
	// 	} else {
	// 		data.Status = &constants.PendingSubscriptionStatus
	// 	}
	case uint16(constants.LeaveEvent):
		data.Status = &constants.UnsubscribedSubscriptionStatus
	case uint16(constants.ApprovedEvent):
		data.Status = &constants.SubscribedSubscriptionStatus
	case uint16(constants.BanMemberEvent):
		data.Status = &constants.BannedSubscriptionStatus
	case uint16(constants.UnbanMemberEvent):
		data.Status = &constants.SubscribedSubscriptionStatus
	default:

	}

	data.Event = *entities.NewEventPath(event.Validator, entities.SubscriptionModel, event.Hash)
	data.Agent = entities.AddressFromString(agent).ToDeviceString()

	if markAsSynced && eventError == "" {
		updateState = true
	}

	data.Subnet = event.Payload.Subnet
	var newState *models.SubscriptionState
	if updateState {
		newState, _, err = query.SaveRecord(models.SubscriptionState{
			Subscription: entities.Subscription{ID: data.ID, Subnet: data.Subnet, Subscriber: data.Subscriber, Topic: data.Topic},
		}, &models.SubscriptionState{
			Subscription: *data,
		}, &models.SubscriptionState{
			Subscription: *data,
		}, tx)
		if err != nil {
			tx.Rollback()
			logger.Error("5000: Db Error", err)
			return
		}
	}
	tx.Commit()
	if markAsSynced {
		go OnFinishProcessingEvent(ctx, &data.Event, utils.IfThenElse(newState!=nil, &newState.ID, nil), utils.IfThenElse(event.Error!="", apperror.Internal(event.Error), nil))
	}
	if string(event.Validator) != (*cfg).PublicKey  {
		dependent, err := query.GetDependentEvents(event)
		if err != nil {
			logger.Info("Unable to get dependent events", err)
		}
		for _, dep := range *dependent {
			go HandleNewPubSubSubscriptionEvent(&dep, ctx)
		}
	}

	// TODO Broadcast the updated state

}
