package service

import (
	"context"
	"fmt"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/p2p"
	"gorm.io/gorm"
)


type PayloadData struct {
	Subnet string
	localDataState *LocalDataState
	
	localDataStateEvent *LocalDataStateEvent
}



type LocalDataState struct {
	Hash string
	ID string
	Event *entities.EventPath
	Timestamp uint64
}
type LocalDataStateEvent struct {
	Hash string
	ID string
	Timestamp uint64
}

func ProcessEvent(event *entities.Event, data PayloadData, validAgentRequired bool, saveEvent func (entities.Event, *entities.Event, *entities.Event, *gorm.DB) (*entities.Event, error), tx *gorm.DB, ctx *context.Context) (bool, bool, *models.AuthorizationState, bool, error) {
	cfg, _ := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)
	// logger.WithFields(logrus.Fields{"event": event}).Debug("New topic event from pubsub channel")
	// markAsSynced := false
	// updateState := false
	// var eventError string
	// // hash, _ := event.GetHash()

	logger.Infof("Event is a valid event %s", event.PayloadHash)
	//cfg, _ := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)

	// Extract and validate the Data of the paylaod which is an Events Payload Data,
	
	logger.Infof("NEWTOPICEVENT: %s", event.Hash)
	previousEventUptoDate := false
	authEventUptoDate :=  false
	eventIsMoreRecent := true
	authMoreRecent := false
	var badEvent  error
	// eventIsMoreRecent := true
	// authMoreRecent := false

	err := ValidateEvent(*event)

	if err != nil {
		logger.Debug(err)
		return false, false, nil, eventIsMoreRecent, err
	}
	d, err := event.Payload.EncodeBytes()
	if err != nil || len(d) == 0 {
		logger.Debug("Invalid event payload")
		return false, false, nil, eventIsMoreRecent, fmt.Errorf("invalid event payload")
	}
	
	subnet := models.SubnetState{}
	
	if event.EventType != uint16(constants.CreateSubnetEvent) && event.EventType != uint16(constants.UpdateSubnetEvent) {
	err = query.GetOne(models.SubnetState{Subnet: entities.Subnet{ID: data.Subnet}}, &subnet)
	if err != nil {
		logger.Infof("EVENTINFO: %v %s", err, data.Subnet)
		if err == gorm.ErrRecordNotFound {
			// get the subnetstate from the sending node
			subPath := entities.NewEntityPath(event.Validator, entities.SubnetModel, data.Subnet)
			pp, err := p2p.GetState(cfg, *subPath, &subnet)
			if err != nil {
				logger.Error(err)
				return false, false, nil, eventIsMoreRecent, fmt.Errorf("unable to get subnetdata")
			}
			subnetEvent, err := entities.UnpackEvent(pp.Event, entities.Subnet{})
			if err != nil {
				logger.Error(err)
				return false, false, nil, eventIsMoreRecent, fmt.Errorf("Unable to unpack subnetdata")
			}
			if subnetEvent != nil {
				HandleNewPubSubSubnetEvent(subnetEvent, ctx)
			}
			// save it
			// query.SaveRecord(models.SubnetState{Subnet: entities.Subnet{ID: data.Subnet}}, &subnet, nil, nil )
			// if err != nil {
			// 	return
			// }
		}
	
	}
	}
	var agent string
	if validAgentRequired {
		agent, err = crypto.GetSignerECC(&d, &event.Payload.Signature)
		if err != nil {
			logger.Debug(err)
			return false, false, nil, eventIsMoreRecent, fmt.Errorf("invalid agent signature")
		}
	
		if entities.StringToDeviceString(string(event.Payload.Agent)) != entities.StringToDeviceString(agent) {
			logger.Debug("Invalid agent signer")
			return false, false, nil, eventIsMoreRecent, fmt.Errorf("invalid payload signer")
		}
	}
	// eventModel, _, err := query.SaveRecord(models.TopicEvent{Event: entities.Event{Hash: event.Hash}}, &models.TopicEvent{Event: *event}, nil, sql.SqlDb)
	_, err = saveEvent( entities.Event{Hash: event.Hash},  event, nil, tx)
	if err != nil {
		return false, false,nil, eventIsMoreRecent, fmt.Errorf("event storage failed")
	}

	
	
	// get agent auth state
	
	var eventAuthState *models.AuthorizationState
	var agentAuthState *models.AuthorizationState
	var agentAuthStateEvent *models.AuthorizationEvent
	if validAgentRequired {
		err = query.GetOne(models.AuthorizationState{Authorization: entities.Authorization{Agent: entities.StringToDeviceString(agent), Subnet: event.Payload.Subnet}}, agentAuthState)
		if err != nil && err != query.ErrorNotFound {
			return false, false,nil, eventIsMoreRecent, fmt.Errorf("db error: %s", err.Error())
		}
	}
	// lets determine which authstate to use to validate this event
	

	// get all events and auth
	var previousEvent *entities.Event
	var authEvent *entities.Event

	// if agentAuthState == nil { // we dont have any info about the agent within this subnet
	// 	// we need to know if the agent has the right to process this event, else we cant do anything
	// 	// check the node that sent the event to see if it has the record

	// }

	if len(event.PreviousEventHash.Hash) > 0 {
		previousEvent, err = query.GetEventFromPath(&event.PreviousEventHash)
		if err != nil && err != query.ErrorNotFound {
			logger.Debug(err)
			return false, false, nil, eventIsMoreRecent, fmt.Errorf("db err: %s", err.Error())
		}
		// check if we have the previous event locally, if we dont we can't proceed until we get it
		if previousEvent != nil {
			previousEventUptoDate = true
		} else {
			// get the previous event from the sending node and process it as well
			previousEvent, _, err = p2p.GetEvent(cfg, event.PreviousEventHash, nil)
			if err != nil {
				logger.Error(err)
				if event.Validator != event.PreviousEventHash.Validator {
					previousEvent, _, err = p2p.GetEvent(cfg, event.PreviousEventHash, &event.Validator)
					logger.Error(err)
				}
			}
			if previousEvent != nil {
				go HandleNewPubSubTopicEvent(previousEvent, ctx)
			}
			
		}

	} else {
		previousEventUptoDate = true
	}

	
		if len(event.AuthEventHash.Hash) > 0 && (agentAuthState == nil || agentAuthState.Event.Hash != event.AuthEventHash.Hash) {
			// check if we have the associated auth event locally, if we dont we can't proceed until we get it
			authEvent, err = query.GetEventFromPath(&event.AuthEventHash)
			if err != nil && err != query.ErrorNotFound {
				logger.Debug(err)
				return previousEventUptoDate, authEventUptoDate, nil, eventIsMoreRecent, fmt.Errorf("auth event not found")
			} else {
				if authEvent != nil {
					authEventUptoDate = true
				} else {
					// get the event from the sending node and process it and process it
					authEvent, _, err = p2p.GetEvent(cfg, event.AuthEventHash, nil)
					if err != nil {
						logger.Error(err)
						if event.Validator != event.AuthEventHash.Validator {
							authEvent, _, err = p2p.GetEvent(cfg, event.AuthEventHash, &event.Validator)
							logger.Error(err)
						}
					}
					if authEvent != nil {

						go HandleNewPubSubAuthEvent(authEvent, ctx)
					}
				}
			}
			
		
	}
	if previousEventUptoDate &&  authEventUptoDate {
		// check a situation where we have either of current auth and event state locally, but the states events are not same as the events prev auth and event
		// get the topics state
		// var id = data.ID
		// var badEvent  error
		// if len(data.ID) == 0 {
		// 	id, _ = entities.GetId(data)
		// } else {
		// 	id = data.ID
		// }
		
		var topicState = data.localDataState
		var topicEvent = data.localDataStateEvent
		// err := query.GetOne(models.TopicState{Topic: entities.Topic{ID: id}}, &topic)
		// if err != nil && err != query.ErrorNotFound {
		// 	logger.Debug(err)
		// }
		if topicState != nil &&  len(topicState.ID) > 0 {
			// check if state.Event is same as events previous has
			if topicState.Event.Hash != event.PreviousEventHash.Hash {
				// either we are not upto date, or the sender is not
				// get the event that resulted in current state
				// topicEvent, err = query.GetEventFromPath(&topicState.Event)
				// if err != nil && err != query.ErrorNotFound {
				// 	logger.Debug(err)
				// }
				if len(topicEvent.ID) > 0 {
					eventIsMoreRecent = IsMoreRecentEvent(topicEvent.Hash, int(topicEvent.Timestamp), event.Hash, int(event.Timestamp))


					 // if this event is more recent, then it must referrence our local event or an event after it
					if eventIsMoreRecent  && topicEvent.Hash != event.PreviousEventHash.Hash {
						previousEventMoreRecent := IsMoreRecentEvent(topicEvent.Hash, int(topicEvent.Timestamp), previousEvent.Hash, int(previousEvent.Timestamp))
						if !previousEventMoreRecent {
							badEvent = fmt.Errorf(constants.ErrorBadRequest)
						}
					}

				}
				
			}
		}

		





		// We need to determin which authstate is valid for this event
		// if its an old event, we can just save it since its not updating state
		// if its a new one, we have to confirm that it is referencing the true latest auth event

		// so lets get the referrenced authorization
		
		err = query.GetOne(models.AuthorizationState{Authorization: entities.Authorization{Event: event.AuthEventHash}}, eventAuthState)
		if err != nil && err != query.ErrorNotFound {
			return false, false, nil, false,fmt.Errorf("db error: %s", err.Error())
		}
		// if event is more recent that our local state, we have to check its validity since it updates state
		if eventIsMoreRecent && validAgentRequired && agentAuthState != nil && agentAuthState.Event.Hash != event.AuthEventHash.Hash && authEvent != nil {
			// get the event that is responsible for the current state
			err := query.GetOne(models.AuthorizationEvent{Event: entities.Event{Hash: agentAuthState.Event.Hash}}, &agentAuthStateEvent)
			if err != nil && err != query.ErrorNotFound {
				logger.Debug(err)
			}
			if len(agentAuthStateEvent.ID) > 0 {
				authMoreRecent = IsMoreRecentEvent(agentAuthStateEvent.Hash, int(agentAuthStateEvent.Timestamp), authEvent.Hash, int(authEvent.Timestamp))
				if !authMoreRecent {
					// this is a bad event using an old auth state.
					// REJECT IT
					badEvent = fmt.Errorf(constants.ErrorUnauthorized)
				}
			}
		}



		if badEvent != nil {
			// update the event state with the error
			// _, _, err = query.SaveRecord(models.TopicEvent{Event: entities.Event{Hash: event.Hash}}, &models.TopicEvent{Event: *event}, &models.TopicEvent{Event: entities.Event{Error: badEvent.Error(), IsValid: false, Synced: true}}, nil)

			_, err = saveEvent(entities.Event{Hash: event.Hash}, event,  &entities.Event{Error: badEvent.Error(), IsValid: false, Synced: true}, tx)
			if err != nil {
				logger.Error(err)
			}
			// notify the originator so it can correct it e.g. let it know that there is a new authorization

			// decide whether to report the node
			return false, false, eventAuthState, eventIsMoreRecent, fmt.Errorf("db error: %s", badEvent)
		}
		
	}
	return previousEventUptoDate, authEventUptoDate, eventAuthState, eventIsMoreRecent, nil

}