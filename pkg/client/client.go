package client

import (
	"encoding/hex"
	"os"

	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/chain"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/p2p"
	"gorm.io/gorm"
)

type NodeInfo struct {
	Account string `json:"account"` 
	NodeType constants.NodeType `json:"node_type"` 
	NodePublicKey string `json:"node_pubkey"` 
	// ChainPublicKey string `json:"chain_pubkey"` 
	ChainId string `json:"chain_id"`
	CurrentCycle uint64 `json:"current_cycle"`
	CurrentBlock uint64 `json:"current_block"`
	CurrentEpoch uint64 `json:"current_epoch"`
	Listeners []string `json:"listeners"`
	Client string `json:"client"`
	ClientVersion string `json:"client_version"`
	ClientReleaseDate string `json:"client_release_date"`
}

func Info(cfg *configs.MainConfiguration) (*NodeInfo, error) {
	provider := chain.Provider(cfg.ChainId)
	info, err := provider.GetChainInfo()
	if err != nil  {
		return  nil, err
	}
	var owner []byte
	if cfg.Validator {
		owner, err = provider.GetValidatorLicenseOwnerAddress(cfg.PublicKeySECP)
	} else {
		owner, err = provider.GetSentryLicenseOwnerAddress(cfg.PublicKeySECP)
	}
	if err != nil {
		return nil, err
	}
	nodeType := constants.ValidatorNodeType
	if !cfg.Validator {
		nodeType = constants.SentryNodeType
	}
	return &NodeInfo{
		Account: hex.EncodeToString(owner),
		NodeType: nodeType,
		NodePublicKey: hex.EncodeToString(cfg.PublicKeySECP),
		//ChainPublicKey: hex.EncodeToString(cfg.PublicKeyEDD),
		ChainId: string(cfg.ChainId),
		Listeners: p2p.GetMultiAddresses(p2p.Host),
		CurrentCycle: info.CurrentCycle.Uint64(),
		CurrentEpoch: info.CurrentEpoch.Uint64(),
		CurrentBlock: info.CurrentBlock.Uint64(),
		Client: "goml",
		ClientVersion: os.Getenv("CLIENT_VERSION"),
		ClientReleaseDate: os.Getenv("RELEASE_DATE"),
	}, nil
	

}
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
	// logger.Debug("ENCODEDBYTESSS"," ", hex.EncodeToString(d), " ", hex.EncodeToString(crypto.Keccak256Hash(d)))

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
	// logger.Debugf("AGENTTTT %s", agent)
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
		logger.Debugf("New Event for Agent/Device %s", agent)
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
