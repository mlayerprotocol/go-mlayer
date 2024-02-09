package p2p

import (

	// "github.com/gin-gonic/gin"
	"context"
	"reflect"
	"time"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	// rest "messagingprotocol/pkg/core/rest"
	// dhtConfig "github.com/libp2p/go-libp2p-kad-dht/internal/config"
)

/**

**/
func isChannelClosed(ch interface{}) bool {
	// Reflect on the channel to check its state
	c := reflect.ValueOf(ch)
	if c.Kind() != reflect.Chan {
		return false
	}
	_, ok := c.TryRecv()
	return !ok
}

/***
Publish Events to a specified p2p broadcast channel
*****/
func PublishEvent(channelPool chan *entities.Event, pubsubChannel *Channel, mainCtx *context.Context) {
	_, cancel := context.WithCancel(*mainCtx)
	cfg, ok := (*mainCtx).Value(constants.ConfigKey).(*configs.MainConfiguration)
	defer cancel()
	if !ok {
		logger.Fatalf("Unable to read config")
		return
	}
	for {
			event, ok := <-channelPool
			if !ok {
				logger.Fatalf("Channel pool closed. %v", &channelPool)
				return
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
			// 	logger.Infof("ADEDEEDDD %v", b)
				err := pubsubChannel.Publish(entities.NewPubSubMessage(pack))
				if err != nil {
					logger.Errorf("Unable to publish message. Please restart server to try again or adjust buffer size in config. Failed with error %v", err)
					return
				}
			}
	}
	// for {
	// 	select {
	// 	case outAuthorization, ok := <-channelpool.IncomingAuthorizationEventInternal_PubSubC:
	// 		if cfg.Validator {
	// 			if !ok {
	// 				logger.Errorf("Outgoing channel closed. Please restart server to try or adjust buffer size in config")
	// 				return
	// 			}
	// 			err := pubsubChannel.Publish(entities.NewPubSubMessage(outAuthorization.MsgPack()))
	// 			if err != nil {
	// 				logger.Errorf("Failed to publish message. Please restart server to try or adjust buffer size in config")
	// 				return
	// 			}
	// 		}
	// 	case outMessage, ok := <-channelpool.NewPayload_Cli_D_C:
	// 		if cfg.Validator {
	// 			if !ok {
	// 				logger.Errorf("Outgoing channel closed. Please restart server to try or adjust buffer size in config")
	// 				return
	// 			}
	// 			err := messagePubSub.Publish(entities.NewPubSubMessage(outMessage.MsgPack(), cfg.NetworkPrivateKey))
	// 			if err != nil {
	// 				logger.Errorf("Failed to publish message. Please restart server to try or adjust buffer size in config")
	// 				return
	// 			}
	// 		}
	// 	case subscription, ok := <-*subscriptionC:
	// 		if cfg.Validator {
	// 			if !ok {
	// 				logger.Errorf("Subscription channel not found in the context")
	// 				return
	// 			}
	// 			logger.Info("subscription channel:::", subscription.TopicId)

	// 			err := subscriptionPubSub.Publish(entities.NewPubSubMessage(subscription.MsgPack(), cfg.NetworkPrivateKey))
	// 			if err != nil {
	// 				logger.Errorf("Failed to publish subscription.")
	// 				return
	// 			}
	// 		}
	// 	case block, ok := <-*outgoingDPBlockCh:
	// 		if cfg.Validator {
	// 			if !ok {
	// 				logger.Errorf("Subscription channel not found in the context")
	// 				return
	// 			}
	// 			logger.Info("subscription channel:::", block.BlockId)
	// 			err := batchPubSub.Publish(entities.NewPubSubMessage(block.MsgPack(), cfg.NetworkPrivateKey))
	// 			if err != nil {
	// 				logger.Errorf("Failed to publish subscription.")
	// 				return
	// 			}
	// 		}
	// 	}
	// }
}

func ReceiveEvent[PayloadData any](payload *PayloadData, toGoChannel chan *entities.Event, fromPubSubChannel *Channel, mainCtx *context.Context) {
	time.Sleep(5 * time.Second)
	_, cancel := context.WithCancel(*mainCtx)
	defer cancel()
	for {
		message, ok := <-fromPubSubChannel.Messages
		if !ok {
			logger.Fatalf("Primary Message channel closed. Please restart server to try or adjust buffer size in config")
			return
		}
		
		event, errT := entities.UnpackEvent(message.Data,  &entities.Authorization{})
		if errT != nil {
			logger.Errorf("Error receiving event  %v\n", errT)
			continue;
		}
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
		toGoChannel <- event
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