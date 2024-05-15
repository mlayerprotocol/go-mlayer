package service

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"github.com/sirupsen/logrus"
)

/*
Validate an agent authorization
*/
func ValidateSubnetData(subnet *entities.Subnet, addressPrefix string) (currentSubnetState *models.SubnetState, err error) {
	// check fields of Subnet
	

	if len(subnet.Ref) > 60 {
		return nil, apperror.BadRequest("Subnet ref cannont be more than 40 characters")
	}
	if len(subnet.Ref) > 0 && !utils.IsAlphaNumericDot(subnet.Ref) {
		return nil, apperror.BadRequest("Ref must be alpha-numeric, and .")
	}
	var valid bool
	b, _ := subnet.EncodeBytes()
	switch subnet.SignatureData.Type {
	case entities.EthereumPubKey:
		valid = crypto.VerifySignatureECC(entities.AddressFromString(string(subnet.Account)).Addr, &b, subnet.SignatureData.Signature)

	case entities.TendermintsSecp256k1PubKey:

		decodedSig, err := base64.StdEncoding.DecodeString(subnet.SignatureData.Signature)
		if err != nil {
			return nil, err
		}

		msg, err := subnet.GetHash()

		

		if err != nil {
			return nil, err
		}

		account := entities.AddressFromString(string(subnet.Account))
		publicKeyBytes, err := base64.RawStdEncoding.DecodeString(subnet.SignatureData.PublicKey)

		if err != nil {
			return nil, err
		}
		authMsg := fmt.Sprintf(constants.SignatureMessageString, "CreateSubnet", addressPrefix, subnet.Ref, encoder.ToBase64Padded(msg))
		logger.Info("MSG:: ", authMsg)
		valid, err = crypto.VerifySignatureAmino(encoder.ToBase64Padded([]byte(authMsg)), decodedSig, account.Addr, publicKeyBytes)
		if err != nil {
			return nil, err
		}

	}
	if !valid {
		return nil, apperror.Unauthorized("Invalid signature signer")
	}
	query.GetOne(models.SubnetState{Subnet: entities.Subnet{Ref: subnet.Ref}}, currentSubnetState)
	return currentSubnetState, nil
}

func HandleNewPubSubSubnetEvent(event *entities.Event, ctx *context.Context) {
	logger.WithFields(logrus.Fields{"event": event}).Debug("New Subnet event from pubsub channel")
	markAsSynced := false
	updateState := false
	var eventError string
	// hash, _ := event.GetHash()
	err := ValidateEvent(*event)

	if err != nil {
		logger.Error(err)
		return
	}

	logger.Infof("Event is a valid event %s", event.PayloadHash)
	cfg, _ := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)

	// Extract and validate the Data of the paylaod which is an Events Payload Data,
	data := event.Payload.Data.(*entities.Subnet)
	hash, _ := data.GetHash()
	data.Hash = hex.EncodeToString(hash)
	// authEventHash := event.AuthEventHash
	// authState, authError := query.GetOneAuthorizationState(entities.Authorization{Event: authEventHash})
	logger.Info("data.Meta Ref ", data.Meta, " ", data.Ref)
	h, _ := data.GetHash()
	logger.Infof("data.Hash %v", h)

	currentState, err := ValidateSubnetData(data, cfg.AddressPrefix)
	if err != nil {
		// penalize node for broadcasting invalid data
		logger.Infof("Invalid Subnet data %v. Node should be penalized", err)
		return
	}

	// check if we are upto date on this event
	prevEventUpToDate := query.EventExist(&event.PreviousEventHash) || (currentState == nil && event.PreviousEventHash.Hash == "") || (currentState != nil && currentState.Event.Hash == event.PreviousEventHash.Hash)
	// authEventUpToDate := query.EventExist(&event.AuthEventHash) || (authState == nil && event.AuthEventHash.Hash == "") || (authState != nil && authState.Event == authEventHash)

	// Confirm if this is an older event coming after a newer event.
	// If it is, then we only have to update our event history, else we need to also update our current state
	isMoreRecent := false
	if currentState != nil && currentState.Hash != data.Hash {
		var currentStateEvent = &models.SubnetEvent{}
		_ = query.GetOne(entities.Event{Hash: currentState.Event.Hash}, currentStateEvent)
		isMoreRecent, markAsSynced = IsMoreRecent(
			currentStateEvent.ID,
			currentState.Event.Hash,
			currentStateEvent.Payload.Timestamp,
			event.Hash,
			event.Payload.Timestamp,
			markAsSynced,
		)
		// if uint64(currentStateEvent.Payload.Timestamp) < uint64(event.Payload.Timestamp) {
		// 	isMoreRecent = true
		// }
		// if uint64(currentStateEvent.Payload.Timestamp) > uint64(event.Payload.Timestamp) {
		// 	isMoreRecent = false
		// }
		// // if the authorization was created at exactly the same time but their hash is different
		// // use the last 4 digits of their event hash
		// if uint64(currentStateEvent.Payload.Timestamp) == uint64(event.Payload.Timestamp) {
		// 	// get the event payload of the current state

		// 	if err != nil && err != gorm.ErrRecordNotFound {
		// 		logger.Error("DB error", err)
		// 	}
		// 	if currentStateEvent.ID == "" {
		// 		markAsSynced = false
		// 	} else {
		// 		// if currentStateEvent.Payload.Timestamp < event.Payload.Timestamp {
		// 		// 	isMoreRecent = true
		// 		// }
		// 		// if currentStateEvent.Payload.Timestamp == event.Payload.Timestamp {
		// 			// logger.Infof("Current state %v", currentStateEvent.Payload)
		// 			csN := new(big.Int)
		// 			csN.SetString(currentState.Event.Hash[56:], 16)
		// 			nsN := new(big.Int)
		// 			nsN.SetString(event.Hash[56:], 16)

		// 			if csN.Cmp(nsN) < 1 {
		// 				isMoreRecent = true
		// 			}
		// 		//}
		// 	}
		// }
	}

	// If no error, then we should act accordingly as well
	// If are upto date, then we should update the state based on if its a recent or old event
	if len(eventError) == 0 {
		if prevEventUpToDate { // we are upto date
			if currentState == nil || isMoreRecent {
				updateState = true
				markAsSynced = trupe
			} else {
				// Its an old event
				markAsSynced = true
				updateState = false
			}
		} else {
			updateState = false
			markAsSynced = false
		}

	}

	// Save stuff permanently
	tx := sql.Db.Begin()
	logger.Info(":::::updateState: Db Error", updateState, currentState == nil)

	// If the event was not signed by your node
	if string(event.Validator) != (*cfg).NetworkPublicKey {
		// save the event
		event.Error = eventError
		event.IsValid = markAsSynced && len(eventError) == 0.
		event.Synced = markAsSynced
		event.Broadcasted = true
		_, _, err := query.SaveRecord(models.SubnetEvent{
			Event: entities.Event{
				PayloadHash: event.PayloadHash,
			},
		}, models.SubnetEvent{
			Event: *event,
		}, false, tx)
		if err != nil {
			tx.Rollback()
			logger.Error("1000: Db Error", err)
			return
		}
	} else {
		if markAsSynced {
			_, _, err := query.SaveRecord(models.SubnetEvent{
				Event: entities.Event{PayloadHash: event.PayloadHash},
			}, models.SubnetEvent{
				Event: entities.Event{Synced: true, Broadcasted: true, Error: eventError, IsValid: len(eventError) == 0},
			}, true, tx)
			if err != nil {
				logger.Error("DB error", err)
			}
		} else {
			// mark as broadcasted
			_, _, err := query.SaveRecord(models.SubnetEvent{
				Event: entities.Event{PayloadHash: event.PayloadHash, Broadcasted: false},
			},
				models.SubnetEvent{
					Event: entities.Event{Broadcasted: true},
				}, true, tx)
			if err != nil {
				logger.Error("DB error", err)
			}
		}
	}

	// d, err := event.Payload.EncodeBytes()
	if err != nil {
		logger.Errorf("Invalid event payload")
	}
	// agent, err := crypto.GetSignerECC(&d, &event.Payload.Signature)
	if err != nil {
		logger.Errorf("Invalid event payload")
	}
	data.Event = *entities.NewEventPath(event.Validator, entities.SubnetEventModel, event.Hash)

	data.Account = event.Payload.Account
	// logger.Error("data.Public ", data.Public)

	if updateState {
		_, _, err := query.SaveRecord(models.SubnetState{
			Subnet: entities.Subnet{ID: data.ID},
		}, models.SubnetState{
			Subnet: *data,
		}, event.EventType == uint16(constants.UpdateSubnetEvent), tx)
		if err != nil {
			tx.Rollback()
			logger.Error("7000: Db Error", err)
			return
		}
	}
	tx.Commit()

	if string(event.Validator) != (*cfg).NetworkPublicKey {
		dependent, err := query.GetDependentEvents(*event)
		if err != nil {
			logger.Info("Unable to get dependent events", err)
		}
		for _, dep := range *dependent {
			go HandleNewPubSubSubnetEvent(&dep, ctx)
		}
	}

	// TODO Broadcast the updated state
}
