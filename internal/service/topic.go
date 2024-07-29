package service

import (
	"context"

	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
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
		return nil, apperror.BadRequest("Topic handle can not be more than 40 characters")
	}
	if !utils.IsAlphaNumericDot(topic.Ref) {
		return nil, apperror.BadRequest("Handle must be alphanumeric, _ and . but cannot start with a number")
	}
	return nil, nil
}

func saveTopicEvent(where entities.Event, createData *entities.Event, updateData *entities.Event, tx *gorm.DB) (*entities.Event, error) {
	var createModel *models.TopicEvent
	if createData != nil {
		createModel = &models.TopicEvent{Event: *createData}
	}
	var updateModel *models.TopicEvent
	if updateData != nil {
		updateModel = &models.TopicEvent{Event: *updateData}
	}
	model, _, err := query.SaveRecord(models.TopicEvent{Event: where},  createModel, updateModel, tx)
	if err != nil {
		return nil, err
	}
	return &model.Event, err
}

func HandleNewPubSubTopicEvent(event *entities.Event, ctx *context.Context) {
	
	data := event.Payload.Data.(entities.Topic)

	var id = data.ID
	if len(data.ID) == 0 {
		id, _ = entities.GetId(data)
	} else {
		id = data.ID
	}
	logger.Info("Processing 1...")
	var localTopicState models.TopicState
	// err := query.GetOne(&models.TopicState{Topic: entities.Topic{ID: id}}, &localTopicState)
	err := sql.SqlDb.Where(&models.TopicState{Topic: entities.Topic{ID: id}}).Take(&localTopicState).Error
	if err != nil {
		logger.Error(err)
	}
	logger.Info("Processing 2...")
	var one = 1
	if one+3 == 4 {
		return
	}
	var localDataState *LocalDataState
	if localTopicState.ID == "" {
		localDataState = &LocalDataState{
			ID: localTopicState.ID,
			Hash: localTopicState.Hash,
			Event: &localTopicState.Event,
			Timestamp: localTopicState.Timestamp,
		}
	}
	// localDataState := utils.IfThenElse(localTopicState != nil, &LocalDataState{
	// 	ID: localTopicState.ID,
	// 	Hash: localTopicState.Hash,
	// 	Event: &localTopicState.Event,
	// 	Timestamp: localTopicState.Timestamp,
	// }, nil)
	var topicEvent *entities.Event
	if localTopicState.ID != "" {
		topicEvent, err = query.GetEventFromPath(&localTopicState.Event)
		if err != nil && err != query.ErrorNotFound {
			logger.Debug(err)
		}
	}
	var localDataStateEvent *LocalDataStateEvent
	if topicEvent != nil {
		localDataStateEvent = &LocalDataStateEvent{
			ID: topicEvent.ID,
			Hash: topicEvent.Hash,
			Timestamp: topicEvent.Timestamp,
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
			saveTopicEvent(entities.Event{Hash: event.Hash}, nil, &entities.Event{IsValid: true, Synced: true}, tx )
			if eventIsMoreRecent {
				// update state
				_, _, err := query.SaveRecord(models.TopicState{
					Topic: entities.Topic{ID: data.ID},
				}, &models.TopicState{
					Topic: data,
				}, utils.IfThenElse(event.EventType == uint16(constants.UpdateTopicEvent), &models.TopicState{
					Topic: data,
				}, &models.TopicState{}) , tx)
				if err != nil {
					// tx.Rollback()
					return
				}
			}
		}



		// check a situation where we have either of current auth and event state locally, but the states events are not same as the events prev auth and event
		// get the topics state
		// var id string
		// var badEvent  error
		// if len(data.ID) == 0 {
		// 	id, _ = entities.GetId(data)
		// } else {
		// 	id = data.ID
		// }

		// var topic models.TopicState
		// var topicEvent *entities.Event
		// err := query.GetOne(models.TopicState{Topic: entities.Topic{ID: id}}, &topic)
		// if err != nil && err != query.ErrorNotFound {
		// 	logger.Debug(err)
		// }
		// if len(topic.ID) > 0 {
		// 	// check if state.Event is same as events previous has
		// 	if topic.Event.Hash != event.PreviousEventHash.Hash {
		// 		// either we are not upto date, or the sender is not
		// 		// get the event that resulted in current state
		// 		topicEvent, err = query.GetEventFromPath(&topic.Event)
		// 		if err != nil && err != query.ErrorNotFound {
		// 			logger.Debug(err)
		// 		}
		// 		if len(topicEvent.ID) > 0 {
		// 			eventIsMoreRecent = IsMoreRecentEvent(topicEvent.Hash, int(topicEvent.Timestamp), event.Hash, int(event.Timestamp))


		// 			 // if this event is more recent, then it must referrence our local event or an event after it
		// 			if eventIsMoreRecent  && topicEvent.Hash != event.PreviousEventHash.Hash {
		// 				previousEventMoreRecent := IsMoreRecentEvent(topicEvent.Hash, int(topicEvent.Timestamp), previousEvent.Hash, int(previousEvent.Timestamp))
		// 				if !previousEventMoreRecent {
		// 					badEvent = fmt.Errorf(constants.ErrorBadRequest)
		// 				}
		// 			}

		// 		}
				
		// 	}
		// }

		





		// // We need to determin which authstate is valid for this event
		// // if its an old event, we can just save it since its not updating state
		// // if its a new one, we have to confirm that it is referencing the true latest auth event

		// // so lets get the referrenced authorization
		// var auth models.AuthorizationState
		// err = query.GetOne(models.AuthorizationState{Authorization: entities.Authorization{Event: event.AuthEventHash}}, &auth)
		// if err != nil && err != query.ErrorNotFound {
		// 	return
		// }
		// // if event is more recent that our local state, we have to check its validity since it updates state
		// if eventIsMoreRecent && agentAuthState != nil && agentAuthState.Event.Hash != event.AuthEventHash.Hash && authEvent != nil {
		// 	// get the event that is responsible for the current state
		// 	err := query.GetOne(models.AuthorizationEvent{Event: entities.Event{Hash: agentAuthState.Event.Hash}}, &agentAuthStateEvent)
		// 	if err != nil && err != query.ErrorNotFound {
		// 		logger.Debug(err)
		// 	}
		// 	if len(agentAuthStateEvent.ID) > 0 {
		// 		authMoreRecent = IsMoreRecentEvent(agentAuthStateEvent.Hash, int(agentAuthStateEvent.Timestamp), authEvent.Hash, int(authEvent.Timestamp))
		// 		if !authMoreRecent {
		// 			// this is a bad event using an old auth state.
		// 			// REJECT IT
		// 			badEvent = fmt.Errorf(constants.ErrorUnauthorized)
		// 		}
		// 	}
		// }



		// if badEvent != nil {
		// 	// update the event state with the error
		// 	_, _, err = query.SaveRecord(models.TopicEvent{Event: entities.Event{Hash: event.Hash}}, &models.TopicEvent{Event: *event}, &models.TopicEvent{Event: entities.Event{Error: badEvent.Error(), IsValid: false, Synced: true}}, nil)
		// 	if err != nil {
		// 		logger.Error(err)
		// 	}
		// 	// notify the originator so it can correct it e.g. let it know that there is a new authorization

		// 	// decide whether to report the node
		// 	return
		// }
		
		// _, err = ValidateTopicData(data, &auth)
		// if err != nil {
		// 	return
		// }
		// // TODO if event is older than our state, just save it and mark it as synced
		// if !eventIsMoreRecent {
		// 	return
		// }
		// HandleNewPubSubEventV2(eventModel, event, ctx)
	}
	// HandleNewPubSubEvent(entities.Topic{}, event, ctx)
	// err = sql.SqlDb.Where(models.TopicEvent{Event: entities.Event{Hash: event.Hash}}).FirstOrCreate(event)
	// hash, _ := data.GetHash()
	// data.Hash = hex.EncodeToString(hash)
	// authEventHash := event.AuthEventHash
	// authState, authError := query.GetOneAuthorizationState(entities.Authorization{Event: authEventHash})

	// currentState, err := ValidateTopicData(data, authState)
	// if err != nil {
	// 	// penalize node for broadcasting invalid data
	// 	logger.Infof("Invalid topic data %v. Node should be penalized", err)
	// 	return
	// }

	// // check if we are upto date on this event

	// prevEventUpToDate := query.EventExist(&event.PreviousEventHash) || (currentState == nil && event.PreviousEventHash.Hash == "") || (currentState != nil && currentState.Event.Hash == event.PreviousEventHash.Hash)
	// authEventUpToDate := query.EventExist(&event.AuthEventHash) || (authState == nil && event.AuthEventHash.Hash == "") || (authState != nil && authState.Event == authEventHash)

	// // Confirm if this is an older event coming after a newer event.
	// // If it is, then we only have to update our event history, else we need to also update our current state
	// isMoreRecent := false
	// if currentState != nil && currentState.Hash != data.Hash {
	// 	var currentStateEvent = &models.TopicEvent{}
	// 	query.GetOne(entities.Event{Hash: currentState.Event.Hash}, currentStateEvent)
	// 	isMoreRecent, markAsSynced = IsMoreRecent(
	// 		currentStateEvent.ID,
	// 		currentState.Event.Hash,
	// 		currentStateEvent.Payload.Timestamp,
	// 		event.Hash,
	// 		event.Payload.Timestamp,
	// 		markAsSynced,
	// 	)
	// 	// if uint64(currentStateEvent.Payload.Timestamp) < uint64(event.Payload.Timestamp) {
	// 	// 	isMoreRecent = true
	// 	// }
	// 	// if uint64(currentStateEvent.Payload.Timestamp) > uint64(event.Payload.Timestamp) {
	// 	// 	isMoreRecent = false
	// 	// }
	// 	// // if the authorization was created at exactly the same time but their hash is different
	// 	// // use the last 4 digits of their event hash
	// 	// if uint64(currentStateEvent.Payload.Timestamp) == uint64(event.Payload.Timestamp) {
	// 	// 	// get the event payload of the current state

	// 	// 	if err != nil && err != gorm.ErrRecordNotFound {
	// 	// 		logger.Error("DB error", err)
	// 	// 	}
	// 	// 	if currentStateEvent.ID == "" {
	// 	// 		markAsSynced = false
	// 	// 	} else {
	// 	// 		if currentStateEvent.Payload.Timestamp < event.Payload.Timestamp {
	// 	// 			isMoreRecent = true
	// 	// 		}
	// 	// 		if currentStateEvent.Payload.Timestamp == event.Payload.Timestamp {
	// 	// 			// logger.Infof("Current state %v", currentStateEvent.Payload)
	// 	// 			csN := new(big.Int)
	// 	// 			csN.SetString(currentState.Event.Hash[56:], 16)
	// 	// 			nsN := new(big.Int)
	// 	// 			nsN.SetString(event.Hash[56:], 16)

	// 	// 			if csN.Cmp(nsN) < 1 {
	// 	// 				isMoreRecent = true
	// 	// 			}
	// 	// 		}
	// 	// 	}
	// 	// }
	// }

	// if authError != nil {
	// 	// check if we are upto date. If we are, then the error is an actual one
	// 	// the error should be attached when saving the event
	// 	// But if we are not upto date, then we might need to wait for more info from the network

	// 	if prevEventUpToDate && authEventUpToDate {
	// 		// we are upto date. This is an actual error. No need to expect an update from the network
	// 		eventError = authError.Error()
	// 		markAsSynced = true
	// 	} else {
	// 		if currentState == nil || (currentState != nil && isMoreRecent) { // it is a morer ecent event
	// 			if strings.HasPrefix(authError.Error(), constants.ErrorForbidden) || strings.HasPrefix(authError.Error(), constants.ErrorUnauthorized) {
	// 				markAsSynced = false
	// 			} else {
	// 				// entire event can be considered bad since the payload data is bad
	// 				// this should have been sorted out before broadcasting to the network
	// 				// TODO penalize the node that broadcasted this
	// 				eventError = authError.Error()
	// 				markAsSynced = true
	// 			}

	// 		} else {
	// 			// we are upto date. We just need to store this event as well.
	// 			// No need to update state
	// 			markAsSynced = true
	// 			eventError = authError.Error()
	// 		}
	// 	}

	// }

	// // If no error, then we should act accordingly as well
	// // If are upto date, then we should update the state based on if its a recent or old event
	// if len(eventError) == 0 {
	// 	if prevEventUpToDate && authEventUpToDate { // we are upto date
	// 		if currentState == nil || isMoreRecent {
	// 			updateState = true
	// 			markAsSynced = true
	// 		} else {
	// 			// Its an old event
	// 			markAsSynced = true
	// 			updateState = false
	// 		}
	// 	} else {
	// 		updateState = false
	// 		markAsSynced = false
	// 	}

	// }

	// // Save stuff permanently
	// tx := sql.Db.Begin()
	// logger.Info(":::::updateState: Db Error", updateState, currentState == nil)

	// // If the event was not signed by your node
	// if string(event.Validator) != (*cfg).PublicKey  {
	// 	// save the event
	// 	event.Error = eventError
	// 	event.IsValid = markAsSynced && len(eventError) == 0.
	// 	event.Synced = markAsSynced
	// 	event.Broadcasted = true
	// 	_, _, err := query.SaveRecord(models.TopicEvent{
	// 		Event: entities.Event{
	// 			PayloadHash: event.PayloadHash,
	// 		},
	// 	}, &models.TopicEvent{
	// 		Event: *event,
	// 	}, nil, tx)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		logger.Error("1000: Db Error", err)
	// 		return
	// 	}
	// } else {
	// 	if markAsSynced {
	// 		_, _, err := query.SaveRecord(models.TopicEvent{
	// 			Event: entities.Event{PayloadHash: event.PayloadHash},
	// 		},
	// 		&models.TopicEvent{
	// 			Event: *event,
	// 		},
	// 		&models.TopicEvent{
	// 			Event: entities.Event{Synced: true, Broadcasted: true, Error: eventError, IsValid: len(eventError) == 0},
	// 		}, tx)
	// 		if err != nil {
	// 			logger.Error("DB error", err)
	// 		}
	// 	} else {
	// 		// mark as broadcasted
	// 		_, _, err := query.SaveRecord(models.TopicEvent{
	// 			Event: entities.Event{PayloadHash: event.PayloadHash, Broadcasted: false},
	// 		},
	// 		&models.TopicEvent{
	// 			Event: *event,
	// 		},
	// 			&models.TopicEvent{
	// 				Event: entities.Event{Broadcasted: true},
	// 			}, tx)
	// 		if err != nil {
	// 			logger.Error("DB error", err)
	// 		}
	// 	}
	// }

	// d, err := event.Payload.EncodeBytes()
	// if err != nil {
	// 	logger.Errorf("Invalid event payload")
	// }
	// agent, err := crypto.GetSignerECC(&d, &event.Payload.Signature)
	// if err != nil {
	// 	logger.Errorf("Invalid event payload")
	// }
	// data.Event = *entities.NewEventPath(event.Validator, entities.TopicModel, event.Hash)
	// data.Agent = entities.AddressFromString(agent).ToDeviceString()
	// data.Account = event.Payload.Account
	// // logger.Error("data.Public ", data.Public)
	// var newState *models.TopicState
	// if updateState {
	// 	newState, _, err = query.SaveRecord(models.TopicState{
	// 		Topic: entities.Topic{ID: data.ID},
	// 	}, &models.TopicState{
	// 		Topic: *data,
	// 	}, utils.IfThenElse(event.EventType == uint16(constants.UpdateTopicEvent), &models.TopicState{
	// 		Topic: *data,
	// 	}, nil) , tx)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		logger.Error("7000: Db Error", err)
	// 		return
	// 	}
	// }
	// tx.Commit()
	// if markAsSynced {
	// 	go OnFinishProcessingEvent(ctx, &data.Event, utils.IfThenElse(newState!=nil, &newState.ID, nil), utils.IfThenElse(event.Error!="", apperror.Internal(event.Error), nil))
	// }

	// if string(event.Validator) != (*cfg).PublicKey  {
	// 	dependent, err := query.GetDependentEvents(*event)
	// 	if err != nil {
	// 		logger.Info("Unable to get dependent events", err)
	// 	}
	// 	for _, dep := range *dependent {
	// 		go HandleNewPubSubTopicEvent(&dep, ctx)
	// 	}
	// }

	// TODO Broadcast the updated state
}
