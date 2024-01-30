package subscription

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ipfs/go-datastore/query"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/chain"
	"github.com/mlayerprotocol/go-mlayer/internal/channelpool"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/db"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
	"github.com/mlayerprotocol/go-mlayer/utils/constants"
)

var logger = &log.Logger

func ProcessNewSubscription(
	ctx context.Context,
	subscriptionBlockStateStore *db.Datastore,
	channelSubscriptionCountStore *db.Datastore,
	newChannelSubscriptionStore *db.Datastore,
	channelSubscriberStore *db.Datastore,
	wg *sync.WaitGroup,
) {

	defer (*wg).Done()
	defer time.Sleep(5 * time.Second)
	for {

		var block *entities.Block
		blockStateTxn, err := subscriptionBlockStateStore.NewTransaction(ctx, false)
		if err != nil {
			logger.Errorf("Subscription Block state store connection error %o", err)
			// invalid proof or proof has been tampered with
			panic(err)
		}

		blockData, err := blockStateTxn.Get(ctx, db.Key(constants.CurrentSubscriptionBlockStateKey))
		if err != nil {
			logger.Debugf("Block state store error - key: %s, %v", constants.CurrentSubscriptionBlockStateKey, err)
			// invalid proof or proof has been tampered with

			if string(err.Error()) == "datastore: key not found" {
				block = entities.NewBlock()
				// logger.Debugf("Error: %s", block.ToJSON())
				blockStateTxn.Put(ctx, db.Key(constants.CurrentDeliveryProofBlockStateKey), block.MsgPack())
				err = blockStateTxn.Commit(ctx)
				if err != nil {
					panic(err)
				}
			}
			continue
		}
		if len(blockData) > 0 {
			block, err = entities.UnpackBlock(blockData)
			if err != nil {
				logger.Errorf("Invalid block data %o", err)
				// invalid proof or proof has been tampered with
				blockStateTxn.Discard(ctx)
				continue
			}
		}

		results, err := newChannelSubscriptionStore.Query(ctx, query.Query{
			Prefix: "/",
		})
		if err != nil {
			logger.Errorf("New Channel Subscriber Store Query Error %o", err)
			return
		}
		for {
			result, ok := <-results.Next()
			if !ok {

			}
			sub, _err := entities.UnpackSubscription(result.Value)
			if _err != nil {
				// delete the subscription
				newChannelSubscriptionStore.Delete(ctx, db.Key(result.Key))
				logger.Infof("Invalid subscription %s", result.Value)
			}

			block.Size += 1
			if block.Size >= constants.MaxBlockSize {
				block.Closed = true
				block.NodeHeight = chain.MLChainApi.GetCurrentBlockNumber()
			}
			subCountKey := db.Key("/" + block.BlockId + "/" + sub.Signature)
			subCount := 0
			hasSub, err := blockStateTxn.Has(ctx, subCountKey)
			if err != nil {

			}

			if hasSub {
				subCountInBlock, err := blockStateTxn.Get(ctx, subCountKey)
				if err != nil {
					logger.Info("Could not get sub form block store %s/%s: %o", block.BlockId, sub.Signature, err)
				}
				subCount, convErr := strconv.Atoi(string(subCountInBlock))
				if convErr != nil {
					blockStateTxn.Delete(ctx, subCountKey)
					logger.Error("Invalid sub count for %s", sub.Signature)
				}
				blockStateTxn.Put(ctx, subCountKey, []byte(strconv.Itoa(subCount+1)))
			} else {
				blockStateTxn.Put(ctx, subCountKey, []byte("1"))
			}
			// tag the subscription with the blockid???
			subTotalCountData, err := channelSubscriptionCountStore.Get(ctx, db.Key("/"+sub.Signature))
			if err != nil {
				logger.Infof("Could not get sub total count sub store %s: %o", sub.Signature, err)
			}
			subTotalCount, convErr := strconv.Atoi(string(subTotalCountData))
			if convErr != nil {
				channelSubscriptionCountStore.Delete(ctx, db.Key("/"+sub.Signature))
				logger.Errorf("Invalid sub count for %s", sub.Signature)
			}
			channelSubCount := strconv.Itoa(subTotalCount + subCount)
			// save the proof and the batch
			block.Hash = hexutil.Encode(crypto.Keccak256Hash([]byte(sub.Signature + channelSubCount + block.Hash)))
			err = blockStateTxn.Put(ctx, db.Key(constants.CurrentDeliveryProofBlockStateKey), block.MsgPack())
			if err != nil {
				logger.Errorf("Unable to update State store errror %o", err)
				blockStateTxn.Discard(ctx)
				return
			}
			err = blockStateTxn.Commit(ctx)
			if err != nil {
				logger.Errorf("Unable to commit state update transaction errror %o", err)
				blockStateTxn.Discard(ctx)
				return
			}
			// dispatch the subscription and the block
			if block.Closed {
				// dispatch the block
				channelpool.OutgoingDeliveryProof_BlockC <- block
			}
			// constants.OutgoingDeliveryProofC <- proof
		}

	}

}

// func ProcessNewSubscriptionV1(ctx context.Context, sub *entities.Subscription, channelsubscribersRPC_D_countStore *db.Datastore, channelSubscriberStore *db.Datastore) {
// 	if sub.Broadcast {
// 		UnpackServerIdentity.SubscriptionD_P2P_C <- sub
// 	}
// 	trx, err := channelsubscribersRPC_D_countStore.NewTransaction(ctx, false)
// 	logger.Info("TRANSACTION INITIATED ******")
// 	if err != nil {
// 		logger.Infof("Transaction err::: %o", err)
// 	}
// 	cscstore, err := trx.Get(ctx, db.Key(sub.Key()))
// 	increment := -1
// 	if sub.Action == utils.Join {
// 		increment = 1
// 		channelSubscriberStore.Set(ctx, db.Key(sub.Key()), sub.MsgPack(), false)

// 	} else {
// 		channelSubscriberStore.Delete(ctx, db.Key(sub.Key()))
// 	}
// 	if len(cscstore) == 0 {
// 		cscstore = []byte("0")
// 	}
// 	cscstoreint, err := strconv.Atoi(string(cscstore))
// 	cscstoreint += increment
// 	channelsubscribersRPC_D_countStore.Set(ctx, db.Key(sub.Channel), []byte(strconv.Itoa(cscstoreint)), true)
// 	logger.Info("TRANSACTION ENDED ******")
// 	trx.Commit(ctx)
// }
