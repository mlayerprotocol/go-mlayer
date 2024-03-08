package message

import (
	"context"
	"sync"

	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/channelpool"
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/db"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
)

var logger = &log.Logger
func ProcessNewMessageEvent(ctx context.Context, unsentMessageP2pStore *db.Datastore,  wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {

			// attempt to push into outgoing message channel
			case outMessage, ok := <-channelpool.NewPayload_Cli_D_C:
				if !ok {
					logger.Errorf("Outgoing Message channel closed. Please restart server to try or adjust buffer size in config")

					return
				}
				go ProcessSentMessage(ctx, unsentMessageP2pStore, outMessage)

			// case sub, ok := <-channelpool.SubscribersRPC_D_c:
			// 	if !ok {
			// 		logger.Errorf("Subscription channel closed!")
			// 		return
			// 	}
			// 	if !entities.IsValidSubscription(*sub, true) {
			// 		logger.Info("ITS NOT VALID!")
			// 		continue
			// 	}
			// 	// go processor.ProcessNewSubscription(ctx, sub, channelsubscribersRPC_D_countStore, channelSubscriptionStore)

			// case clientHandshake, ok := <-channelpool.ClientHandshakeC:
			// 	if !ok {
			// 		logger.Errorf("Verification channel closed. Please restart server to try or adjust buffer size in config")
			// 		wg.Done()
			// 		return
			// 	}
			// 	go processor.ValidateMessageClient(ctx, &connectedSubscribers, clientHandshake, channelSubscriptionStore)

			// case proof, ok := <-channelpool.IncomingDeliveryProofsC:
			// 	if !ok {
			// 		logger.Errorf("Incoming delivery proof channel closed. Please restart server to try or adjust buffer size in config")
			// 		wg.Done()
			// 		return
			// 	}
			// 	go processor.ValidateAndAddToDeliveryProofToBlock(ctx,
			// 		proof,
			// 		deliveryProofStore,
			// 		channelSubscriptionStore,
			// 		deliveryProofBlockStateStore,
			// 		localDPBlockStore,
			// 		MaxDeliveryProofBlockSize,
			// 		&deliveryProofBlockMutex,
			// 	)

			// case batch, ok := <-channelpool.PubSubInputBlockC:
			// 	if !ok {
			// 		logger.Errorf("PubsubInputBlock channel closed. Please restart server to try or adjust buffer size in config")
			// 		wg.Done()
			// 		return
			// 	}
			// 	go func() {
			// 		unconfurmedBlockStore.Put(ctx, db.Key(batch.Key()), batch.MsgPack())
			// 	}()
			// case proof, ok := <-channelpool.PubSubInputProofC:
			// 	if !ok {
			// 		logger.Errorf("PubsubInputBlock channel closed. Please restart server to try or adjust buffer size in config")
			// 		wg.Done()
			// 		return
			// 	}
			// 	go func() {
			// 		unconfurmedBlockStore.Put(ctx, db.Key(proof.BlockKey()), proof.MsgPack())
			// 	}()

			 }

		}
	
}
func ProcessSentMessage(ctx context.Context, unsentMessageP2pStore *db.Datastore, outMessage *entities.ClientPayload) {
	// VALIDATE AND DISTRIBUTE
	// outMessage := (outEvent.Data).(*entities.Message)
	// construct the event
	// newMessageEvent := entities.Event{}
	//d, _:= (outMessage.Data).GetHash()
	// logger.Infof("\nSending out message %s\n", d)
	// unsentMessageP2pStore.Set(ctx, db.Key(outMessage.Key()), outMessage.MsgPack(), false)
	// channelpool.OutgoingMessageEvents_D_P2P_C <- outEvent
	// channelpool.IncomingMessageEvent_P2P_D_C <- outEvent
	logger.Infof("\nSending out complete\n")
}
