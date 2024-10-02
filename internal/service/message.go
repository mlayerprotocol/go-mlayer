package service

import (
	"context"
	"encoding/hex"

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
func ValidateMessageData(payload *entities.ClientPayload, topic *entities.Topic) (currentSubscription *models.SubscriptionState, err error) {
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
	message := payload.Data.(entities.Message)
	var subscription models.SubscriptionState
	// err = query.GetOne(models.SubscriptionState{
	// 	Subscription: entities.Subscription{Subscriber: payload.Account, Topic: topicData.ID},
	// }, &subscription)
	if payload.Account != message.Sender {
		return nil,  apperror.BadRequest("Invalid message signer")
	}

	
	subsribers := []entities.DIDString{entities.DIDString(payload.Agent), entities.DIDString(payload.Account.ToString())}
	subscriptions, err := query.GetSubscriptionStateBySubscriber(payload.Subnet, message.Topic, subsribers, sql.SqlDb)
	
	if err != nil {
		return
	}
	
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
		
		if *topic.ReadOnly && payload.Account != topic.Account && *subscription.Role < constants.TopicManagerRole {
			return nil, apperror.Unauthorized("Not allowed to post to this topic")
		}
		if payload.Account != topic.Account && *subscription.Role < constants.TopicWriterRole {
			return  nil, apperror.Unauthorized("Not allowed to post to this topic")
		}
		logger.Debugf("Found Subscribers: %v", subscription)
		return &subscription, nil
	} else {
		// check if the sender is a subnet admin
		 subnet := models.SubnetState{}
		err = query.GetOneState(entities.Subnet{ID: payload.Subnet}, &subnet)
		if err != nil {
			return nil, apperror.Unauthorized("Invalid subnet")
		}
		if payload.Account != subnet.Account {
			// check if its an admin
			auth := models.AuthorizationState{}
			err = query.GetOneState(entities.Authorization{Agent: payload.Agent, Account: payload.Account}, &auth)
			if err != nil {
				return nil, apperror.Unauthorized("Invalid subnet")
			}
			if *auth.Priviledge != constants.AdminPriviledge {
				return nil, apperror.Unauthorized("Not subscribed to topic or admin")
			}
			
		}
	}
	
	
	return &subscription, nil
}
func saveMessageEvent(where entities.Event, createData *entities.Event, updateData *entities.Event, tx *gorm.DB) (*entities.Event, error) {
	var createModel *models.MessageEvent
	if createData != nil {
		createModel = &models.MessageEvent{Event: *createData}
	} else {
		createModel = &models.MessageEvent{}
	}
	var updateModel *models.MessageEvent
	if updateData != nil {
		updateModel = &models.MessageEvent{Event: *updateData}
	}
	model, _, err := query.SaveRecord(models.MessageEvent{Event: where},  createModel, updateModel, tx)
	if err != nil {
		return nil, err
	}
	return &model.Event, err
}
func HandleNewPubSubMessageEvent(event *entities.Event, ctx *context.Context) {
	cfg, ok := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)
	if !ok {
		panic("Unable to load config from context")
	}
	data := event.Payload.Data.(entities.Message)
	var id = data.ID
	if len(data.ID) == 0 {
		id, _ = entities.GetId(data)
	} else {
		id = data.ID
	}
	var topic =  models.TopicState{}
	data.BlockNumber = event.BlockNumber
	data.Cycle = event.Cycle
	data.Epoch = event.Epoch
	data.Event = *event.GetPath()
	hash, err := data.GetHash()
	if err != nil {
		return
	}
	data.Hash = hex.EncodeToString(hash)
	var subnet = event.Payload.Subnet
	
	var localState models.MessageState
	// err := query.GetOne(&models.TopicState{Topic: entities.Topic{ID: id}}, &localTopicState)
	err = sql.SqlDb.Where(&models.MessageState{Message: entities.Message{ ID: id, Agent: entities.AddressFromString(string(data.Agent)).ToDeviceString()}}).Take(&localState).Error
	if err != nil {
		logger.Error(err)
	}
	
	
	var localDataState *LocalDataState
	if localState.ID != "" {
		localDataState = &LocalDataState{
			ID: localState.ID,
			Hash: localState.Hash,
			Event: &localState.Event,
			Timestamp: uint64(event.Payload.Timestamp),
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

	
	
	previousEventUptoDate,  authEventUpToDate, _, eventIsMoreRecent, err := ProcessEvent(event,  eventData, true, saveMessageEvent, tx, ctx)
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
					return
				}
				if len(pp.Event) < 2 {
					return 
				}
				topicEvent, err := entities.UnpackEvent(pp.Event, entities.TopicModel)
				if err != nil {
					logger.Error(err)
					return 
				}
				if topicEvent != nil {
					HandleNewPubSubEvent(*topicEvent, ctx)
					return
				}
		}
		
	
		 _, err = ValidateMessageData(&event.Payload, &topic.Topic)
		if err != nil {
			// update error and mark as synced
			// notify validator of error
			saveMessageEvent(entities.Event{Hash: event.Hash}, nil, &entities.Event{Error: err.Error(), IsValid: false, Synced: true}, tx )
			
		} else {
			// TODO if event is older than our state, just save it and mark it as synced
			
			savedEvent, err := saveMessageEvent(entities.Event{Hash: event.Hash}, nil, &entities.Event{IsValid: true, Synced: true}, tx );
			if eventIsMoreRecent && err == nil {
				// update state
				if data.ID != "" {
					_, _, err = query.SaveRecord(models.MessageState{
						Message: entities.Message{ID: data.ID},
					}, &models.MessageState{
						Message: data,
					},  &models.MessageState{
						Message: data,
					}, tx)
				} else {
					createModel := models.MessageState{ Message: data}
					err = tx.Create(&createModel).Error
				}
				if err != nil {
					// tx.Rollback()
					logger.Errorf("SaveStateError %v", err)
					return
				}
				
			}
			if err == nil {
				go OnFinishProcessingEvent(ctx, event,  &models.MessageState{
					Message: data,
				},  &savedEvent.Payload.Subnet)
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
