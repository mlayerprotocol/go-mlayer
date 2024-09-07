package service

import (
	"context"
	"encoding/hex"

	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"gorm.io/gorm"
)

/*
Validate an agent authorization
*/
func ValidateTopicData(topic *entities.Topic, authState *models.AuthorizationState) (currentTopicState *models.TopicState, err error) {
	
	//subnet := models.SubnetState{}

	// TODO state might have changed befor receiving event, so we need to find state that is relevant to this event.
	// err = query.GetOne(models.SubnetState{Subnet: entities.Subnet{ID: topic.Subnet}}, &subnet)
	// if err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		return nil, apperror.Forbidden("Invalid subnet id")
	// 	}
	// 	return nil, apperror.Internal(err.Error())
	// }
	if authState != nil && *authState.Priviledge < constants.StandardPriviledge {
		return nil, apperror.Forbidden("Agent does not have enough permission to create topics")
	}

	if len(topic.Ref) > 40 {
		return nil, apperror.BadRequest("Topic reference can not be more than 40 characters")
	}
	if !utils.IsAlphaNumericDot(topic.Ref) {
		return nil, apperror.BadRequest("Reference must be alphanumeric, _ and . but cannot start with a number")
	}
	return nil, nil
}

func saveTopicEvent(where entities.Event, createData *entities.Event, updateData *entities.Event, tx *gorm.DB) (*entities.Event, error) {
	var createModel = models.TopicEvent{}
	if createData != nil {
		createModel = models.TopicEvent{Event: *createData}
	}
	var updateModel = models.TopicEvent{}
	if updateData != nil {
		updateModel = models.TopicEvent{Event: *updateData}
	}
	model, _, err := query.SaveRecord(models.TopicEvent{Event: where},  &createModel, &updateModel, tx)
	if err != nil {
		return nil, err
	}
	
	return &model.Event, err
}

func HandleNewPubSubTopicEvent(event *entities.Event, ctx *context.Context) {
	cfg, ok := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)
	if !ok {
		panic("Unable to get config from context")
	}
	data := event.Payload.Data.(entities.Topic)

	var id = data.ID
	if len(data.ID) == 0 {
		id, _ = entities.GetId(data)
	} else {
		id = data.ID
	}
	data.Event = *event.GetPath()
	data.BlockNumber = event.BlockNumber
	data.Cycle = event.Cycle
	data.Epoch = event.Epoch
	hash, err := data.GetHash()
	if err != nil {
		return
	}
	data.Hash = hex.EncodeToString(hash)
	data.Account = event.Payload.Account
	data.Agent = event.Payload.Agent
	data.Timestamp = event.Payload.Timestamp
	logger.Debug("Processing 1...")
	var localState models.TopicState
	// err := query.GetOne(&models.TopicState{Topic: entities.Topic{ID: id}}, &localState)
	err = sql.SqlDb.Where(&models.TopicState{Topic: entities.Topic{ID: id}}).Take(&localState).Error
	if err != nil {
		logger.Error(err)
	}
	logger.Debug("Processing 2...")
	
	var localDataState *LocalDataState
	if localState.ID == "" {
		localDataState = &LocalDataState{
			ID: localState.ID,
			Hash: localState.Hash,
			Event: &localState.Event,
			Timestamp: localState.Timestamp,
		}
	}
	// localDataState := utils.IfThenElse(localState != nil, &LocalDataState{
	// 	ID: localState.ID,
	// 	Hash: localState.Hash,
	// 	Event: &localState.Event,
	// 	Timestamp: localState.Timestamp,
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

	eventData := PayloadData{Subnet: data.Subnet, localDataState: localDataState, localDataStateEvent:  localDataStateEvent}
	tx := sql.SqlDb
	// defer func () {
	// 	if tx.Error != nil {
	// 		tx.Rollback()
	// 	} else {
	// 		tx.Commit()
	// 	}
	// }()
	previousEventUptoDate,  authEventUptoDate, authState, eventIsMoreRecent, err := ProcessEvent(event,  eventData, true, saveTopicEvent, tx, ctx)
	if err != nil {
		return
	}
	
	if previousEventUptoDate &&  authEventUptoDate {
		_, err = ValidateTopicData(&data, authState)
		if err != nil {
			// update error and mark as synced
			// notify validator of error
			saveTopicEvent(entities.Event{Hash: event.Hash}, nil, &entities.Event{Error: err.Error(), IsValid: false, Synced: true}, tx )
			
		} else {
			// TODO if event is older than our state, just save it and mark it as synced
			savedEvent, err := saveTopicEvent(entities.Event{Hash: event.Hash}, nil, &entities.Event{IsValid: true, Synced: true}, tx )
			if eventIsMoreRecent {
				logger.Debug("ISMORERECENT", eventIsMoreRecent)
				// update state
				if data.ID != "" {
					_, _, err = query.SaveRecord(models.TopicState{
						Topic: entities.Topic{ID: data.ID},
					}, &models.TopicState{
						Topic: data,
					}, utils.IfThenElse(event.EventType == uint16(constants.UpdateTopicEvent), &models.TopicState{
						Topic: data,
					}, &models.TopicState{}) , tx)
				} else {
					createModel := models.TopicState{ Topic: data}
					err = tx.Create(&createModel).Error
				}
				if err != nil {
					// tx.Rollback()
					return
				}
			}
			if err == nil {
				go OnFinishProcessingEvent(ctx, event.GetPath(), &savedEvent.Payload.Subnet, err)
			}
			
			
			if string(event.Validator) != cfg.PublicKey {
				go func () {
				dependent, err := query.GetDependentEvents(event)
				if err != nil {
					logger.Debug("Unable to get dependent events", err)
				}
				for _, dep := range *dependent {
					HandleNewPubSubEvent(&dep, ctx)
				}
				}()
			}
		}


	}
	
}
