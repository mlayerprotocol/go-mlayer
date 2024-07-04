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
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"github.com/sirupsen/logrus"
)

/*
Validate an agent authorization
*/
func ValidateMessageData(message *entities.Message, payload *entities.ClientPayload, topic *entities.Topic) (currentMessageState *models.MessageState, subscriptionState *models.SubscriptionState, err error) {
	// check fields of message
	// var currentState *models.MessageState

	// err = query.GetOne(models.MessageState{
	// 	Message: entities.Message{Subscriber: message.Subscriber, Topic: message.Topic},
	// }, &currentState)
	// if err != nil {
	// 	if err != gorm.ErrRecordNotFound {
	// 		//return nil, nil, apperror.Unauthorized("Not a subscriber")
	// 		// } else {
	// 		return nil, err
	// 	}
	// }
	var subscription models.SubscriptionState
	// err = query.GetOne(models.SubscriptionState{
	// 	Subscription: entities.Subscription{Subscriber: payload.Account, Topic: topicData.ID},
	// }, &subscription)
	subsribers := []entities.DIDString{entities.DIDString(payload.Agent), entities.DIDString(payload.Account.ToString())}
	subscriptions, err := query.GetSubscriptionStateBySuscriber(payload.Subnet, message.TopicId, subsribers, nil)
	if len(*subscriptions) > 0 {
		if  len(*subscriptions) > 1 {
			// if string(payload.Account)  != "" && (*subscriptions)[0].Subscription.Subscriber.ToString() == string(payload.Account) {
			// 	subscription = (*subscriptions)[0]
			// } else {
			// 	subscription = (*subscriptions)[1]
			// }
			if  *((*subscriptions)[0].Subscription.Role) > *((*subscriptions)[1].Subscription.Role) {
				subscription = (*subscriptions)[0]
			} else {
				subscription = (*subscriptions)[1]
			}
		} else {
			subscription = (*subscriptions)[0]
		}
	}
	
	if *topic.ReadOnly && payload.Account != topic.Account && *subscription.Role < constants.TopicManagerRole {
		return nil, nil, apperror.Unauthorized("Not allowed to post to this topic")
	}
	if payload.Account != topic.Account && *subscription.Role < constants.TopicWriterRole {
		return nil, nil, apperror.Unauthorized("Not allowed to post to this topic")
	}
	return nil,  subscriptionState, nil
}

func HandleNewPubSubMessageEvent(event *entities.Event, ctx *context.Context) {
	logger.WithFields(logrus.Fields{"event": event}).Debug("New message event from pubsub channel")
	markAsSynced := false
	updateState := false
	var eventError string
	// hash, _ := event.GetHash()
	err := ValidateEvent(*event)

	if err != nil {
		logger.Error(err)
		return
	}
	// d, err := event.Payload.EncodeBytes()
	// if err != nil {
	// 	logger.Errorf("Invalid event payload")
	// }
	// agent, err := crypto.GetSignerECC(&d, &event.Payload.Signature)
	// if err != nil {
	// 	logger.Errorf("Invalid event payload")
	// }

	logger.Infof("HandleNewPubSubMessageEvent Event is a valid event %s", event.PayloadHash)
	cfg, _ := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)

	// Extract and validate the Data of the paylaod which is an Events Payload Data,
	data := event.Payload.Data.(*entities.Message)
	hash, _ := data.GetHash()
	data.Hash = hex.EncodeToString(hash)
	data.Agent = event.Payload.Agent

	topicData, err := query.GetTopicById(data.TopicId)
	if err != nil {
		eventError = err.Error()
	}

	if topicData == nil {
		eventError = "Invalid topic id"
	}


	currentState, _, _ := ValidateMessageData(data, &event.Payload, &topicData.Topic)
	prevEventUpToDate := false
	authEventUpToDate := false

	// check if we are upto date on this event
	prevEventUpToDate = query.EventExist(&event.PreviousEventHash) || (currentState == nil && event.PreviousEventHash.Hash == "") || (currentState != nil && currentState.Event.Hash == event.PreviousEventHash.Hash)

	authState, authError := query.GetOneAuthorizationState(entities.Authorization{Event: event.AuthEventHash})

	authEventUpToDate = query.EventExist(&event.AuthEventHash) || (authState.ID == "" && event.AuthEventHash.Hash == "") || (authState.ID != "" && authState.Event.Hash == event.AuthEventHash.Hash)
	logger.Infof("prevEventUpToDate %t ----- authEventUpToDate %t ---- event.PreviousEventHash.Hash %s ", prevEventUpToDate, authEventUpToDate, event.PreviousEventHash)
	// Confirm if this is an older event coming after a newer event.
	// If it is, then we only have to update our event history, else we need to also update our current state
	isMoreRecent := false
	if currentState != nil && currentState.Hash != data.Hash {
		var currentStateEvent *models.MessageEvent
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

	// Save stuff permanently
	tx := sql.Db.Begin()

	// If the event was not signed by your node
	if string(event.Validator) != (*cfg).NetworkPublicKey {
		// save the event
		event.Error = eventError
		event.IsValid = markAsSynced && len(eventError) == 0.
		event.Synced = markAsSynced
		event.Broadcasted = true
		_, _, err := query.SaveRecord(models.MessageEvent{
			Event: entities.Event{
				PayloadHash: event.PayloadHash,
			},
		}, models.MessageEvent{
			Event: *event,
		}, false, tx)
		if err != nil {
			tx.Rollback()
			logger.Fatal("5000: Db Error", err)
			return
		}
	} else {
		if markAsSynced {
			_, _, err := query.SaveRecord(models.MessageEvent{
				Event: entities.Event{PayloadHash: event.PayloadHash},
			}, models.MessageEvent{
				Event: entities.Event{Synced: true, Broadcasted: true, Error: eventError, IsValid: len(eventError) == 0},
			}, true, tx)
			if err != nil {
				logger.Fatal("DB error", err)
			}
		} else {
			// mark as broadcasted
			_, _, err := query.SaveRecord(models.MessageEvent{
				Event: entities.Event{PayloadHash: event.PayloadHash, Broadcasted: false},
			},
				models.MessageEvent{
					Event: entities.Event{Broadcasted: true},
				}, true, tx)
			if err != nil {
				logger.Fatal("DB error", err)
			}
		}
	}

	//Update message status based on the event type

	data.Event = *entities.NewEventPath(event.Validator, entities.MessageEventModel, event.Hash)
	// data.Agent = entities.DIDString(agent)

	if markAsSynced && eventError == "" {
		updateState = true
	}
	logger.Infof("Lst Event is a valid event %t --- %s", markAsSynced, eventError)
	var newState *models.MessageState
	if updateState {
		newState, _, err = query.SaveRecord(models.MessageState{
			Message: entities.Message{Hash: data.Hash},
		}, models.MessageState{
			Message: *data,
		}, true, tx)
		if err != nil {
			tx.Rollback()
			logger.Fatal("5000: Db Error", err)
			return
		}
	}
	tx.Commit()
	go OnFinishProcessingEvent(ctx, &data.Event, utils.IfThenElse(newState!=nil, &newState.ID, nil), utils.IfThenElse(event.Error!="", apperror.Internal(event.Error), nil))

	if string(event.Validator) != (*cfg).NetworkPublicKey {
		dependent, err := query.GetDependentEvents(*event)
		if err != nil {
			logger.Info("Unable to get dependent events", err)
		}
		for _, dep := range *dependent {
			go HandleNewPubSubMessageEvent(&dep, ctx)
		}
	}

	// TODO Broadcast the updated state

}
