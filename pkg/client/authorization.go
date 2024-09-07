package client

import (
	"encoding/json"
	"time"

	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/service"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"gorm.io/gorm"
)

func getAuthorizations(auth entities.Authorization) (*[]models.AuthorizationState, error) {
	var authState []models.AuthorizationState

	err := query.GetMany(models.AuthorizationState{
		Authorization: auth,
	}, &authState, nil)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &authState, nil
}

func ValidateAuthPayloadData(cfg *configs.MainConfiguration, payload entities.ClientPayload) (assocPrevEvent *entities.EventPath, assocAuthEvent *entities.EventPath, err error) {
	authData := entities.Authorization{}
	
	d, _ := json.Marshal(payload.Data)
	e := json.Unmarshal(d, &authData)
	if e != nil {
		logger.Errorf("UnmarshalError %v", e)
	}
	payload.Data = authData
	if uint64(*authData.Timestamp) == 0 || uint64(*authData.Timestamp) > uint64(time.Now().UnixMilli())+15000 || uint64(*authData.Timestamp) < uint64(time.Now().UnixMilli())-15000 {
		return nil, nil, apperror.BadRequest("Invalid event timestamp")
	}
	if *authData.Duration != 0 && uint64(time.Now().UnixMilli()) >
		(uint64(*authData.Timestamp)+uint64(*authData.Duration)) {
		return nil, nil, apperror.BadRequest("Authorization duration exceeded")
	}
	
	currentState, grantorAuthState, subnet, err := service.ValidateAuthPayloadData(&payload, cfg.ChainId)
	logger.Debugf("CurrentState %v, %v", currentState, subnet)
	// TODO If error is because the subnet was not found, check the dht for the subnet
	if err != nil {
		logger.Error(err)
		return nil, nil, err
	}

	// generate associations
	if currentState != nil {
		
		assocPrevEvent = &currentState.Event
		// assocPrevEvent = entities.EventPath{
		// 	Relationship: entities.PreviousEventAssoc,
		// 	Hash: currentState.Event,
		// 	Model: entities.AuthorizationEventModel,
		// }.ToString()
	} else {
		// Get the subnets state event
		subnetState := &models.SubnetState{}
		err = query.GetOne(&models.SubnetState{Subnet: entities.Subnet{ID: authData.Subnet }}, subnetState)
		if err != nil {
			// find ways to get the subnet
		} else {
			assocPrevEvent = &subnetState.Event
		}


	}
	if grantorAuthState != nil {
		assocAuthEvent = &grantorAuthState.Event
		// assocAuthEvent =  entities.EventPath{
		// 	Relationship: entities.AuthorizationEventAssoc,
		// 	Hash: grantorAuthState.Event,
		// 	Model: entities.AuthorizationEventModel,
		// }
	}
	return assocPrevEvent, assocAuthEvent, nil
}

func GetAccountAuthorizations(auth *entities.Authorization) (*[]models.AuthorizationState, error) {
	//agentAuthState, _ := ValidateClientPayload(clientPayload)

	// if agentAuthState == nil || agentAuthState.Priviledge == 0 {
	// 	return nil, apperror.Unauthorized("Agent not authorized")
	// }
	var authState []models.AuthorizationState
	// // auth.Account = clientPayload.Account

	err := query.GetMany(models.AuthorizationState{
		Authorization: *auth,
	}, &authState, nil)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &authState, nil
}
