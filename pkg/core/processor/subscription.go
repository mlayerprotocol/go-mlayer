package processor

import (
	"context"
	"strconv"
	"sync"
	"time"

	db "github.com/ByteGum/go-icms/pkg/core/db"
	"github.com/ByteGum/go-icms/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ipfs/go-datastore/query"
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

		var block utils.Block
		txn, err := subscriptionBlockStateStore.NewTransaction(ctx, false)
		if err != nil {
			utils.Logger.Errorf("Subscription Block state store error %w", err)
			// invalid proof or proof has been tampered with
			panic(err)
		}
		blockData, err := txn.Get(ctx, db.Key(utils.CurrentSubscriptionBlockStateKey))
		if err != nil {
			logger.Errorf("State store error %w", err)
			// invalid proof or proof has been tampered with
			txn.Discard(ctx)
			return
		}
		if len(blockData) > 0 {
			block, err = utils.BlockFromBytes(blockData)
			if err != nil {
				logger.Errorf("Invalid block data %w", err)
				// invalid proof or proof has been tampered with
				txn.Discard(ctx)
				continue
			}
		} else {
			// generate a new batch
			block = utils.NewBlock()

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
			subKey := db.Key("/" + block.BlockId + "/" + sub.Signature)
			hasSub, err := txn.Has(ctx, subKey)
			if err != nil {

			}
			if hasSub {
				subCountInBlock, err := txn.Get(ctx, subKey)
				if err != nil {
					logger.Info("Could not get sub form block store %s/%s: %w", block.BlockId, sub.Signature, err)
				}
				subCount, convErr := strconv.Atoi(string(subCountInBlock))
				if convErr != nil {
					txn.Delete(ctx, subKey)
					logger.Error("Invalid sub count for %s", sub.Signature)
				}
				txn.Put(ctx, subKey, []byte(strconv.Itoa(subCount+1)))
			} else {
				txn.Put(ctx, subKey, []byte("1"))
			}
			// tag the subscription with the blockid???
			
			// save the proof and the batch
			block.Hash = hexutil.Encode(utils.Hash(sub.Signature + totalChannelSubscriptionCount + block.Hash))
			err = txn.Put(ctx, db.Key(utils.CurrentDeliveryProofBlockStateKey), block.ToJSON())
			if err != nil {
				logger.Errorf("Unable to update State store errror %w", err)
				txn.Discard(ctx)
				return
			}
			proof.Block = block.BlockId
			proof.Index = block.Size
			err = deliveryProofStore.Put(ctx, db.Key(proof.Key()), proof.ToJSON())
			if err != nil {
				txn.Discard(ctx)
				logger.Errorf("Unable to save proof to store error %w", err)
				return
			}
			err = localBlockStore.Put(ctx, db.Key(utils.CurrentDeliveryProofBlockState), block.ToJSON())
			if err != nil {
				logger.Errorf("Unable to save batch error %w", err)
				txn.Discard(ctx)
				return
			}
			err = txn.Commit(ctx)
			if err != nil {
				logger.Errorf("Unable to commit state update transaction errror %w", err)
				txn.Discard(ctx)
				return
			}
			// dispatch the proof and the batch
			if block.Closed {
				utils.OutgoingDeliveryProof_BlockC <- &block
			}
			utils.OutgoingDeliveryProofC <- proof
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
