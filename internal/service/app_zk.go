package service

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	dsquery "github.com/mlayerprotocol/go-mlayer/internal/ds/query"
	"github.com/mlayerprotocol/go-mlayer/zk/circuits"
	// query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
)

/*
Validate an agent authorization
*/
func ValidateApplicationDataZK(clientPayload *entities.ClientPayload, chainID configs.ChainId) (appState *entities.Application, err error) {
	// check fields of Application

	var currentApplicationState *entities.Application
	app := clientPayload.Data.(entities.Application)

	if len(app.Ref) > 64 {
		return nil, apperror.BadRequest("Application ref cannot be more than 64 characters")
	}

	if len(app.Ref) > 0 && !utils.IsAlphaNumericDot(app.Ref) {
		return nil, apperror.BadRequest("Ref can only include alpha-numerics, and .")
	}
	if strings.Contains(strings.ToLower(app.Ref), "global") || strings.Contains(strings.ToLower(app.Ref), "giobal") {
		return nil, apperror.BadRequest("Application ref cannot contain word \"global\"")
	}

	if app.ZKData != nil {

		// verify ZK
		proof := groth16.NewProof(circuits.CURVE)
		proofBuf := bytes.NewReader([]byte(app.ZKData.Proof))
		_, err := proof.ReadFrom(proofBuf)
		if err != nil {
			logger.Fatal(err)
		}
		cd := circuits.AppCircuitData[app.ZKData.Version]
		publicWitnessData := circuits.AppCircuit{
			Public: circuits.AppCircuitPublic{
				ID:            app.Commitment,
				AppCommitment: app.Commitment,
				ChainID:       chainID,
				// Name:                 app.Name,
				Meta:                 []byte(app.Meta),
				Ref:                  []byte(app.Ref),
				Categories:           []byte(app.Categories),
				Status:               new(big.Int).SetUint64(uint64(*app.Status)).Bytes(),
				DefaultAuthPrivilege: new(big.Int).SetUint64(uint64(*app.DefaultAuthPrivilege)).Bytes(),
				Timestamp:            new(big.Int).SetUint64(app.Timestamp).Bytes(),
			},

			// // Public inputs (placeholders)
			// CreatorCommitment: 0,

			// NewName:           0,
		}
		witness, err := frontend.NewWitness(&publicWitnessData, circuits.CURVE.ScalarField())
		if err != nil {
			logger.Errorf("Error creating witness: %v\n", err)
			return nil, err
		}
		groth16.Verify(proof, cd.VerificationKey, witness)
	} else {

		agent, err := entities.DeviceFromString(string(app.AppKey))
		if err != nil {
			return nil, fmt.Errorf("invalid agent")
		}
		account, err := entities.AccountFromString(string(app.Account))
		if err != nil {
			return nil, fmt.Errorf("invalid account")
		}

		logger.Infof("APPID %s", app.ID)
		if len(app.AppKey) > 0 && app.ID != "" {

			// TODO Check that this agent is an admin of app. Return error if not
			priv := constants.AdminPriviledge

			// err := query.GetOne(models.AuthorizationState{Authorization: entities.Authorization{
			// 	Agent: agent.ToDeviceString(),
			// 	Application: app.ID,
			// 	Priviledge: &priv,
			// 	Account: account.ToString(),
			// }}, &auth)
			authorizations, err := dsquery.GetAccountAuthorizations(entities.Authorization{
				Authorized:  agent.ToAddressString(),
				Application: app.ID,
				Account:     account.ToString(),
			}, dsquery.DefaultQueryLimit, nil)
			if err != nil {
				if dsquery.IsErrorNotFound(err) {
					return nil, apperror.Unauthorized("agent not authorized")
				}
				return nil, apperror.Internal("internal database error")
			}
			authorized := false
			// var auth models.AuthorizationState
			for _, _auth := range authorizations {
				if *(_auth.Priviledge) == priv {
					authorized = true
				}
				// auth = models.AuthorizationState{Authorization: *_auth}
			}
			if !authorized {
				return nil, apperror.Unauthorized("agent not authorized")
			}

		}

		// TODO if agent is specified, ensure agent is allowed to sign on behalf of Owner

		var valid bool
		// b, _ := app.EncodeBytes()
		msg, err := clientPayload.GetHash()
		if err != nil {
			return nil, err
		}
		logger.Infof("HELLOSJSLIJSDMSG: %s", hex.EncodeToString(msg))
		action := "write_app"
		switch app.SignatureData.Type {
		case entities.EthereumPubKey:
			authMsg := fmt.Sprintf(constants.SignatureMessageString, action, app.Ref, chainID, encoder.ToBase64Padded(msg))
			msgByte := crypto.EthMessage([]byte(authMsg))
			logger.Infof("AUTHMESSAGE %s", authMsg)
			addr, err := entities.AddressFromString(string(app.Commitment))
			if err != nil {
				return nil, apperror.BadRequest("invalid account address")
			}
			valid = crypto.VerifySignatureECC(addr.Addr, &msgByte, string(app.SignatureData.Signature))

		case entities.TendermintsSecp256k1PubKey:

			decodedSig, err := base64.StdEncoding.DecodeString(string(app.SignatureData.Signature))
			if err != nil {
				return nil, err
			}
			// account := entities.AddressFromString(string(app.Account))
			publicKeyBytes, err := base64.RawStdEncoding.DecodeString(string(app.SignatureData.PublicKey))

			if err != nil {
				return nil, err
			}
			authMsg := fmt.Sprintf(constants.SignatureMessageString, action, chainID, app.Ref, encoder.ToBase64Padded(msg))
			logger.Debug("MSG:: ", authMsg)
			valid, err = crypto.VerifySignatureAmino(encoder.ToBase64Padded([]byte(authMsg)), decodedSig, account.Addr, publicKeyBytes)
			if err != nil {
				return nil, err
			}
		}
		if !valid {
			return nil, apperror.Unauthorized("Invalid app data signature")
		}

	}
	if app.ID != "" {
		// var  curSt  models.ApplicationState
		// query.GetOne(models.ApplicationState{Application: entities.Application{ID: app.ID}}, &curSt)
		currentApplicationState, err = dsquery.GetApplicationStateById(app.ID)
		if err != nil {
			if !dsquery.IsErrorNotFound(err) {
				return nil, err
			} else {
				return nil, nil
			}
		}

	}
	// logger.Infof("IsValidSigner %v, subId: %s, currentstate: %v, error: %v", valid, app.ID, currentApplicationState, err)
	// logger.Infof("IsValidSigner %v, subId: %s, currentstate: %v, error: %v", valid, app.ID, currentApplicationState, err)
	return currentApplicationState, nil
}
