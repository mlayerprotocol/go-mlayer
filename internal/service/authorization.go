package service

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/channelpool"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ValidateAuthData(auth *entities.Authorization, addressPrefix string) (prevAuthState *models.AuthorizationState, grantorAuthState *models.AuthorizationState, err error) {
	b, err := auth.EncodeBytes()
	if err != nil {
		return nil, nil, err
	}
	var valid bool

	// if string(auth.Account) == string(auth.Grantor) {
	// 	valid, err = crypto.VerifySignatureEDD(string(auth.Account), &b, auth.Signature );
	// 	if err != nil {
	// 		return nil, nil, err
	// 	}
	// } else {
	// 	valid = crypto.VerifySignatureECC(string(auth.Grantor), &b, auth.Signature );
	// 	if valid {
	// 		// check if the grantor is authorized
	// 		grantorAuthState, err = query.GetOneAuthorizationState(entities.Authorization{Account: auth.Account, Agent: string(auth.Grantor)})
	// 		if err == gorm.ErrRecordNotFound {
	// 			return nil, nil, apperror.Unauthorized( "Grantor not authorized agent")
	// 		}
	// 		if grantorAuthState.Authorization.Priviledge != constants.AdminPriviledge {
	// 			return nil, grantorAuthState, apperror.Forbidden(" Grantor does not have enough permission")
	// 		}

	// 	}
	// }
	switch auth.SignatureData.Type {
	case entities.EthereumPubKey:
		valid = crypto.VerifySignatureECC(string(auth.Grantor), &b, auth.SignatureData.Signature)
		if valid {
			// check if the grantor is authorized
			grantorAuthState, err = query.GetOneAuthorizationState(entities.Authorization{Account: auth.Account, Agent: string(auth.Grantor)})
			if err == gorm.ErrRecordNotFound {
				return nil, nil, apperror.Unauthorized("Grantor not authorized agent")
			}
			if grantorAuthState.Authorization.Priviledge != constants.AdminPriviledge {
				return nil, grantorAuthState, apperror.Forbidden(" Grantor does not have enough permission")
			}
		}

	case entities.TendermintsSecp256k1PubKey:

		decodedSig, err := base64.StdEncoding.DecodeString(auth.SignatureData.Signature)
		if err != nil {
			return nil, nil, err
		}

		msg, err := auth.GetHash()

		logger.Info("MSG:: ", msg)

		if err != nil {
			return nil, nil, err
		}

		grantorAddress, err := entities.AddressFromString(auth.Grantor)
		publicKeyBytes, err := base64.RawStdEncoding.DecodeString(auth.SignatureData.PublicKey)

		if err != nil {
			return nil, nil, err
		}
		// grantor, err := entities.AddressFromString(auth.Grantor)

		// if err != nil {
		// 	return nil, nil, err
		// }

		// decoded, err := hex.DecodeString(grantor.Addr)
		// if err == nil {
		// 	address = crypto.ToBech32Address(decoded, "cosmos")
		// }

		authMsg := fmt.Sprintf("Approve %s for %s: %s", auth.Agent, addressPrefix, encoder.ToBase64Padded(msg))
		logger.Info("AUTHMESS ", authMsg, " ", auth.Hash)

		valid, err = crypto.VerifySignatureAmino(encoder.ToBase64Padded([]byte(authMsg)), decodedSig, grantorAddress.Addr, publicKeyBytes)
		if err != nil {
			return nil, nil, err
		}

	}

	prevAuthState, err = query.GetOneAuthorizationState(entities.Authorization{Account: auth.Account, Agent: auth.Agent})
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, nil, err
	}
	if !valid {
		return prevAuthState, grantorAuthState, errors.New("4000: Invalid signature signer")
	}
	return prevAuthState, grantorAuthState, nil

}

func HandleNewPubSubAuthEvent(event *entities.Event, ctx *context.Context) {
	logger.WithFields(logrus.Fields{"event": event}).Debug("New auth event from pubsub channel")
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
	authRequest := event.Payload.Data.(*entities.Authorization)
	hash, _ := authRequest.GetHash()
	authRequest.Hash = hex.EncodeToString(hash)
	currentState, authState, stateError := ValidateAuthData(authRequest, cfg.AddressPrefix)

	// check if we are upto date on this event
	prevEventUpToDate := (currentState == nil && event.PreviousEventHash.Hash == "") || (currentState != nil && currentState.EventHash == event.PreviousEventHash.Hash)
	authEventUpToDate := (authState == nil && event.AuthEventHash.Hash == "") || (authState != nil && authState.EventHash == event.AuthEventHash.Hash)

	// Confirm if this is an older event coming after a newer event.
	// If it is, then we only have to update our event history, else we need to also update our current state
	isMoreRecent := false
	if currentState != nil && currentState.Hash != authRequest.Hash {
		if currentState.Timestamp < authRequest.Timestamp {
			isMoreRecent = true
		}
		if currentState.Timestamp > authRequest.Timestamp {
			isMoreRecent = false
		}
		// if the authorization was created at exactly the same time but their hash is different
		// use the last 4 digits of their event hash
		if currentState.Timestamp == authRequest.Timestamp {
			// get the event payload of the current state
			currentStateEvent, err := query.GetOneAuthorizationEvent(entities.Event{Hash: currentState.EventHash})
			if err != nil && err != gorm.ErrRecordNotFound {
				logger.Fatal("DB error", err)
			}
			if currentStateEvent == nil {
				markAsSynced = false
			} else {
				if currentStateEvent.Payload.Timestamp < event.Payload.Timestamp {
					isMoreRecent = true
				}
				if currentStateEvent.Payload.Timestamp == event.Payload.Timestamp {
					// logger.Infof("Current state %v", currentStateEvent.Payload)
					csN := new(big.Int)
					csN.SetString(currentState.EventHash[56:], 16)
					nsN := new(big.Int)
					nsN.SetString(event.Hash[56:], 16)

					if csN.Cmp(nsN) < 1 {
						isMoreRecent = true
					}
				}
			}
		}
	}
	if stateError != nil {
		// check if we are upto date. If we are, then the error is an actual one
		// the error should be attached when saving the event
		// But if we are not upto date, then we might need to wait for more info from the network

		if prevEventUpToDate && authEventUpToDate {
			// we are upto date. This is an actual error. No need to expect an update from the network
			eventError = stateError.Error()
			markAsSynced = true
		} else {
			if currentState == nil || (currentState != nil && isMoreRecent) { // it is a morer ecent event
				if strings.HasPrefix(stateError.Error(), constants.ErrorForbidden) || strings.HasPrefix(stateError.Error(), constants.ErrorUnauthorized) {
					markAsSynced = false
				} else {
					// entire event can be considered bad since the payload data is bad
					// this should have been sorted out before broadcasting to the network
					// TODO penalize the node that broadcasted this
					eventError = stateError.Error()
					markAsSynced = true
				}

			} else {
				// we are upto date. We just need to store this event as well.
				// No need to update state
				markAsSynced = true
				eventError = stateError.Error()
			}
		}

	}

	// If no error, then we should act accordingly as well
	// If are upto date, then we should update the state based on if its a recent or old event
	if len(eventError) == 0 {
		if prevEventUpToDate && authEventUpToDate { // we are upto date
			if currentState == nil || (currentState != nil && isMoreRecent) {
				updateState = true
				markAsSynced = true
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

	// If the event was not signed by your node
	if string(event.Validator) != (*cfg).NetworkPublicKey {
		// save the event
		event.Error = eventError
		event.IsValid = markAsSynced && len(eventError) == 0.
		event.Synced = markAsSynced
		event.Broadcasted = true
		_, _, err := query.SaveAuthorizationEvent(event, true, tx)
		if err != nil {
			tx.Rollback()
			logger.Fatal("5000: Db Error", err)
			return
		}
	} else {
		if markAsSynced {
			_, err := query.UpdateAuthorizationEvent(entities.Event{PayloadHash: event.PayloadHash},
				entities.Event{Synced: true, Broadcasted: true, Error: eventError, IsValid: len(eventError) == 0}, tx)

			if err != nil {
				logger.Fatal("DB error", err)
			}
		} else {
			// mark as broadcasted
			_, err := query.UpdateAuthorizationEvent(entities.Event{PayloadHash: event.PayloadHash, Broadcasted: false},
				entities.Event{Broadcasted: true}, tx)
			if err != nil {
				logger.Fatal("DB error", err)
			}
		}
	}
	authRequest.EventHash = event.Hash
	if updateState {
		_, err := query.SaveAuthorizationState(authRequest, tx)
		if err != nil {
			tx.Rollback()
			logger.Fatal("5000: Db Error", err)
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
			channelpool.AuthorizationEventPublishC <- &dep.Event
		}
	}

	// TODO Broadcast the updated state

}
