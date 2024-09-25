package service

import (
	"context"
	"encoding/hex"
	"slices"

	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/p2p"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"gorm.io/gorm"
)

/*
Validate an agent authorization
*/
func ValidateSubscriptionData(payload *entities.ClientPayload, topic *entities.Topic) (currentSubscriptionState *models.SubscriptionState, err error) {
	// check fields of subscription

	subscription := payload.Data.(entities.Subscription)
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
func saveSubscriptionEvent(where entities.Event, createData *entities.Event, updateData *entities.Event, tx *gorm.DB) (*entities.Event, error) {
	var createModel *models.SubscriptionEvent
	if createData != nil {
		createModel = &models.SubscriptionEvent{Event: *createData}
	} else {
		createModel = &models.SubscriptionEvent{}
	}
	var updateModel *models.SubscriptionEvent
	if updateData != nil {
		updateModel = &models.SubscriptionEvent{Event: *updateData}
	}
	model, _, err := query.SaveRecord(models.SubscriptionEvent{Event: where},  createModel, updateModel, tx)
	if err != nil {
		return nil, err
	}
	return &model.Event, err
}
func HandleNewPubSubSubscriptionEvent(event *entities.Event, ctx *context.Context) {
	
	cfg, ok := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)
	if !ok {
		panic("Unable to load config from context")
	}
	data := event.Payload.Data.(entities.Subscription)
	// var id = data.ID
	// if len(data.ID) == 0 {
	// 	id, _ = entities.GetId(data)
	// } else {
	// 	id = data.ID
	// }
	var topic =  models.TopicState{}
	data.Event = *event.GetPath()
	data.BlockNumber = event.BlockNumber
	data.Cycle = event.Cycle
	data.Epoch = event.Epoch
	hash, err := data.GetHash()
	if err != nil {
		return
	}
	data.Hash = hex.EncodeToString(hash)
	var subnet = data.Subnet

	var localState models.SubscriptionState
	// err := query.GetOne(&models.TopicState{Topic: entities.Topic{ID: id}}, &localTopicState)
	err = sql.SqlDb.Where(&models.SubscriptionState{Subscription: entities.Subscription{Subnet: subnet, Topic: data.Topic, Agent: entities.AddressFromString(string(data.Agent)).ToDeviceString()}}).Take(&localState).Error
	if err != nil {
		logger.Error(err)
	}
	

	var localDataState *LocalDataState
	if localState.ID != "" {
		localDataState = &LocalDataState{
			ID: localState.ID,
			Hash: localState.Hash,
			Event: &localState.Event,
			Timestamp: *localState.Timestamp,
		}
	}
	// localDataState := utils.IfThenElse(localTopicState != nil, &LocalDataState{
	// 	ID: localTopicState.ID,
	// 	Hash: localTopicState.Hash,
	// 	Event: &localTopicState.Event,
	// 	Timestamp: localTopicState.Timestamp,
	// }, nil)
	var stateEvent *entities.Event
	if localState.ID != "" {
		stateEvent, err = query.GetEventFromPath(&localState.Event)
		if err != nil && err != query.ErrorNotFound {
			logger.Debug(err)
		}
	}
	var localDataStateEvent *LocalDataStateEvent
	if stateEvent != nil {
		localDataStateEvent = &LocalDataStateEvent{
			ID: stateEvent.ID,
			Hash: stateEvent.Hash,
			Timestamp: stateEvent.Timestamp,
		}
	}

	eventData := PayloadData{Subnet: subnet, localDataState: localDataState, localDataStateEvent:  localDataStateEvent}
	tx := sql.SqlDb
	// defer func () {
	// 	if tx.Error != nil {
	// 		tx.Rollback()
	// 	} else {
	// 		tx.Commit()
	// 	}
	// }()

	
	
	previousEventUptoDate,  authEventUpToDate, _, eventIsMoreRecent, err := ProcessEvent(event,  eventData, true, saveSubscriptionEvent, tx, ctx)
	if err != nil {
		logger.Debugf("Processing Error...: %v", err)
		return
	}
	logger.Debugf("Processing 2...: %v,  %v", previousEventUptoDate, authEventUpToDate)
	// get the topic, if not found retrieve it
	
	if previousEventUptoDate  && authEventUpToDate {

		err = query.GetOneState(entities.Topic{ID: data.Topic}, &topic)

		if topic.ID == "" || (err != nil && err == query.ErrorNotFound) {
			// get topic like we got subnet
			topicPath := entities.NewEntityPath(event.Validator, entities.TopicModel, data.Topic)
				pp, err := p2p.GetState(cfg, *topicPath, &event.Validator, &topic)
				if err != nil {
					logger.Error(err)
					
				}
				
				topicEvent, err := entities.UnpackEvent(pp.Event, entities.TopicModel)
				if err != nil {
					logger.Error(err)
					
				}
				if topicEvent != nil {
					HandleNewPubSubSubnetEvent(topicEvent, ctx)
					
				}
				return
		}
	
		_, err = ValidateSubscriptionData(&event.Payload, &topic.Topic)
		if err != nil {
			// update error and mark as synced
			// notify validator of error
			saveSubscriptionEvent(entities.Event{Hash: event.Hash}, nil, &entities.Event{Error: err.Error(), IsValid: false, Synced: true}, tx )
			
		} else {
			// TODO if event is older than our state, just save it and mark it as synced
			
			savedEvent, err := saveSubscriptionEvent(entities.Event{Hash: event.Hash}, nil, &entities.Event{IsValid: true, Synced: true}, tx );
			if eventIsMoreRecent && err == nil {
				// update state
				_, _, err := query.SaveRecord(models.SubscriptionState{
					Subscription: entities.Subscription{Topic: data.Topic, Subnet: data.Subnet, Agent: data.Agent},
				}, &models.SubscriptionState{
					Subscription: data,
				},  &models.SubscriptionState{
					Subscription: data,
				}, tx)
				if err != nil {
					// tx.Rollback()
					logger.Errorf("SaveStateError %v", err)
					return
				}
				
			}
			if err == nil {
				go OnFinishProcessingEvent(ctx, *event.GetPath(), &savedEvent.Payload.Subnet)
			}
			
			
			if string(event.Validator) != cfg.PublicKeyEDDHex {
				go func () {
				dependent, err := query.GetDependentEvents(event)
				if err != nil {
					logger.Debug("Unable to get dependent events", err)
				}
				for _, dep := range *dependent {
					HandleNewPubSubEvent(dep, ctx)
				}
				}()
			}
			
		}
	} 
		
	
}
