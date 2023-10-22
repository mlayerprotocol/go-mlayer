package processor

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ipfs/go-datastore/query"
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/db"
	"github.com/mlayerprotocol/go-mlayer/utils"
)

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

		var block *utils.Block
		blockStateTxn, err := subscriptionBlockStateStore.NewTransaction(ctx, false)
		if err != nil {
			utils.Logger.Errorf("Subscription Block state store error %w", err)
			// invalid proof or proof has been tampered with
			panic(err)
		}

		blockData, err := blockStateTxn.Get(ctx, db.Key(utils.CurrentSubscriptionBlockStateKey))
		if err != nil {
			logger.Infof("Block state store error - key: %s, %v", utils.CurrentSubscriptionBlockStateKey, err)
			// invalid proof or proof has been tampered with

			if string(err.Error()) == "datastore: key not found" {
				block = utils.NewBlock()
				logger.Infof("Erro: %s", block.ToJSON())
				blockStateTxn.Put(ctx, db.Key(utils.CurrentDeliveryProofBlockStateKey), block.ToJSON())
				err = blockStateTxn.Commit(ctx)
				if err != nil {
					panic(err)
				}
			}
			continue
		}
		if len(blockData) > 0 {
			block, err = utils.BlockFromBytes(blockData)
			if err != nil {
				logger.Errorf("Invalid block data %w", err)
				// invalid proof or proof has been tampered with
				blockStateTxn.Discard(ctx)
				continue
			}
		}

		results, err := newChannelSubscriptionStore.Query(ctx, query.Query{
			Prefix: "/",
		})
		if err != nil {
			utils.Logger.Errorf("New Channel Subscriber Store Query Error %w", err)
			return
		}
		for {
			result, ok := <-results.Next()
			if !ok {

			}
			sub, _err := utils.SubscriptionFromBytes(result.Value)
			if _err != nil {
				// delete the subscription
				newChannelSubscriptionStore.Delete(ctx, db.Key(result.Key))
				logger.Info("Invalid subscription %s", result.Value)
			}

			block.Size += 1
			if block.Size >= utils.MaxBlockSize {
				block.Closed = true
				block.NodeHeight = utils.GetNodeHeight()
			}
			subCountKey := db.Key("/" + block.BlockId + "/" + sub.Signature)
			subCount := 0
			hasSub, err := blockStateTxn.Has(ctx, subCountKey)
			if err != nil {

			}

			if hasSub {
				subCountInBlock, err := blockStateTxn.Get(ctx, subCountKey)
				if err != nil {
					logger.Info("Could not get sub form block store %s/%s: %w", block.BlockId, sub.Signature, err)
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
				logger.Info("Could not get sub total count sub store %s: %w", sub.Signature, err)
			}
			subTotalCount, convErr := strconv.Atoi(string(subTotalCountData))
			if convErr != nil {
				channelSubscriptionCountStore.Delete(ctx, db.Key("/"+sub.Signature))
				logger.Error("Invalid sub count for %s", sub.Signature)
			}
			channelSubCount := strconv.Itoa(subTotalCount + subCount)
			// save the proof and the batch
			block.Hash = hexutil.Encode(utils.Hash(sub.Signature + channelSubCount + block.Hash))
			err = blockStateTxn.Put(ctx, db.Key(utils.CurrentDeliveryProofBlockStateKey), block.ToJSON())
			if err != nil {
				logger.Errorf("Unable to update State store errror %w", err)
				blockStateTxn.Discard(ctx)
				return
			}
			err = blockStateTxn.Commit(ctx)
			if err != nil {
				logger.Errorf("Unable to commit state update transaction errror %w", err)
				blockStateTxn.Discard(ctx)
				return
			}
			// dispatch the subscription and the block
			if block.Closed {
				// dispatch the block
				utils.OutgoingDeliveryProof_BlockC <- block
			}
			// utils.OutgoingDeliveryProofC <- proof
		}

	}

}

// func ProcessNewSubscriptionV1(ctx context.Context, sub *utils.Subscription, channelsubscribersRPC_D_countStore *db.Datastore, channelSubscriberStore *db.Datastore) {
// 	if sub.Broadcast {
// 		utils.SubscriptionD_P2P_C <- sub
// 	}
// 	trx, err := channelsubscribersRPC_D_countStore.NewTransaction(ctx, false)
// 	utils.Logger.Info("TRANSACTION INITIATED ******")
// 	if err != nil {
// 		utils.Logger.Infof("Transaction err::: %w", err)
// 	}
// 	cscstore, err := trx.Get(ctx, db.Key(sub.Key()))
// 	increment := -1
// 	if sub.Action == utils.Join {
// 		increment = 1
// 		channelSubscriberStore.Set(ctx, db.Key(sub.Key()), sub.ToJSON(), false)

// 	} else {
// 		channelSubscriberStore.Delete(ctx, db.Key(sub.Key()))
// 	}
// 	if len(cscstore) == 0 {
// 		cscstore = []byte("0")
// 	}
// 	cscstoreint, err := strconv.Atoi(string(cscstore))
// 	cscstoreint += increment
// 	channelsubscribersRPC_D_countStore.Set(ctx, db.Key(sub.Channel), []byte(strconv.Itoa(cscstoreint)), true)
// 	utils.Logger.Info("TRANSACTION ENDED ******")
// 	trx.Commit(ctx)
// }
