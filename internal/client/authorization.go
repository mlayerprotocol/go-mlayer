package client

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"

	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/chain"
	"github.com/mlayerprotocol/go-mlayer/internal/channelpool"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"github.com/mlayerprotocol/go-mlayer/utils/constants"
	"gorm.io/gorm"
)



func ProcessAuthorizationPayload(
	payload entities.ClientPayload, ctx *context.Context,
	) (*models.AuthorizationEvent, error) {
	
	cfg, _ :=(*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)

	// check if client payload is valid
	
	if err := payload.Validate(cfg.NetworkPrivateKey); err != nil {
		return nil, err
	}

	
	authData := entities.Authorization{} 
	d, _ := json.Marshal(payload.Data)
	e := json.Unmarshal(d, &authData)
	if e!=nil {
		logger.Errorf("UnmarshalError %v", e)
	}
	payload.Data = authData
	data := payload
	// authPayload := entities.AuthorizationPayload {
	// 	ClientPayload: entities.ClientPayload{
	// 		Data: authData,
	// 	},
	// }

	if err := authData.Validate(cfg.NetworkPrivateKey); err != nil {
		return nil, err
	}
	
	// authData := payload.ClientPayload.Data.(entities.Authorization)
	pDataBytes, err := authData.EncodeBytes()
	if err != nil {
		return nil, err
	}
	// logger.Infof("ENDOSDSDSDDS---> %v", pDataBytes)
	if payload.EventType == uint16(constants.AuthorizationEvent) {
		// dont worry validating the AuthHash for Authorization requests
		if err != nil {
			return nil, err
		}
		
		if authData.Duration != 0 && uint64(time.Now().UnixMilli()) > 
		(uint64(data.Timestamp) + uint64(authData.Duration)) {
			return nil, errors.New("Authorization duration exceeded")
		}
		valid, err := crypto.VerifySignatureEDD(authData.Grantor, &pDataBytes, authData.Signature)
		if err != nil {
			return nil, err
		}
		if !valid {
			return nil, errors.New("Invalid authorization signer")
		}
		logger.Infof("TimestampLapse %d", uint64(time.Now().UnixMilli()) - uint64(authData.Timestamp))
		if (uint64(authData.Timestamp) > uint64(time.Now().UnixMilli()) + 15000)  {
			return nil,  errors.New("Authorization timestamp exceeded")
		}

	}
	// payload.Data.Hash =  hex.EncodeToString(crypto.Sha256(pDataBytes))

	// check for existing authorization 
	authModel, err := query.GetAuthorizationState(authData.Grantor, authData.Agent)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	if authModel != nil {

	}
	
	// create event struct
	event :=  entities.Event{
			Payload: payload,
			Timestamp: uint64(time.Now().UnixMilli()),
			EventType: uint16(payload.EventType),
			Parents : []string{},
			Synced: false,
			Broadcasted : false,
			BlockNumber: chain.MLChainApi.GetCurrentBlockNumber(),
			Node: crypto.GetPublicKeyEDD( cfg.NetworkPrivateKey),
		}
	
	
	b, err := event.EncodeBytes()
	logger.Infof("BYTESSS %v", b)
	event.Hash = hex.EncodeToString(crypto.Sha256(b))
	_, event.Signature  = crypto.SignEDD(b, cfg.NetworkPrivateKey)
	
	
	eModel, created, err := query.SaveAuthorizationEvent(&event, nil)
	if err != nil {
		return nil, err
	}
	
	channelpool.BroadcastAuthorizationEventInternal_PubSubC <- &(eModel.Event)
	if created {
		// encoded, _ := encoder.MsgPackStruct(*eModel)
		// packed, _ := encoder.MsgPackStruct(*eModel)
		
		channelpool.BroadcastAuthorizationEventInternal_PubSubC <- &(eModel.Event)
	}
	return eModel, nil
}

func ProcessAuthorizationFromPubSub (mainCtx *context.Context) {
	_, cancel := context.WithCancel(*mainCtx)
	defer cancel()
	incomingAuthorizationC, ok := (*mainCtx).Value(constants.IncomingAuthorizationEventChId).(*chan *entities.Event)
	if !ok {
		logger.Errorf("incomingAuthorizationC closed")
		return
	}
	for {
		event, ok :=  <-*incomingAuthorizationC
		if !ok {
			logger.Errorf("incomingAuthorizationC closed")
			return
		}
		
		// TODO 
		
		// if err != nil {
		// 	logger.Errorf("Error unmarshaling event payload %v", err)
		// }
		// payload:
		
		// event.Event.Payload = event.Event.Payload.(entities.AuthorizationPayload)
		// event.Payload.ClientPayload.Data = event.Payload.Data 

		//authPl := entities.AuthorizationPayload{}
		// authPlByte, _ := json.Marshal( &event.Payload)
		// _	= json.Unmarshal(authPlByte, &authPl)

		// logger.Infof("Event ----===> %v", event.Payload)
		// auth := entities.Authorization{}
		// authByte, _ := json.Marshal( authPl.Data)
		// _	= json.Unmarshal(authByte, &auth)

		// clPl := event.Event.Payload.(entities.ClientPayload)
		// clPl.Data = auth
		
		// authPl.Data = auth
		// event.Event.Payload = event.Event.Payload.(entities.AuthorizationPayload)
		// event.Event.Payload = authPl
		//  
		
		
		err := ValidateEvent(*event)
		if err != nil {
			logger.Errorf("Invalid Event %v", err)
		}
		logger.Infof("Event is a valid event %s",  event.ID)
	}
}
