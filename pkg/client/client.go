package client

import (
	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"gorm.io/gorm"
)

func ValidateClientPayload(
	payload *entities.ClientPayload,
	strictAuth bool,
	chainId configs.ChainId,
) (*models.AuthorizationState, *entities.DeviceString, error) {
	
	// _, err := payload.EncodeBytes()
	// if err != nil {
	// 	logger.Error(err)
	// 	return nil, apperror.Internal(err.Error())
	// }
	// logger.Info("ENCODEDBYTESSS"," ", hex.EncodeToString(d), " ", hex.EncodeToString(crypto.Keccak256Hash(d)))

	if payload.Subnet == "" {
		return nil, nil, apperror.Forbidden("Subnet Id is required")
	}
	if string(payload.ChainId) != string(chainId) {
		return nil, nil, apperror.Forbidden("Invalid chain Id")
	}
	// payload.ChainId = chainId
	agent, err := payload.GetSigner()
	
	if err != nil {
		return nil, nil, err
	}
	
	// if agent != payload.Agent {
	// 	return nil, nil, apperror.BadRequest("Agent is required")
	// }
	// logger.Infof("AGENTTTT %s", agent)
	subnet := models.SubnetState{}
	err = query.GetOne(models.SubnetState{Subnet: entities.Subnet{ID: payload.Subnet}}, &subnet)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil,  nil, apperror.Forbidden("Invalid subnet id")
		}
		return nil, nil, apperror.Internal(err.Error())
	}
	if *subnet.Status ==  0 {
		return nil, nil, apperror.Forbidden("Subnet is disabled")
	}


	// check if device is authorized
	if agent != "" {
		logger.Infof("New Event for Agent/Device %s", agent)
		agent = entities.DeviceString(agent)
		
		if strictAuth || string(payload.Account) != ""  {
			authData := models.AuthorizationState{}
			err := query.GetOne(models.AuthorizationState{
				Authorization: entities.Authorization{Account: payload.Account,
					Subnet: payload.Subnet,
					Agent:  agent},
			}, &authData)
			if err != nil {
				// if err == gorm.ErrRecordNotFound {
				// 	return nil, nil
				// }
				return nil, &agent, err
			} else {
				return &authData, &agent,  nil
			}
		} else {
			return nil, &agent,  nil
		}
	}
	return nil, nil, apperror.BadRequest("Unable to resolve agent")
}

func SyncRequest(payload *entities.ClientPayload) entities.SyncResponse {
	var response = entities.SyncResponse{}
	return response
}
