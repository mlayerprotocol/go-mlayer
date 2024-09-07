package service

import (
	"context"
	"encoding/base64"
	"encoding/hex"
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
	"github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"gorm.io/gorm"
)

func ValidateAuthPayloadData(clientPayload *entities.ClientPayload, chainId configs.ChainId) (prevAuthState *models.AuthorizationState, grantorAuthState *models.AuthorizationState, subnet *models.SubnetState, err error) {
	auth  := clientPayload.Data.(entities.Authorization)
	if err != nil {
		return nil, nil, nil, err
	}
	logger.Debug("auth.SignatureData.Signature:: ", auth.SignatureData.Signature)
	var valid bool
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
	msg, err := clientPayload.GetHash()
	if err != nil {
		return nil, nil, nil, err
	}
	switch auth.SignatureData.Type {
	case entities.EthereumPubKey:
		authMsg := fmt.Sprintf(constants.SignatureMessageString, "AuthorizeAgent", chainId, agent.Addr, encoder.ToBase64Padded(msg))
		logger.Debug("MSG:: ", authMsg)
		
		msgByte := crypto.EthMessage([]byte(authMsg))
		signer := utils.IfThenElse(len(string(auth.Grantor)) == 0, account.Addr,  grantor.Addr )
		// if len(agent.Addr) > 0 {
		// 	signer = agent.Addr
		// } 
		// logger.Debug("Signer:: ", signer)
		valid = crypto.VerifySignatureECC(signer, &msgByte, auth.SignatureData.Signature)
		if valid {
			// check if the grantor is authorized
			if auth.Grantor != auth.Account {
				grantorAuthState, err = query.GetOneAuthorizationState(entities.Authorization{Account: entities.DIDString(string(account.ToDeviceString())), Subnet: auth.Subnet, Agent: grantor.ToDeviceString()})
				if err == gorm.ErrRecordNotFound {
					return nil, nil, subnet, apperror.Unauthorized("Grantor not authorized agent")
				}
				if *grantorAuthState.Authorization.Priviledge != constants.AdminPriviledge {
					return nil, grantorAuthState,  subnet, apperror.Forbidden(" Grantor does not have enough permission")
				}
			}
			// check if agent is authorized by grantor
			// if agent.Addr != "" {
			// 	agentAuthState, err := query.GetOneAuthorizationState(entities.Authorization{Account: entities.DIDString(grantor.ToDeviceString()), Subnet: auth.Subnet, Agent: agent.ToDeviceString()})
			// 	if err == gorm.ErrRecordNotFound {
			// 		return nil, nil, subnet, apperror.Unauthorized("Agent not authorized to act on behalf of grantor")
			// 	}
			// 	if *agentAuthState.Authorization.Priviledge != constants.AdminPriviledge {
			// 		return nil, grantorAuthState,  subnet, apperror.Forbidden("Agent does not have enough permission")
			// 	}
			// }
		}

	case entities.TendermintsSecp256k1PubKey:

		decodedSig, err := base64.StdEncoding.DecodeString(auth.SignatureData.Signature)
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
		authMsg := fmt.Sprintf(constants.SignatureMessageString, "AuthorizeAgent", chainId, agent.Addr, encoder.ToBase64Padded(msg))

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

func saveAuthorizationEvent(where entities.Event, createData *entities.Event, updateData *entities.Event, tx *gorm.DB) (*entities.Event, error) {
	var createModel *models.AuthorizationEvent
	if createData != nil {
		createModel = &models.AuthorizationEvent{Event: *createData}
	} else {
		createModel = &models.AuthorizationEvent{}
	}
	var updateModel *models.AuthorizationEvent
	if updateData != nil {
		updateModel = &models.AuthorizationEvent{Event: *updateData}
	}
	model, _, err := query.SaveRecord(models.AuthorizationEvent{Event: where},  createModel, updateModel, tx)
	if err != nil {
		return nil, err
	}
	return &model.Event, err
}


func HandleNewPubSubAuthEvent(event *entities.Event, ctx *context.Context) {
	
	cfg, ok := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)
	if !ok {
		panic("Unable to load config from context")
	}
	data := event.Payload.Data.(entities.Authorization)
	// var id = data.ID
	// if len(data.ID) == 0 {
	// 	id, _ = entities.GetId(data)
	// } else {
	// 	id = data.ID
	// }
	data.BlockNumber = event.BlockNumber
	data.Cycle = event.Cycle
	data.Epoch = event.Epoch
	data.Event = *event.GetPath()
	hash, err := data.GetHash()
	data.Agent = entities.AddressFromString(string(data.Agent)).ToDeviceString()
	if err != nil {
		return
	}
	data.Hash = hex.EncodeToString(hash)
	var subnet = data.Subnet
	
	var localState models.AuthorizationState
	// err := query.GetOne(&models.TopicState{Topic: entities.Topic{ID: id}}, &localTopicState)
	err = sql.SqlDb.Where(&models.AuthorizationState{Authorization: entities.Authorization{Subnet: subnet, Agent: entities.AddressFromString(string(data.Agent)).ToDeviceString()}}).Take(&localState).Error
	if err != nil {
		logger.Error(err)
	}
	
	
	var localDataState *LocalDataState
	if localState.ID != "" {
		localDataState = &LocalDataState{
			ID: localState.ID,
			Hash: localState.Hash,
			Event: &localState.Event,
			Timestamp: *localState.Timestamp,
		}
	}
	// localDataState := utils.IfThenElse(localTopicState != nil, &LocalDataState{
	// 	ID: localTopicState.ID,
	// 	Hash: localTopicState.Hash,
	// 	Event: &localTopicState.Event,
	// 	Timestamp: localTopicState.Timestamp,
	// }, nil)
	var stateEvent *entities.Event
	if localState.ID != "" {
		stateEvent, err = query.GetEventFromPath(&localState.Event)
		if err != nil && err != query.ErrorNotFound {
			logger.Debug(err)
		}
	}
	var localDataStateEvent *LocalDataStateEvent
	if stateEvent != nil {
		localDataStateEvent = &LocalDataStateEvent{
			ID: stateEvent.ID,
			Hash: stateEvent.Hash,
			Timestamp: stateEvent.Timestamp,
		}
	}

	eventData := PayloadData{Subnet: subnet, localDataState: localDataState, localDataStateEvent:  localDataStateEvent}
	tx := sql.SqlDb
	// defer func () {
	// 	if tx.Error != nil {
	// 		tx.Rollback()
	// 	} else {
	// 		tx.Commit()
	// 	}
	// }()
	
	previousEventUptoDate,  authEventUpToDate, _, eventIsMoreRecent, err := ProcessEvent(event,  eventData, false, saveAuthorizationEvent, tx, ctx)
	if err != nil {
		logger.Debugf("Processing Error...: %v", err)
		return
	}
	logger.Debugf("Processing 2...: %v,  %v", previousEventUptoDate, authEventUpToDate)
	if previousEventUptoDate  && authEventUpToDate {
		_, _, _, err = ValidateAuthPayloadData(&event.Payload, cfg.ChainId)
		if err != nil {
			// update error and mark as synced
			// notify validator of error
			saveAuthorizationEvent(entities.Event{Hash: event.Hash}, nil, &entities.Event{Error: err.Error(), IsValid: false, Synced: true}, tx )
			
		} else {
			// TODO if event is older than our state, just save it and mark it as synced
			
			savedEvent, err := saveAuthorizationEvent(entities.Event{Hash: event.Hash}, nil, &entities.Event{IsValid: true, Synced: true}, tx );
			if eventIsMoreRecent && err == nil {
				// update state
				_, _, err := query.SaveRecord(models.AuthorizationState{
					Authorization: entities.Authorization{Subnet: subnet, Agent: data.Agent},
				}, &models.AuthorizationState{
					Authorization: data,
				},  &models.AuthorizationState{
					Authorization: data,
				}, tx)
				if err != nil {
					// tx.Rollback()
					logger.Errorf("SaveStateError %v", err)
					return
				}
				
			}
			if err == nil {
				go OnFinishProcessingEvent(ctx, event.GetPath(), &savedEvent.Payload.Subnet, err)
			}
			
			
			if string(event.Validator) != cfg.PublicKey {
				go func () {
				dependent, err := query.GetDependentEvents(event)
				if err != nil {
					logger.Debug("Unable to get dependent events", err)
				}
				for _, dep := range *dependent {
					logger.Debugf("Processing Dependend Event %s", dep.Hash)
					go HandleNewPubSubEvent(&dep, ctx)
				}
				}()
			}
			
		}
	}
}


// {"action":"AuthorizeAgent","network":"84532","identifier":"0xD466f0C2506b69e091b4356cd55b55f6DF00491b","hash":"kXTFkj7NkzQt5VQKvTcxXq6RHd5KvhO7eEDxtcsy+Ec="}