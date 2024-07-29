package node

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/dgraph-io/badger"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/query"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/chain"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto/schnorr"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/db"
	p2p "github.com/mlayerprotocol/go-mlayer/pkg/core/p2p"
)

// Keep a record of all messages sent within a cycle per subnet
func TrackReward(ctx *context.Context) {
	cfg, _ := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)
	// validator := (*cfg).PublicKey  
	claimedRewardStore, ok := (*ctx).Value(constants.ClaimedRewardStore).(*db.Datastore)
	if !ok {
		panic("Unable to load claimedRewardStore")
	}
	eventCounterStore, ok := (*ctx).Value(constants.EventCountStore).(*db.Datastore)
	if !ok {
		panic("Unable to load eventCounterStore")
	}
	currentCycle := chain.API.GetCurrentCycle()
	if !ok {
		panic("Unable to access reward store")
	}
	lastCycleClaimedKey :=  datastore.NewKey("claimed")
	lastClaimedCycle := uint64(0)
	//batch, err :=	claimedRewardStore.Batch(*ctx)
	// if err != nil {
	// 	panic(err)
	// }
	lastClaimed, err := claimedRewardStore.Get(*ctx, lastCycleClaimedKey)
		logger.Infof("CURRENTCYCLECLAIM %d", lastClaimed)
		if err != nil  {
			if err != badger.ErrKeyNotFound {
				logger.Error(err)
				return;
			}
		} else {
			lastClaimedCycle = encoder.NumberFromByte(lastClaimed)
		}
	if lastClaimedCycle >= currentCycle {
		return
	}
	for i := lastClaimedCycle+1; i < currentCycle; i++ {
		// get message count per subnet in this cycle in groups of 100 and create a proof batch
		cycleKey :=  fmt.Sprintf("%s/%d", cfg.PublicKey, i)
	
		subnetList, err := eventCounterStore.Query(*ctx, query.Query{
			Prefix: cycleKey,
		})

		
		if err != nil {
			continue
		}
		defer subnetList.Close()
		cost, err := p2p.GetCycleMessageCost(ctx, i)
		if err != nil {
			logger.Errorf("GetCycleMessageCost: %v", err)
			break;
		}
		index := uint64(0)
		rewardBatch := entities.NewRewardBatch(cfg, i, index, cost)
		for  rsl := range subnetList.Next() {
				rewardBatch.Append(entities.SubnetCount{
					Subnet: rsl.Key,
				EventCount: encoder.NumberFromByte(rsl.Value),
				})
				if rewardBatch.Closed {
					// save batch
					unclaimedBatchKey :=  datastore.NewKey(fmt.Sprintf("unclaimed/%s", rewardBatch.Id))
					 err :=	claimedRewardStore.Put(*ctx, unclaimedBatchKey, rewardBatch.MsgPack())
					if err != nil {
						panic(err)
					}
					go processSentryRewardBatch(ctx, cfg, *rewardBatch)
					index++
					rewardBatch = entities.NewRewardBatch(cfg, i, index, cost)
					
					// set time stamp and sign batch
					// rewardBatch.Sign(cfg.PrivateKey)
		
					// // identify nodes to provide proof
					// totalLicenses  := big.NewInt(int64(100))
					// hashNumber := big.NewInt(0).SetBytes(rewardBatch.DataHash)
					// salt := big.NewInt(0).Mod(hashNumber, totalLicenses)
					// startLicence := big.NewInt(0).Mod(big.NewInt(0).Div(hashNumber, salt), totalLicenses)
					// var nodeIds = []int{}(15)
					// for i = 0; i < len(nodeIds); i++ {
					// 	nodeIds[i] =  big.NewInt(0).Mod( big.NewInt(0).SetBytes(utils.lcg(startLicence.Uint64() + i * salt.Uint64())), totalLicenses)
					// }


					// get nodes associated with licences
					// request proof from nodes

					// send proof to blokchain for redemption
				}
			}
		}
			
	
	// key := fmt.Sprintf("%s", validator)
	// for {
	// 	result, err := claimedRewardStore.Query(*ctx, query.Query{
	// 		Prefix: key,
	// 	})
	// 	if err != nil {
	// 		continue
	// 	}
	// 	for {
	// 		select {
	// 		case rsl := <-result.Next():
	// 			rsl.Value

	// 		}
	// 	}
		
	// }
}


func processSentryRewardBatch(ctx *context.Context, cfg *configs.MainConfiguration, batch entities.RewardBatch) {
	hashNumber :=  new(big.Int).SetBytes(batch.DataHash)
		totalLicenses  := chain.API.GetLicenseCount(batch.Cycle)
		salt := new(big.Int).Mod(hashNumber, big.NewInt(1000)).Add(big.NewInt(1), big.NewInt(0))
		startLicence := new(big.Int).Mod(new(big.Int).Div(hashNumber, salt), totalLicenses).Add(big.NewInt(1000), big.NewInt(0))
		// var licenses = [15]*big.Int{}
		//1. Identify nodes licences associated with the batch hash
		//2. loop through licence and request commitment from nodes by sending them the batch data. The nodes will call "schnorr.ComputeNonce(pk, msg)"
				// and return the nonce publicKey to you
		noncePubKeys := []*btcec.PublicKey{}
		sentryPubKeys := []*btcec.PublicKey{}
		signatures := [][]byte{}
		operatorPubKeys := []string{}
		for i := 0; i<30; i++ {
			decodedLicence :=  new(big.Int).Mod(utils.Lcg(startLicence.Uint64() + (uint64(i) * salt.Uint64())), totalLicenses)
			operator, err := chain.API.LicenceOperator(decodedLicence)
			if err != nil {
				continue
			} else {
				if len(operator) > 0 {
					// get the operators address from the dht
				validator, err := p2p.GetDhtValue("/ml/val/" + hex.EncodeToString(operator))
				if err != nil {
					continue
				}
				
				// clear out the batch data. Its unncessary to send as the receiving node wont use it
				tmpBatch := batch
				tmpBatch.Data = []entities.SubnetCount{}
				// request commitment from validator
				payload := p2p.NewP2pPayload(cfg, p2p.P2pActionGetCommitment, tmpBatch.MsgPack())
				if err != nil {
					continue
				}
				response, err := payload.SendRequest( cfg.PrivateKeyBytes, hex.EncodeToString(validator)) // nonce public key
				if err != nil {
					continue
				}
				if response.IsValid(cfg.ChainId) {
					noncePubKey, err := btcec.ParsePubKey(response.Data)
					if err != nil {
						continue
					}
					pubKey, err := btcec.ParsePubKey(operator)
					if err != nil {
						continue
					}
					
					noncePubKeys = append(noncePubKeys, noncePubKey)
					sentryPubKeys = append(sentryPubKeys, pubKey)
					operatorPubKeys = append(operatorPubKeys, hex.EncodeToString(operator))
					
				}
				if len(noncePubKeys) == 20 {
					hash, err := batch.GetHash(cfg.ChainId)
					if err != nil {

					}
					
					aggPubKey, challenge, commitment := schnorr.ComputeSigningParams(sentryPubKeys, noncePubKeys, hash )
					signReq := entities.SignatureRequestData{
						AggPubKey: aggPubKey.SerializeCompressed(),
						Challenge: challenge,
						Commitment: commitment,
					}
					
					signingPayload := p2p.NewP2pPayload(cfg, p2p.P2pActionGetSentryProof, signReq.MsgPack())
					if err != nil {
						continue
					}
					for _, validator := range operatorPubKeys {
						signResp, err := signingPayload.SendRequest(cfg.PrivateKeyBytes, validator)
						if err != nil {
							continue
						}
						if signResp.IsValid(cfg.ChainId) {
							// we have received a valid signature from this node
							signatures = append(signatures, signResp.Data)
						}
					}
					if len(signatures) == 20 { // all 20 have signed
						// aggregate signature
						aggSig := schnorr.AggregateSignatures(signatures)
						_, err := chain.API.ClaimReward(chain.ClaimData{
							SubnetRewardCount: batch.Data,
							Signature: [32]byte(aggSig),
							Commitment: commitment,
							PubKeys: sentryPubKeys,
							Cycle: batch.Cycle,
						})
						if err != nil {
							// retry again
							logger.Errorf("processSentryRewardBatch: %v", err)
						} else {
							// mark it as finalized
							logger.Infof("processSentryRewardBatch: Successful..... %d/%s", batch.Cycle, batch.Id)
						}

					}
					
					break
				}
			}
		}
		// get the node that owns the license
	}
					
	//3. make a list of the public keys 20 nodes that respond. then get the signing parameters with schnoor.ComputeSigningParams(pks,noncePks,message)
	//4. send the signing params to each node on the list. Each node will call ComputeSignature the return the signature
	//5. aggregate the signatures with schnorr.AggregateSignatures()

}

// get the cycle from the store else from the network
