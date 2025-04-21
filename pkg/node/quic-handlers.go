package node

import (
	"bytes"
	"context"
	"encoding/hex"
	"sort"

	"fmt"
	"io"
	"log"

	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto/bls"
	dsquery "github.com/mlayerprotocol/go-mlayer/internal/ds/query"
	"github.com/mlayerprotocol/go-mlayer/internal/service"
	"github.com/mlayerprotocol/go-mlayer/internal/system"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/p2p"
	"github.com/quic-go/quic-go"
)

func HandleQuicConnection(ctx *context.Context, cfg *configs.MainConfiguration, connection quic.Connection) {
	for {
		logger.Infof("NewRemoteStreamStarted: %s", connection.RemoteAddr().String())
		newStream, err := connection.AcceptStream(*ctx)
		if err != nil {
			if _, ok := err.(*quic.StreamError); ok {
				continue
			} else {
				return
			}
		}
		go func(stream quic.Stream) {
			if err != nil {
				logger.Fatal(err)
			}
			defer stream.Close()

			// Read the client's request (the filename)
			buf := make([]byte, 1024)
			data := bytes.Buffer{}
			for {

				n, err := stream.Read(buf) // Read into the buffer
				data.Write(buf[:n])
				if n < len(buf) || n == 0 || err == io.EOF {
					break // End of file, stop reading
				}
				if err != nil {
					logger.Error("HandleQuicConnection/writedata: ", err) // Handle error
					return
				}
			}
			if len(data.Bytes()) == 0 {
				return
			}
			payload, err := p2p.UnpackP2pPayload(data.Bytes())

			if err != nil {

				logger.Errorf("HandleQuicConnection/UnpackP2pPayload: %v, %d", err, data.Len())
				return
			}
			// if !payload.IsValid(cfg.ChainId) {
			// 	logger.Error(fmt.Errorf("HandleQuicConnection: invalid payload signature for action %d", payload.Action))
			// 	return
			// }
			p2p.NodeQuicPool.AddConnection(*ctx, connection.RemoteAddr().String(), hex.EncodeToString(payload.Signer), connection)
			response := p2p.NewP2pPayload(cfg, p2p.P2pActionResponse, []byte{})
			response.Id = payload.Id
			switch payload.Action {
			case p2p.P2pActionPostEvent:
				eventPayload := entities.EventPayload{}
				err := encoder.MsgPackUnpackStruct(payload.Data, &eventPayload)
				modelType := entities.GetModelTypeFromEventType(eventPayload.EventType)
				if err == nil {
					event, err := entities.UnpackEvent(eventPayload.Event, modelType)
					if err == nil {
						eventResp, err := service.HandleNewPubSubEvent(*event, cfg.Context)
						logger.Infof("EVENTRESP: %+v", eventResp)
						resp := entities.EventDelta{}
						if err == nil {
							for _, stateData := range eventResp.States {
								delta, errD := service.GetStateDelta(stateData, *event)
								if errD != nil {
									err = errD
									break 
								}
								resp.Deltas = append(resp.Deltas, *delta)
								// stateBytes, err := encoder.MsgPackStruct(stateData.StateData)
								// // stateBytes, err := dsquery.GetStateBytesFromEventPath(event.GetPath())
								// if err == nil {
								// 	delta := make(map[string]interface{})
								// 	state := make(map[string]interface{})
								// 	prevState := map[string]interface{}{}
								// 	encoder.MsgPackUnpackStruct(stateBytes, &state)
								// 	var previousStateBytes []byte
								
								// 	if event.PreviousEvent.ID != "" && event.PreviousEvent.EntityPath.Model == event.GetPath().Model {
								// 		previousStateBytes, err = dsquery.GetStateBytesFromEventPath(&event.PreviousEvent)
								// 		logger.Infof("STATEMAP %v", previousStateBytes)
								// 		if err == nil {
								// 			encoder.MsgPackUnpackStruct(previousStateBytes, &prevState)
								// 			delta = utils.GetDifference(prevState, state)
								// 		}
								// 	} else {
								// 		delta = state
								// 	}
								
								// 	if _, ok := delta["id"]; !ok {
								// 		delta["id"] = prevState["id"]
								// 	}
								// 	b, err := encoder.MsgPackStruct(delta)
								// 	if err == nil {
								// 		// logger.Infof("STATEDELTA %v", delta)
								// 		// pack, err := encoder.MsgPackStruct(delta)
								// 		//if err == nil {
								// 		var previousHash []byte
								// 		if len(fmt.Sprint(prevState["h"])) > 0 {
								// 			previousHash, _ = hex.DecodeString(fmt.Sprint(prevState["h"]))
								// 		}
								// 		resp.Deltas = append(resp.Deltas, entities.StateDelta{Type: stateData.Type, Delta: b, StateID: stateData.StateID, PreviousHash: previousHash})
								// 	}

								// 	// sig, _ := crypto.SignSECP(resp.Hash, cfg.PrivateKeySECP)
								// 	// resp.Signatures = []entities.SignatureData{
								// 	// 	{Signature: entities.HexString(hex.EncodeToString(sig)), PublicKey: entities.PublicKeyString(cfg.PublicKeySECPHex)},
								// 	// }
								// 	// resp.Signature, _ = crypto.SignSECP(resp.Hash, cfg.PrivateKeySECP)
								// 	// resp.Validator = cfg.PublicKeySECP

								// 	//}

								// }
							}
							resp.Hash, _ = resp.GetHash()
							resp.Event = event.ID
							sig, _ := bls.BlsProofGenerator.Sign(cfg.PrivateKeyBLS, resp.Hash)
							resp.SignatureData = entities.BlsSignatureData{
								Signature: entities.HexString(hex.EncodeToString(sig)), PublicKeys: []entities.PublicKeyString{entities.PublicKeyString(cfg.PublicKeyBLSPHex)},
							}
							// resp.Validator = cfg.PublicKeyBLSPHex
							response.Data, _ = encoder.MsgPackStruct(resp)

						}
					}

				}
				if err != nil {
					logger.Errorf("ErrorProccessingEvent: %v", err)
					response.Error = err.Error()
					response.ResponseCode = 500
				}
			case p2p.P2pActionNotifyValidEvent:

				eventPath := entities.EventPath{}
				err := encoder.MsgPackUnpackStruct(payload.Data, &eventPath)

				if err == nil {
					if v, err := system.Mempool.GetData(eventPath.ID); err == nil {

						dstate := dsquery.DataStates{}
						err = encoder.MsgPackUnpackStruct(*v, &dstate)
						if err == nil {
							logger.Infof("P2pActionNotifyValidEvent %+v", dstate)
							err = dstate.Commit(nil, nil, nil, eventPath.ID, err)
							if err == nil {
								response.Data = eventPath.MsgPack()
							}
						} else {
							logger.Errorf("P2pActionNotifyValidEvent %v", err)
						}
					}
				} else {
					logger.Errorf("P2pActionNotifyValidEventDsError %v", err)
				}
				if err != nil {
					logger.Errorf("P2pActionNotifyValidEvent %v", err)
					response.Error = err.Error()
					response.ResponseCode = 500
				}
			case p2p.P2pActionSyncState:
				logger.Infof("SyncingValidState")
				stateDelta := entities.EventDelta{}
				err := encoder.MsgPackUnpackStruct(payload.Data, &stateDelta)
				if err == nil {
					valid := true
					hash, _ := stateDelta.GetHash()
					pubKeys := stateDelta.SignatureData.PublicKeys
					sort.Slice(&pubKeys, func(i, j int) bool {
						return string(pubKeys[i]) < string(pubKeys[j])
					})
					// for _, sig := range stateDelta.SignatureData.PublicKeys {
					// 	if v, _ := crypto.VerifySignatureSECP(sig.PublicKey.GetBytes(), hash, sig.Signature.GetBytes()); !v {
					// 		valid = false
					// 	}
					// }
					pubKeysBytes := [][]byte{}
					for _, pubk := range pubKeys {
						b, _ := hex.DecodeString(string(pubk))
						pubKeysBytes = append(pubKeysBytes, b)
					}
					valid, err := bls.BlsProofGenerator.VerifyAggregateSignature(pubKeysBytes, hash, stateDelta.SignatureData.Signature.GetBytes())

					if valid {
						dataState := dsquery.NewDataStates(stateDelta.Event, cfg)
						for _, deltaData := range stateDelta.Deltas {
							delta := map[string]interface{}{}
							encoder.MsgPackUnpackStruct([]byte(deltaData.Delta), delta)
							id := fmt.Sprint(delta["id"])
							model := deltaData.Type
							var newState = entities.GetStateModelFromEntityType(model)
							//var prevEventPath *entities.EventPath
							if deltaData.PreviousHash != nil {
								// prevEvent := &entities.Event{}
								newState, _, err = service.SyncStateFromPeer(fmt.Sprint(delta["id"]), model, cfg, string(stateDelta.SignatureData.PublicKeys[0]))
								if err != nil {
									logger.Errorf("ErrorSyncingOldState")

								}
								// if prevEvent != nil {
								// 	prevEventPath = prevEvent.GetPath()
								// }
							}
							if err == nil {
								utils.MapToStruct(delta, &newState)
								dataState.AddCurrentState(model, id, newState)
							} else {
								logger.Errorf("state sync error: %v", err)
							}

						}
						if err == nil {
							// dataState.AddEvent(stateDelta.Event)
							err = dataState.Commit(nil, nil, nil, stateDelta.Event, nil)
						} else {
							logger.Errorf("state sync error: %v", err)
						}
					}
				}
				if err != nil {
					response.Error = err.Error()
					response.ResponseCode = 500
				}

			default:
				response, err = p2p.ProcessP2pPayload(cfg, payload, false)

			}

			if err != nil {
				logger.Error("HandleQuicConnection/processP2pPayload: ", err)
				return
			}
			_, err = stream.Write(response.MsgPack())
			if err != nil {
				log.Fatalf("Failed to send file: %v", err)
			}
		}(newStream)
	}

}
