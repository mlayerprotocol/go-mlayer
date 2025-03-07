package service

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ipfs/go-datastore"
	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	dsquery "github.com/mlayerprotocol/go-mlayer/internal/ds/query"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/p2p"
	"gorm.io/gorm"
	// query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
)

/*
Validate an agent authorization
*/
func ValidateApplicationData(clientPayload *entities.ClientPayload, chainID configs.ChainId) ( *entities.Application,error) {
	// check fields of Application
	
	var currentApplicationState *entities.Application
	app := clientPayload.Data.(entities.Application)
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
		authorizations, err := dsquery.GetAccountAuthorizations( entities.Authorization{
			Authorized: agent.ToAddressString(),
			Application: app.ID,
			Account: account.ToString(),
		}, dsquery.DefaultQueryLimit, nil)
		if err != nil  {
			if  dsquery.IsErrorNotFound(err) {
				return nil,  apperror.Unauthorized("agent not authorized")
			}
			return nil,  apperror.Internal("internal database error")
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
			return nil,  apperror.Unauthorized("agent not authorized")
		}
		
	}

	// TODO if agent is specified, ensure agent is allowed to sign on behalf of Owner

	if len(app.Ref) > 64 {
		return nil, apperror.BadRequest("Application ref cannot be more than 64 characters")
	}
	
	if len(app.Ref) > 0 && !utils.IsAlphaNumericDot(app.Ref) {
		return nil, apperror.BadRequest("Ref can only include alpha-numerics, and .")
	}
	if strings.Contains(strings.ToLower(app.Ref), "global") ||  strings.Contains(strings.ToLower(app.Ref), "giobal") {
		return nil, apperror.BadRequest("Application ref cannot contain word \"global\"")
	}
	var valid bool
	// b, _ := app.EncodeBytes()
	msg, err := clientPayload.GetHash()
	if err != nil {
		return nil, err
	}
	logger.Infof("HELLOSJSLIJSDMSG: %s", hex.EncodeToString(msg))
	action :=  "write_app"
	switch app.SignatureData.Type {
	case entities.EthereumPubKey:
		authMsg := fmt.Sprintf(constants.SignatureMessageString, action,  app.Ref, chainID, encoder.ToBase64Padded(msg))
		msgByte := crypto.EthMessage([]byte(authMsg))
		logger.Infof("AUTHMESSAGE %s", authMsg)
		addr, err := entities.AddressFromString(string(app.Account))
		if err != nil {
			return nil,  apperror.BadRequest("invalid account address")
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

func saveApplicationEvent(where entities.Event, createData *entities.Event, updateData *entities.Event, txn *datastore.Txn, tx *gorm.DB) (*entities.Event, error) {
	return SaveEvent(entities.ApplicationModel, where, createData, updateData, txn)
 }


func HandleNewPubSubApplicationEvent(event *entities.Event, ctx *context.Context, ) (resp *entities.EventProcessorResponse, err error) {

	cfg, ok := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)
	
	if !ok {
		panic("Unable to load config from context")
	}
	
	dataStates := dsquery.NewDataStates(event.ID, cfg)
	dataStates.AddEvent(*event)
	
	data := event.Payload.Data.(entities.Application)
	data.Event = *event.GetPath()
	data.BlockNumber = event.BlockNumber
	data.Cycle = event.Cycle
	data.Epoch = event.Epoch
	data.EventSignature = event.Signature
	hash, err := data.GetHash()
	if err != nil {
		return nil, err
	}
	data.Hash = hex.EncodeToString(hash)
	logger.Debugf("HandlingNewEvent: %s in app %s", data.ID, event.Payload.Application )
	var id string
	if len(data.ID) == 0 {
		id, _ = entities.GetId(data, data.ID)
	} else {
		id = data.ID
	}

	defer func () {
		if err != nil {
			return
		}
		validators, err := p2p.NewApplicationValidator(cfg, utils.UuidToBytes(data.ID), cfg.PublicKeySECP, data.Cycle)
		if err != nil  {
			return
		}
		dataStates.AddToDhtSync("app", data.ID, validators.MsgPack())
		appRefData, err := p2p.NewApplicationValidator(cfg, utils.UuidToBytes(data.ID), []byte{}, 0)
		if err != nil  {
			return
		}
		dataStates.AddToDhtSync("snetRef", hex.EncodeToString(crypto.Keccak256Hash([]byte(data.Ref))), appRefData.MsgPack())

		// stateUpdateError := dataStates.Commit(nil, nil, nil, event.ID, err)
		// if event.IsLocal(cfg) {
			stateUpdateError := dataStates.Save(event.ID)
			if stateUpdateError != nil {
				logger.Fatalf("ApplicationStateUpdateError: %v", stateUpdateError)
				logger.Fatalf(stateUpdateError.Error())
				return
			} 
		//}
	
	
			resp = &entities.EventProcessorResponse{
				State: dataStates.CurrentStates[entities.EntityPath{Model: entities.ApplicationModel, ID: data.ID}],
				Hash: data.Hash,
			}
			// p2p.StateDhtSyncer
			// keySecP := "/ml/app/" + hex.EncodeToString(crypto.Keccak256Hash([]byte(payloadData.Ref)))
			go  OnFinishProcessingEvent(cfg, event,  &data, nil)
			// go utils.WriteBytesToFile(filepath.Join(cfg.DataDir, "log.txt"), []byte("newMessage" + "\n"))
		


	}()
	
	var localState models.ApplicationState
	
	 app, err := dsquery.GetApplicationStateById(id)
	 if err != nil && !dsquery.IsErrorNotFound(err){
		logger.Debugf("ApplicationStateQueryError: %v", err)
		return nil, err
	 }
	 if (app != nil ) {
	 	localState =  models.ApplicationState{Application: *app}
	 }

	// if err != nil {
	// 	logger.Error(err)
	// }
	
	
	var localDataState *LocalDataState
	if localState.ID != "" {
		localDataState = &LocalDataState{
			ID: localState.ID,
			Hash: localState.ID,
			Event: &localState.Event,
			Timestamp: localState.Timestamp,
		}
	}
	// localDataState := utils.IfThenElse(localTopicState != nil, &LocalDataState{
	// 	ID: localTopicState.ID,
	// 	Hash: localTopicState.ID,
	// 	Event: &localTopicState.Event,
	// 	Timestamp: localTopicState.Timestamp,
	// }, nil)
	var stateEvent *entities.Event
	if localState.ID != "" {
		stateEvent, err = dsquery.GetEventFromPath(&localState.Event)
		if err != nil && err != query.ErrorNotFound && !dsquery.IsErrorNotFound(err) {
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

	eventData := PayloadData{Application: data.ID, localDataState: localDataState, localDataStateEvent:  localDataStateEvent}
	// tx := sql.SqlDb
	// defer func () {
	// 	if tx.Error != nil {
	// 		tx.Rollback()
	// 	} else {
	// 		tx.Commit()
	// 	}
	// }()
	// txn, err := stores.EventStore.NewTransaction(context.Background(), false) // true for read-write, false for read-only
	// if err != nil {
	// 	// either app does not exist or you are not uptodate
	// }
	// defer txn.Discard(context.Background())  
	previousEventUptoDate,  _, _, eventIsMoreRecent, err := ProcessEvent(event,  eventData, false, saveApplicationEvent, nil, nil, ctx, dataStates)
	if err != nil {
		logger.Debugf("Processing Error...: %v", err)
		return nil, err
	}
	
		event.Application = id
		// err = dsquery.IncrementCounters(event.Cycle, event.Validator, event.Application, &txn)
		// if err != nil { 
		// 	return err
		// }
	
	logger.Debugf("Processing 2...: %v", previousEventUptoDate)
	if previousEventUptoDate {

		if event.Validator != entities.PublicKeyString(cfg.PublicKeyEDDHex) {
			_, err = ValidateApplicationData(&event.Payload, cfg.ChainId)
		}
		
		if err != nil {
			// update error and mark as synced
			// notify validator of error
			dataStates.AddEvent(entities.Event{ID: event.ID, Error: err.Error(), IsValid: utils.FalsePtr(), Synced:  utils.TruePtr()})
			
		} else {
			// TODO if event is older than our state, just save it and mark it as synced
			dataStates.AddEvent(entities.Event{ID: event.ID, IsValid:  utils.TruePtr(), Synced:  utils.TruePtr()})
		
			// savedEvent, err := saveApplicationEvent(entities.Event{ID: event.ID}, nil, &entities.Event{IsValid:  utils.TruePtr(), Application: event.Application, Synced:  utils.TruePtr()}, &txn, nil );
			// data.ID, _ = entities.GetId(data, id)
			
			if eventIsMoreRecent {
				// update state
					dataStates.AddCurrentState(entities.ApplicationModel, id, data)
				
				// if err != nil {
				// 	// tx.Rollback()
				// 	logger.Errorf("SaveStateError %v", err)
				// 	return err
				// } else {
				// 	_, err = saveApplicationEvent(entities.Event{ID: event.ID}, nil, &entities.Event{IsValid: utils.TruePtr(), Synced:  utils.TruePtr()}, &txn, nil )
				// }
			} else {
				dataStates.AddHistoricState(entities.ApplicationModel, data.ID, data.MsgPack())
			}
			go dsquery.UpdateAccountCounter(string(event.Payload.Account))
			// if err == nil {
			// 	if err = txn.Commit(context.Background()); err != nil {
			// 		logger.Errorf("ErorrSavingEvent: %v", err)
			// 		return err
			// 	}
			// 	go func ()  {
			// 		dsquery.UpdateAccountCounter(data.Account.ToString())
			// 		//event.Application = savedEvent.ID
			// 		dsquery.IncrementStats(event, nil)

			// 		OnFinishProcessingEvent(ctx, event, &models.ApplicationState{
			// 				Application: data,
			// 			}, &savedEvent.ID)
			// 	}()
				
			// }
			
			
			
		}

}
return resp, nil
}

// func UpdateApplicationFromPeer(appId string , cfg *configs.MainConfiguration, validator string) (*entities.Application, error) {
// 	_app := &entities.Application{}
// 	if validator == "" {
// 		validator = chain.NetworkInfo.GetRandomSyncedNode()
// 	}
// 	if len(validator) == 0 {
// 		return nil, apperror.NotFound("app not found")
// 	}
// 	subPath := entities.NewEntityPath(entities.PublicKeyString(validator), entities.ApplicationModel, appId)
// 	pp, err := p2p.GetState(cfg, *subPath, nil, _app)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(pp.Event) < 2 {
// 		return nil, apperror.NotFound("app not found")
// 	}
// 	appEvent, err := entities.UnpackEvent(pp.Event, entities.ApplicationModel)
// 	if err != nil {
// 		logger.Errorf("UnpackError: %v", err)
// 		return  nil, err
// 	}
// 	err = dsquery.CreateEvent(appEvent, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for _, snetData := range pp.States {
// 		_app, err := entities.UnpackApplication(snetData)
// 		logger.Infof("FoundApplication %v", _app)
// 		if err != nil {
// 			return  nil, apperror.NotFound("unable to retrieve app")
// 		}
// 			s, err := dsquery.CreateApplicationState(&_app, nil)
// 			logger.Infof("FoundApplication 2 %v", _app)
// 			if err != nil {
// 				return  nil, apperror.NotFound("app not saved")
// 			}
// 			_app = *s;
		
// 	}
// 	return _app, nil
// }

// 0000000000000000000000000000000000000000000000000000000000014a34 c313b453da7da4cfd1fb71a6c9c2636d47abf704e851fb8e59b8661b40deb734 00000000000001f56d69643a307835396664386639346464643166653630363664333030663734616664356533613031393730653433ddb466a5dd4a5c0835614c7a46e18943ef750a9d00000000000000000000019519cdb90e
// 0000000000000000000000000000000000000000000000000000000000014a34 f999615aca7732e509cbc8c28ef728273a207c00999a2bef3cd91fdd974d04ee 00000000000001f56d69643a307835396664386639346464643166653630363664333030663734616664356533613031393730653433ddb466a5dd4a5c0835614c7a46e18943ef750a9d00000000000000000000019519cdb90e