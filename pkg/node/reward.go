package node

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/dgraph-io/badger"
	"github.com/ipfs/go-datastore"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"

	"github.com/mlayerprotocol/go-mlayer/internal/chain"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto/schnorr"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/query"
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
	// eventCounterStore, ok := (*ctx).Value(constants.EventCountStore).(*db.Datastore)
	// if !ok {
	// 	panic("Unable to load eventCounterStore")
	// }
	currentCycle, err := chain.DefaultProvider(cfg).GetCurrentCycle()
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
	if lastClaimedCycle >= currentCycle.Uint64() {
		return
	}
	for i := lastClaimedCycle+1; i < currentCycle.Uint64(); i++ {
		// TODO loop through index till no data
		rewardBatch, err := generateBatch(i, 1, ctx, cfg)
		if err != nil {
			break
		}
		if rewardBatch.Closed {
			// save batch
			unclaimedBatchKey :=  datastore.NewKey(fmt.Sprintf("unclaimed/%s", rewardBatch.Id))
			err :=	claimedRewardStore.Put(*ctx, unclaimedBatchKey, rewardBatch.MsgPack())
			if err != nil {
				panic(err)
			}
			go processSentryRewardBatch(ctx, cfg, rewardBatch)
			// index++
			// rewardBatch = entities.NewRewardBatch(cfg, i, index, cost)
			
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

func generateBatch(cycle uint64, index int, ctx *context.Context, cfg *configs.MainConfiguration) (*entities.RewardBatch, error) {
	
	subnetList := []models.EventCounter{}
		claimed := false
		err := query.GetManyWithLimit(models.EventCounter{Cycle: cycle, Validator: entities.PublicKeyString(cfg.PublicKey), Claimed: &claimed }, &subnetList, &map[string]query.Order{"count": query.OrderDec}, entities.MaxBatchSize, index*entities.MaxBatchSize)
		if err != nil {
			return nil, err
		}
		// defer subnetList.Close()
		logger.Infof("ListLen: %d", len(subnetList))
		if len(subnetList) == 0 {
			return nil, fmt.Errorf("list empty")
		}
		cost, err := p2p.GetCycleMessageCost(ctx, cycle)
		if err != nil {
			logger.Errorf("GetCycleMessageCost: %v", err)
			return nil, err
		}
		
		rewardBatch := entities.NewRewardBatch(cfg, cycle, index, cost, len(subnetList), cfg.PublicKeySECP)
		for  _, rsl := range subnetList {
				rewardBatch.Append(entities.SubnetCount{
					Subnet: rsl.Subnet,
				EventCount: rsl.Count,
				})
				if rewardBatch.Closed {
					break
				}
		}
		return rewardBatch, nil
}

func processSentryRewardBatch(ctx *context.Context, cfg *configs.MainConfiguration, batch *entities.RewardBatch) {
	logger.Infof("Processing Batch....: %v", batch.Id)
	claimedRewardStore, ok := (*ctx).Value(constants.ClaimedRewardStore).(*db.Datastore)
	if !ok {
		panic("Unable to load claimedRewardStore") 
	}
	hashNumber :=  new(big.Int).SetBytes(batch.DataHash)
	logger.Infof("Hash: %s", hashNumber)
		totalLicenses, err := chain.DefaultProvider(cfg).GetTotalSentryLicenseCount(big.NewInt(int64(batch.Cycle)))
		logger.Infof("License count: %s", totalLicenses)
		if err != nil {
			panic(err)
		}
		salt := new(big.Int).Mod(hashNumber, big.NewInt(1000))
		salt = new(big.Int).Add(salt, big.NewInt(1))
		logger.Infof("Salt: %s", new(big.Int).Mod(hashNumber, big.NewInt(1000)))
		startLicence := new(big.Int).Mod(new(big.Int).Div(hashNumber, salt), totalLicenses).Add(big.NewInt(1000), big.NewInt(0))
		
		// var licenses = [15]*big.Int{}
		//1. Identify nodes licences associated with the batch hash
		//2. loop through licence and request commitment from nodes by sending them the batch data. The nodes will call "schnorr.ComputeNonce(pk, msg)"
				// and return the nonce publicKey to you
		noncePubKeys := []*btcec.PublicKey{}
		sentryPubKeys := []*btcec.PublicKey{}
		signatures := [][]byte{}
		operatorPubKeys := []string{}
		max := 30
		if totalLicenses.Uint64() < 120 {
			max = int(totalLicenses.Uint64() / 3)
		}
		if totalLicenses.Uint64() < 3 {
			max = int(totalLicenses.Uint64()) - 1
		}
		logger.Infof("Max Proofs Needed: %d", max)
		for i := 0; i<max*2; i++ {
			decodedLicence :=  new(big.Int).Mod(utils.Lcg(startLicence.Uint64() + (uint64(i) * salt.Uint64())), totalLicenses)
			decodedLicence = new(big.Int).Add(decodedLicence, big.NewInt(1000))
			logger.Infof("Start Licence: %s", decodedLicence)
			operatorBytes, err := chain.DefaultProvider(cfg).GetSentryLicenceOperator(decodedLicence)
			if err != nil {
				logger.Info(err)
				continue
			} else {
				if len(operatorBytes) > 0 {
					operator := hex.EncodeToString(operatorBytes)
					if operator == hex.EncodeToString(cfg.PublicKeySECP) {
						continue
					}
					// get the operators address from the dht
					
				// validator, err := p2p.GetOperatorMultiAddress(hex.EncodeToString(operator), cfg.ChainId)
				// if err != nil {
				// 	continue
				// }

				
				// clear out the batch data. Its unncessary to send as the receiving node wont use it
				batchCopy := *batch
				tmpBatch := &batchCopy
				tmpBatch.Data = []entities.SubnetCount{}
				
				logger.Infof("BATCHDATA: %v", batch)
				// request commitment from validator
				payload := p2p.NewP2pPayload(cfg, p2p.P2pActionGetCommitment, tmpBatch.MsgPack())
				
				response, err := payload.SendRequest( cfg.PrivateKeyBytes, operator) // nonce public key
				if err != nil {
					logger.Error(err)
					continue
				}
				
				logger.Infof("IsValid Response: %s, %v", response.Id, response.IsValid(cfg.ChainId))
				
				if response.IsValid(cfg.ChainId) {
					
					noncePubKey, err := btcec.ParsePubKey(response.Data)
					
					if err != nil {
						logger.Errorf("Response Nonce: %v, %v", err, response.Data)
				
						continue
					}
					
					pubKey, err := btcec.ParsePubKey(operatorBytes)
					if err != nil {
						logger.Errorf("Response Nonce: %v, %v", err, response.Data)
				
						continue
					}
					
					noncePubKeys = append(noncePubKeys, noncePubKey)
					sentryPubKeys = append(sentryPubKeys, pubKey)
					operatorPubKeys = append(operatorPubKeys, operator)
					
				}
				
				
				if len(noncePubKeys) == max {
					
					// hash, _ := batch.GetHash(cfg.ChainId)
					// if err != nil {

					// }
					proofData := batch.GetProofData(cfg.ChainId)
					hash, _ := proofData.GetHash()
					_, challenge, commitment := schnorr.ComputeSigningParams(sentryPubKeys, noncePubKeys, hash )
					signReq := entities.SignatureRequestData{
						ProofHash: hash[:],
						// AggPubKey: aggPubKey.SerializeCompressed(),
						Challenge: challenge,
						// Commitment: commitment,
					}
					logger.Infof("ReceivedAllNonces: challange:%s, commit:%s", hex.EncodeToString(challenge), commitment)
					signingPayload := p2p.NewP2pPayload(cfg, p2p.P2pActionGetSentryProof, signReq.MsgPack())
					if err != nil {
						continue
					}
					for _, validator := range operatorPubKeys {
						signResp, err := signingPayload.SendRequest(cfg.PrivateKeyBytes, validator)
						if err != nil {
							continue
						}
						if signResp.Error != "" {
							logger.Errorf("GetProofRequestError %s", signResp.Error)
							continue
						}
						logger.Infof("IsValid Response For Signature: %s, %v", response.Id, response.IsValid(cfg.ChainId))
						if signResp.IsValid(cfg.ChainId) {
							// we have received a valid signature from this node
							signatures = append(signatures, signResp.Data)
						}
					}
					if len(signatures) == max { // all 20 have signed
						// aggregate signature
						logger.Infof("Signatures: %v", signatures)
						aggSig := schnorr.AggregateSignatures(signatures)
						_, err := chain.DefaultProvider(cfg).ClaimReward(entities.ClaimData{
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
							// store it as a pending claim
							signers := [][]byte{}
							for _, k := range sentryPubKeys {
								signers = append(signers, k.SerializeCompressed())
							}
							proofData.Signature  = aggSig	
							proofData.Signers = signers
							proofData.Commitment = []byte(commitment)
							pendingClaimsKey :=  datastore.NewKey(fmt.Sprintf("validClaim/%s",  hex.EncodeToString(hash[:])))
							err = claimedRewardStore.Put(*ctx,pendingClaimsKey, proofData.MsgPack() )
							if err != nil {
								logger.Error(err)
							} else {
								for _, d := range batch.Data {
									logger.Infof("[")
									logger.Infof("{'subnetId':'0x%s', 'amount':'%s'},", strings.ReplaceAll(d.Subnet, "-", ""), new(big.Int).SetBytes(d.Cost) )
									logger.Infof("]")
								}
								for _, k := range sentryPubKeys {
									logger.Infof("[")
									logger.Infof("{'x':'%s','y':'%s'},", k.X(), k.Y() )
									logger.Infof("]")
								}
								logger.Info("Cycle", batch.Cycle)
								logger.Info("Index", batch.Index)
								logger.Info("Cost", batch.TotalValue)
								logger.Info("Validator", hex.EncodeToString(batch.Validator))
								logger.Info("DataHash", hex.EncodeToString(proofData.DataHash))
								logger.Info("ProofHash: ", hex.EncodeToString(hash[:]))
								logger.Info("commitment: ", hex.EncodeToString(proofData.Commitment))
								logger.Info("signature: ", hex.EncodeToString(proofData.Signature))
								logger.Infof("processSentryRewardBatch: Successful..... %d/%s, %v, %v", batch.Cycle, batch.Id, sentryPubKeys[0], commitment)
							}
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
