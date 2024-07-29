package service

import (
	"context"
	"encoding/base64"
	"errors"
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
	"gorm.io/gorm"
)

func ValidateAuthPayloadData(auth *entities.Authorization, addressPrefix configs.ChainId) (prevAuthState *models.AuthorizationState, grantorAuthState *models.AuthorizationState, subnet *models.SubnetState, err error) {
	
	b, err := auth.EncodeBytes()
	if err != nil {
		return nil, nil, nil, err
	}
	logger.Info("auth.SignatureData.Signature:: ", auth.SignatureData.Signature)
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
	if auth.Subnet == "" {
		return nil, nil, nil, apperror.BadRequest("Subnet is required")
	}
	

	// TODO find subnets state prior to the current state
	err = query.GetOne(models.SubnetState{Subnet: entities.Subnet{ID: auth.Subnet}}, &subnet)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil, nil, apperror.NotFound("subnet not found")
		}
		return nil, nil, nil, apperror.Internal(err.Error())
	}

	if auth.Account != subnet.Account && *auth.Priviledge > *subnet.DefaultAuthPrivilege {
		return nil, nil, subnet, apperror.Internal("invalid auth priviledge. Cannot be higher than subnets default")
	}
	account :=  entities.AddressFromString(string(auth.Account))
	grantor := entities.AddressFromString(string(auth.Grantor))
	agent := entities.AddressFromString(string(auth.Agent))
	if account.Addr == agent.Addr {
		return nil, nil, subnet, apperror.Internal("cannot reassign subnet owner role")
	}

	switch auth.SignatureData.Type {
	case entities.EthereumPubKey:
		signer := utils.IfThenElse(len(string(auth.Grantor)) == 0, account.Addr,  grantor.Addr )
		valid = crypto.VerifySignatureECC(signer, &b, auth.SignatureData.Signature)
		if valid {
			// check if the grantor is authorized
			if auth.Grantor != auth.Account {
				grantorAuthState, err = query.GetOneAuthorizationState(entities.Authorization{Account: auth.Account, Subnet: auth.Subnet, Agent: grantor.ToDeviceString()})
				if err == gorm.ErrRecordNotFound {
					return nil, nil, subnet, apperror.Unauthorized("Grantor not authorized agent")
				}
				if *grantorAuthState.Authorization.Priviledge != constants.AdminPriviledge {
					return nil, grantorAuthState,  subnet, apperror.Forbidden(" Grantor does not have enough permission")
				}
			}
		}

	case entities.TendermintsSecp256k1PubKey:

		decodedSig, err := base64.StdEncoding.DecodeString(auth.SignatureData.Signature)
		if err != nil {
			return nil, nil, subnet, err
		}

		msg, err := auth.GetHash()

		logger.Info("MSG:: ", msg)

		if err != nil {
			return nil, nil, subnet, err
		}
		publicKeyBytes, err := base64.RawStdEncoding.DecodeString(auth.SignatureData.PublicKey)

		if err != nil {
			return nil, nil, subnet, err
		}
		// grantor, err := entities.AddressFromString(auth.Grantor)

		// if err != nil {
		// 	return nil, nil, err
		// }

		// decoded, err := hex.DecodeString(grantor.Addr)
		// if err == nil {
		// 	address = crypto.ToBech32Address(decoded, "cosmos")
		// }
		authMsg := fmt.Sprintf(constants.SignatureMessageString, "AuthorizeAgent", addressPrefix, agent.Addr, encoder.ToBase64Padded(msg))

		valid, err = crypto.VerifySignatureAmino(encoder.ToBase64Padded([]byte(authMsg)), decodedSig, grantor.Addr, publicKeyBytes)
		if err != nil {
			return nil, nil, subnet, err
		}

	}

	prevAuthState, err = query.GetOneAuthorizationState(entities.Authorization{Agent:  agent.ToDeviceString(), Subnet: auth.Subnet})
	
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, nil, subnet, err
	}
	if !valid {
		return prevAuthState, grantorAuthState, subnet, errors.New("4000: Invalid authorization data signature")
	}
	
	return prevAuthState, grantorAuthState, subnet, nil

}

func HandleNewPubSubAuthEvent(event *entities.Event, ctx *context.Context) {
	
	

}
