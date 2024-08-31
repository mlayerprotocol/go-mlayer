package p2p

import (

	// "github.com/gin-gonic/gin"
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ipfs/go-datastore"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
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

func ProcessEventsReceivedFromOtherNodes(modelType entities.EntityModel, fromPubSubChannel *entities.Channel, mainCtx *context.Context, process func(event *entities.Event, ctx *context.Context)) {
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
		
		event, errT := entities.UnpackEvent(message.Data,  modelType)
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
func processP2pPayload(config *configs.MainConfiguration, payload *P2pPayload) (response *P2pPayload, err error) {
	ctx := MainContext
	
	response = NewP2pPayload(config, P2pActionResponse, []byte{})
	response.Id = payload.Id
	claimedRewardStore, ok := (*ctx).Value(constants.ClaimedRewardStore).(*db.Datastore)
	if !ok {
		response.ResponseCode = 500
		response.Error = "Internal error"
		logger.Infof("CommitmentRequest: Error get claim reward store from context")
		response.Sign(config.PrivateKeyEDD)
		return response, err
	}
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
				logger.Errorf("EventFromPathError: %v,%v", err, eventPath)
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
				err = sql.SqlDb.Model(d).Where("event = ?", eventPath.ToString(), &result).Error

				if err != nil {
					logger.Errorf("EventReseponse: %v", err)
				}
				if err == nil {
					for _, st := range result {
						states = append(states, st.MsgPack())
					}
					data := P2pEventResponse{Event: event.MsgPack(), States: states}
					logger.Infof("EventReseponse: %v", (&data).MsgPack())
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
				eventPath := d.FieldByName("Event").Interface()
				// := entities.EventPathFromString(eventPath)
				path := eventPath.(entities.EventPath)
				event, err := query.GetEventFromPath(&path)
				logger.Infof("STATESEVENNNT %v, %v", event, eventPath)
				if err == nil {
					states := []json.RawMessage{}
					states = append(states, state.(IState).MsgPack())
					data := P2pEventResponse{Event: event.MsgPack(), States: states}
					response.Data = (&data).MsgPack()
				} else {
					response.ResponseCode = 500
					response.Error = err.Error()
				}
				
			
				
				
			}
		case P2pActionSyncBlock:
			
			blocks := [][]byte{}
			encoder.MsgPackUnpackStruct(payload.Data, &blocks)
			// cycleKey :=  fmt.Sprintf("%s/%d", response.Signer, batch.Cycle)
	
			// subnetList, err := eventCounterStore.Query(*ctx, dsQuery.Query{
			// 	Prefix: cycleKey,
			// })
			
			var buffer bytes.Buffer
			models := []any{
				models.SubnetEvent{},
				models.AuthorizationEvent{},
				models.TopicEvent{},
				models.TopicEvent{},
				models.SubscriptionEvent{},
				models.MessageEvent{},
			}
			fromBlock :=  new(big.Int).SetBytes(blocks[0])
			toBlock :=  new(big.Int).SetBytes(blocks[1])
			for _, m := range models {
				// get events from each cycle
				// block := new(big.Int).SetBytes(blockByte)
				b, err := generateImportScript(m, fromBlock.Uint64(), toBlock.Uint64())
				if err != nil {
					response.ResponseCode = 404
					response.Error = "Event not found"
					break
				}
				buffer.Write(b)
				buffer.Write([]byte{'|'})
			}
				response.Data = buffer.Bytes()
			//	logger.Infof("FILEPATH: %s", sql)
			
		case P2pActionGetCommitment:
			// logger.Info("ReceivedCommitmentRequest")
			// eventCounterStore, ok := (*ctx).Value(constants.EventCountStore).(*db.Datastore)
			// if !ok {
			// 	panic("Unable to load eventCounterStore")
			// }
			// TODO check if you own any of the next 20? licenses that should validate this, if you dont, no need to commit
			
			realBatch, err := entities.UnpackRewardBatch(payload.Data)
			batchCopy := *realBatch
			batch := &batchCopy
			batch.Clear()
			if err != nil {
				response.ResponseCode = 500
				response.Error = err.Error()
			}
			//  cycleKey :=  fmt.Sprintf("%s/%d", response.Signer, batch.Cycle)
			subnetList := []models.EventCounter{}
			claimed := false
			err = query.GetManyWithLimit(models.EventCounter{Cycle: &batch.Cycle, Validator: entities.PublicKeyString(hex.EncodeToString(payload.Signer)), Claimed: &claimed }, &subnetList, &map[string]query.Order{"count": query.OrderDec}, entities.MaxBatchSize,  batch.Index * entities.MaxBatchSize)
			if err != nil {
				return nil, err
			}
			if len(subnetList) == 0 {
				response.ResponseCode = 500
				response.Error = "empty list"
				break
			}
			if subnetList[0].Subnet != realBatch.DataBoundary[0].Subnet {
				response.ResponseCode = 500
				response.Error = "upper data boundary dont match"
				break
			}
			if subnetList[len(subnetList)-1].Subnet != realBatch.DataBoundary[1].Subnet {
				response.ResponseCode = 500
				response.Error = "lower data boundary dont match"
				break
			}
			
			for  _, rsl := range subnetList {
				// if start  == i {
					// if rsl.Subnet != batch.DataBoundary[0].Subnet {
					// 	response.ResponseCode = 500
					// 	response.Error = "data boundary dont match"
					// 	break
					// }
					batch.Append(entities.SubnetCount{
						Subnet: rsl.Subnet,
						EventCount: *rsl.Count,
					})
					
				// }
				// if i > start + 99 {
				// 	break
				// }
				// i++
			}
			
			claimHash := [32]byte{}
			if len(batch.Data) > 0 && len(response.Error) == 0  {
				//logger.Infof("BATCHINGOF %s", realBatch.GetProofData(config.ChainId).DataHash)
				claimHash, err = realBatch.GetProofData(config.ChainId).GetHash()
				logger.Infof("ValidDataHash %v, %v",[32]byte(batch.DataHash) == [32]byte(realBatch.DataHash), realBatch )
				if err != nil {
					response.ResponseCode = 500
					response.Error = err.Error()
					logger.Errorf("Error getting hash: %v", err)
				}
				if [32]byte(batch.DataHash) != [32]byte(realBatch.DataHash) {
					response.ResponseCode = 400
					response.Error = "Invalid batch hash"
				}
			} else {
				response.ResponseCode = 400
				response.Error = "Invalid batch hash"
			}

			
			validCommitmentKey :=  datastore.NewKey(fmt.Sprintf("commitment/%s",  hex.EncodeToString(claimHash[:])))
			logger.Infof("CommitmentKey1: %s", validCommitmentKey.String())

			if response.ResponseCode == 0 {
				pk, _ := btcec.PrivKeyFromBytes(config.PrivateKeySECP)
				nonce, noncePublicKey := schnorr.ComputeNonce(pk, claimHash)
				err = claimedRewardStore.Put(*ctx, validCommitmentKey, nonce.Bytes())
				if err != nil {
					logger.Errorf("FailedStoringComittemnt: %v", err)
					response.ResponseCode = 500
					response.Error = "Internal error"
				} else{ 
					response.Data = noncePublicKey.SerializeCompressed()
				}
				logger.Infof("NoncePubKey %s", hex.EncodeToString(noncePublicKey.SerializeCompressed()))
			}
	
		case P2pActionGetSentryProof:
			logger.Info("ReceivedProoftRequest")
			// eventCounterStore, ok := (*ctx).Value(constants.EventCountStore).(*db.Datastore)
			// if !ok {
			// 	panic("Unable to load eventCounterStore")
			// }
			
			sigData, err := entities.UnpackSignatureRequestData(payload.Data)
			
			if err != nil {
				response.ResponseCode = 500
				response.Error = err.Error()
			}
			
			validCommitmentKey :=  datastore.NewKey(fmt.Sprintf("commitment/%s",  hex.EncodeToString(sigData.ProofHash)))
			logger.Infof("CommitmentKey2: %s", validCommitmentKey.String())
			
			nonce, err := claimedRewardStore.Get(*ctx, validCommitmentKey)
			if err != nil {
				response.ResponseCode = 500
				response.Error = "Internal error"
				logger.Infof("Error getting commitment from store")
			} 
			if err == nil && response.ResponseCode == 0 {

				pk, _ := btcec.PrivKeyFromBytes(config.PrivateKeySECP)
				// nonce, _ := schnorr.ComputeNonce(pk, [32]byte(sigData.BatchHash))
				sig := schnorr.ComputeSignature(pk, new(big.Int).SetBytes(nonce), sigData.Challenge)
				//  cycleKey :=  fmt.Sprintf("%s/%d", response.Signer, batch.Cycle)
				response.Data = sig
				/// TODO save the nonepublickey with the claimhash in badger
				logger.Infof("NoncePubKey %s", hex.EncodeToString(sig))
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
			
		default:
			response.Error = "invalid action type"
			response.ResponseCode = 400
			
	}
	response.Sign(config.PrivateKeyEDD)
	return response, err
}

func generateImportScript(model any, fromBlock uint64, toBlock uint64) ([]byte, error) {

	sql, err := query.GenerateImportScript(sql.SqlDb, model, sql.SqlDb.Where("block_number >= ? AND block_number <= ?",  fromBlock, toBlock), "", config )
				if err != nil {
					logger.Infof("SQLFILEERROR: %v", err)
				}
				d, err := utils.CompressToGzip(sql)
				if err != nil {
					return nil, err
				}
				return d, nil
}
