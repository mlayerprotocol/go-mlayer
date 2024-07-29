package p2p

import (

	// "github.com/gin-gonic/gin"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	dsQuery "github.com/ipfs/go-datastore/query"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto/schnorr"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/db"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	// rest "messagingprotocol/pkg/core/rest"
	// dhtConfig "github.com/libp2p/go-libp2p-kad-dht/internal/config"
)

/**

**/
// func isChannelClosed(ch interface{}) bool {
// 	// Reflect on the channel to check its state
// 	c := reflect.ValueOf(ch)
// 	if c.Kind() != reflect.Chan {
// 		return false
// 	}
// 	_, ok := c.TryRecv()
// 	return !ok
// }

/***
Publish Events to a specified p2p broadcast channel
*****/
func publishChannelEventToNetwork(channelPool chan *entities.Event, pubsubChannel *entities.Channel, mainCtx *context.Context) {
	_, cancel := context.WithCancel(*mainCtx)
	cfg, ok := (*mainCtx).Value(constants.ConfigKey).(*configs.MainConfiguration)
	
	defer cancel()
	if !ok {
		logger.Fatalf("Unable to read config")
		return
	}
	for {
		if pubsubChannel == nil {
			continue
		}
			event, ok := <-channelPool
			
			
			if !ok {
				logger.Fatalf("Channel pool closed. %v", &channelPool)
				panic("Channel pool closed")
			}
			if cfg.Validator {
				if !ok {
					logger.Errorf("Outgoing channel closed. Please restart server to try or adjust buffer size in config")
					return
				}
				pack := event.MsgPack()
				// if err != nil {
				// 	logger.Error(err)
				// 	continue
				// }
				
		
				// event, errT := entities.UnpackEvent(pack, &entities.Authorization{})
				// if errT != nil {
				// 	logger.Errorf("Error receiving event  %v\n", errT)
				// 	continue;
				// }
				
				// eT := entities.Event{
				// 	Payload: ,
				// }
				// // auth := models.AuthorizationEvent{}
				// err = entities.UnpackEvent(pack, &eT)
				// if err != nil {
				// 	logger.Errorf("Failed to UNPCAKC %v", err)
				// 	continue
				// }
				//payload := entities.AuthorizationPayload{}
				// dbByte, _ := json.Marshal(auth.Event.Payload)
				// _	= json.Unmarshal(dbByte, &payload)
				
			 // auth.Payload = payload
			//  auth.Payload.ClientPayload.Data = auth.Payload.ClientPayload.Data
			//  logger.Infof("Payload----> %v", event.Payload.ClientPayload.Data)
			// 	// auth.Event.Payload = payload
			// 	b, err := (auth).EncodeBytes()
			// 	if err != nil {
			// 		logger.Errorf("Failed to ENCODE %v", err)
			// 		continue
			// 	}
			
				err := pubsubChannel.Publish(entities.NewPubSubMessage(pack))
				
				if err != nil {
					logger.Errorf("Unable to publish message. Please restart server to try again or adjust buffer size in config. Failed with error %v", err)
					return
				}
				
			}
	}
	
}

func ProcessEventsReceivedFromOtherNodes[PayloadData any](payload *PayloadData, fromPubSubChannel *entities.Channel, mainCtx *context.Context, process func(event *entities.Event, ctx *context.Context)) {
	// time.Sleep(5 * time.Second)
	
	_, cancel := context.WithCancel(*mainCtx)
	
	defer cancel()
	
	for {
		if fromPubSubChannel == nil || fromPubSubChannel.Messages == nil {
			logger.Info("Channel is nil")
			time.Sleep(1 * time.Second)
			continue
		}
		logger.Info("Channel no more nil")
		
		message, ok := <-fromPubSubChannel.Messages
		if !ok {
			logger.Fatalf("Primary Message channel closed. Please restart server to try or adjust buffer size in config")
			return
		}
		
		event, errT := entities.UnpackEvent(message.Data,  *payload)
		logger.Infof("UNPASCKEDEVENT %s  %s, %v", event.PreviousEventHash, event.PayloadHash,  event.Payload.Data)
		
		if errT != nil {
			logger.Errorf("Error receiving event  %v\n", errT)
			continue;
		}

		// TODO validate the event
		// pv := models.AuthorizationEvent{
		// 	Event: event,
		// }
		// auth := entities.Authorization{}
		// err := encoder.MsgPackUnpackStruct(message.Data, &event)
		// if err != nil {
		// 	logger.Error(err)
		// }
		// logger.Infof("Event 1 ----===> %v", event)
		// authByte, _ := json.Marshal( pv.Event.Payload.(entities.AuthorizationPayload).ClientPayload.Data)
		// _	= json.Unmarshal(authByte, &auth)
		// // pv.Event.Payload  = payload
		// pv.Event.ClientPayload.Data = payload.Data
		// pv.Payload.ClientPayload = entities.AuthorizationPayload{
		// 	ClientPayload: payload.ClientPayload,
		// }
		
		// authEvent := models.AuthorizationEvent{
		// 	Event: entities.Event{
		// 		Payload: payload,
		// 	},
		// 	Payload: entities.AuthorizationPayload{
		// 		Data: payload.Data.
		// 	},
		// }
		// logger.Infof("ADEDEEDDD %v", pv.Event.Payload.(entities.AuthorizationPayload).ClientPayload)
		// b, err := pv.EncodeBytes()
		// logger.Infof("ADEDEEDDD %v", b)
		// logger.Infof("Event Received ----===> %v", event.GetValidator())
		// toGoChannel <- event
		
		go process(event, mainCtx)
	}
	// for {
	// 	select {

	// 	case authEvent, ok := <-authorizationPubSub.Messages:
	// 		if !ok {
	// 			cancel()
	// 			logger.Fatalf("Primary Message channel closed. Please restart server to try or adjust buffer size in config")
	// 			return
	// 		}
	// 		// !validating message
	// 		// !if not a valid message continue
	// 		// _, err := inMessage.MsgPack()
	// 		// if err != nil {
	// 		// 	continue
	// 		// }
	// 		//TODO:
	// 		// if not a valid message, continue

	// 		logger.Infof("Received new message %s\n", authEvent.ToString())
	// 		cm := models.AuthorizationEvent{}
	// 		err = encoder.MsgPackUnpackStruct(authEvent.Data, cm)
	// 		if err != nil {

	// 		}
	// 		*incomingAuthorizationC <- &cm
	// 	case inMessage, ok := <-batchPubSub.Messages:
	// 		if !ok {
	// 			cancel()
	// 			logger.Fatalf("Primary Message channel closed. Please restart server to try or adjust buffer size in config")
	// 			return
	// 		}
	// 		// !validating message
	// 		// !if not a valid message continue
	// 		// _, err := inMessage.MsgPack()
	// 		// if err != nil {
	// 		// 	continue
	// 		// }
	// 		//TODO:
	// 		// if not a valid message, continue

	// 		logger.Infof("Received new message %s\n", inMessage.ToString())
	// 		cm, err := entities.MsgUnpackClientPayload(inMessage.Data)
	// 		if err != nil {

	// 		}
	// 		*incomingMessagesC <- &cm
	// 	case sub, ok := <-subscriptionPubSub.Messages:
	// 		if !ok {
	// 			cancel()
	// 			logger.Fatalf("Primary Message channel closed. Please restart server to try or adjust buffer size in config")
	// 			return
	// 		}
	// 		// logger.Info("Received new message %s\n", inMessage.Message.Body.Message)
	// 		cm, err := entities.UnpackSubscription(sub.Data)
	// 		if err != nil {

	// 		}
	// 		logger.Info("New subscription updates:::", string(cm.ToJSON()))
	// 		// *incomingMessagesC <- &cm
	// 		cm.Broadcasted = false
	// 		*publishedSubscriptionC <- &cm
	// 	}
	// }
}

type IState interface {
	MsgPack() []byte
}

//process the payload based on the type of request
func processP2pPayload(ctx *context.Context, config *configs.MainConfiguration, payload *P2pPayload) (response *P2pPayload, err error) {
	response = NewP2pPayload(config, P2pActionResponse, []byte{})
	response.Id = payload.Id
	switch(payload.Action) {
	case P2pActionGetEvent:
			eventPath, err := entities.UnpackEventPath(payload.Data)
			if err != nil {
				response.ResponseCode = 500
				response.Error = "Invalid payload data"
				logger.Infof("processP2pPayload: %v", err)
			}
			event, err := query.GetEventFromPath(eventPath)
			if err != nil {
				if err == query.ErrorNotFound {
					response.ResponseCode = 404
					response.Error = "Event not found"
				} else {
				response.ResponseCode = 500
				response.Error = err.Error()
				}
			} else {
				d := models.GetModelFromModelType(eventPath.Model)
				result := []IState{}
				states := []json.RawMessage{}
				// states := query.GetMany(d, &result)
				sql.SqlDb.Model(d).Where("event = ?", eventPath.ToString(), &result)
				for _, st := range result {
					states = append(states, st.MsgPack())
				}
				if err == nil {
					data := P2pEventResponse{Event: event.MsgPack(), States: states}
					response.Data = (&data).MsgPack()
				}
			}
		case P2pActionGetState:
			ePath, err := entities.UnpackEntityPath(payload.Data)
			if err != nil {
				response.ResponseCode = 500
				response.Error = "Invalid payload data"
				logger.Infof("processP2pPayload: %v", err)
			}
			state, err := query.GetStateFromPath(ePath)
			if err != nil {
				if err == query.ErrorNotFound {
					response.ResponseCode = 404
					response.Error = "Event not found"
				} else {
				response.ResponseCode = 500
				response.Error = err.Error()
				}
			} else {
				d := reflect.ValueOf(state).Elem()
				eventPath := fmt.Sprint(d.FieldByName("Event").Interface())
				pathFromString := entities.EventPathFromString(eventPath)
				event, err := query.GetEventFromPath(pathFromString)
				states := []json.RawMessage{}
				states = append(states, state.(IState).MsgPack())
				if err == nil {
					data := P2pEventResponse{Event: event.MsgPack(), States: states}
					response.Data = (&data).MsgPack()
				}
			}

		case P2pActionGetCommitment:
			
			eventCounterStore, ok := (*ctx).Value(constants.EventCountStore).(*db.Datastore)
			if !ok {
				panic("Unable to load eventCounterStore")
			}
			realBatch, err := entities.UnpackRewardBatch(payload.Data)
			batch := realBatch
			if err != nil {
				response.ResponseCode = 500
				response.Error = err.Error()
			}
			cycleKey :=  fmt.Sprintf("%s/%d", response.Signer, batch.Cycle)
	
			subnetList, err := eventCounterStore.Query(*ctx, dsQuery.Query{
				Prefix: cycleKey,
			})
			defer subnetList.Close()
			i := uint64(0)
			start := batch.Index * 100
			for  rsl := range subnetList.Next() {
				if start  == i {
					if rsl.Key != batch.DataBoundary[0].Subnet {
						response.ResponseCode = 500
						response.Error = err.Error()
						break
					}
					batch.Append(entities.SubnetCount{
						Subnet: rsl.Key,
						EventCount: encoder.NumberFromByte(rsl.Value),
					})
					
				}
				if i > start + 99 {
					break
				}
				i++
			}
			claimHash := [32]byte{}
			if len(batch.Data) > 0 && len(response.Error) == 0  {
				claimHash, err = batch.GetHash(config.ChainId)
				if err != nil {
					response.ResponseCode = 500
					response.Error = err.Error()
				}
				if [32]byte(batch.DataHash) != [32]byte(realBatch.DataHash) {
					response.ResponseCode = 400
					response.Error = "Invalid batch hash"
				}
			} else {
				response.ResponseCode = 400
				response.Error = "Invalid batch hash"
			}

			if response.ResponseCode == 0 {
				pk, _ := btcec.PrivKeyFromBytes(config.PrivateKeyBytes)
				_, noncePublicKey := schnorr.ComputeNonce(pk, claimHash)
				response.Data = noncePublicKey.SerializeCompressed()
				/// TODO save the nonepublickey with the claimhash in badger
			}
			

			// if err != nil {
			// 	response.ResponseCode = 500
			// 	response.Error = "Invalid payload data"
			// 	logger.Infof("processP2pPayload: %v", err)
			// }
			
			// 1. Get the reward batch data
			// 2. Loop through the Data field and check your /validator/cycle/subnetId/{batchId} to get the last time a proof was requested
			// 3. If this is less than 10 minutes ago, respond with error - proof requested too early
			// 4. If non exists or most recent is more than 10 minutes
			
			
			
	}
	response.Sign(config.PrivateKeyBytes)
	return response, err
}