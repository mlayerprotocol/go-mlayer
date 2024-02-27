package client

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"

	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/chain"
	"github.com/mlayerprotocol/go-mlayer/internal/channelpool"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/internal/service"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
)

/*
	Validate and Process the authorization request
*/
/*
	Validate and Process the authorization request
*/
func AuthorizeAgent(
	payload entities.ClientPayload, ctx *context.Context,
) (*models.AuthorizationEvent, error) {

	cfg, _ := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)

	logger.Info(cfg.NetworkPublicKey)
	logger.Info(payload.Validator)

	if string(payload.Validator) != cfg.NetworkPublicKey {
		return nil, apperror.Forbidden("Validator not authorized to procces this request")
	}

	authData := entities.Authorization{}
	d, _ := json.Marshal(payload.Data)
	e := json.Unmarshal(d, &authData)
	if e != nil {
		logger.Errorf("UnmarshalError %v", e)
	}
	payload.Data = authData
	data := payload
	var assocPrevEvent string
	var assocAuthEvent string
	if payload.EventType == uint16(constants.AuthorizationEvent) {
		// dont worry validating the AuthHash for Authorization requests

		if authData.Duration != 0 && uint64(time.Now().UnixMilli()) >
			(uint64(data.Timestamp)+uint64(authData.Duration)) {
			return nil, errors.New("Authorization duration exceeded")
		}

		if uint64(authData.Timestamp) > uint64(time.Now().UnixMilli())+15000 {
			return nil, errors.New("Authorization timestamp exceeded")
		}

		currentState, grantorAuthState, err := service.ValidateAuthData(&authData)
		if err != nil {
			return nil, err
		}

		// generate associations
		if currentState != nil {
			assocPrevEvent = currentState.EventHash
			// assocPrevEvent = entities.EventPath{
			// 	Relationship: entities.PreviousEventAssoc,
			// 	Hash: currentState.EventHash,
			// 	Model: entities.AuthorizationEventModel,
			// }.ToString()
		}
		if grantorAuthState != nil {
			assocAuthEvent = grantorAuthState.EventHash
			// assocAuthEvent =  entities.EventPath{
			// 	Relationship: entities.AuthorizationEventAssoc,
			// 	Hash: grantorAuthState.EventHash,
			// 	Model: entities.AuthorizationEventModel,
			// }
		}

	}
	payloadHash, _ := payload.GetHash()
	// hash the payload  Nonce
	// payload.Nonce = string(crypto.Keccak256Hash(encoder.EncodeBytes(encoder.IntEncoderDataType(payload.Nonce))
	// create event struct
	event := entities.Event{
		Payload:           payload,
		Timestamp:         uint64(time.Now().UnixMilli()),
		EventType:         payload.EventType,
		Associations:      []string{},
		PreviousEventHash: assocPrevEvent,
		AuthEventHash:     assocAuthEvent,
		Synced:            false,
		PayloadHash:       hex.EncodeToString(payloadHash),
		Broadcasted:       false,
		BlockNumber:       chain.MLChainApi.GetCurrentBlockNumber(),
		Validator:         entities.PublicKeyString(cfg.NetworkPublicKey),
	}

	b, err := event.EncodeBytes()

	event.Hash = hex.EncodeToString(crypto.Sha256(b))
	_, event.Signature = crypto.SignEDD(b, cfg.NetworkPrivateKey)

	eModel, created, err := query.SaveAuthorizationEvent(&event, false, nil)
	if err != nil {
		return nil, err
	}

	channelpool.AuthorizationEventPublishC <- &(eModel.Event)
	if created {
		channelpool.AuthorizationEventPublishC <- &(eModel.Event)
	}
	return eModel, nil
}

// func ListenForNewAuthEventFromPubSub (mainCtx *context.Context) {
// 	ctx, cancel := context.WithCancel(*mainCtx)
// 	defer cancel()
// 	time.AfterFunc(5*time.Second, func() {
// 		logger.Info("Sending subscription to channel")
// 		//subscriptionPubSub.Publish(entities.NewPubSubMessage((&entities.Subscription{Signature: "channel", Subscriber: "sds"}).MsgPack()))
// 	})
// 	incomingAuthorizationC, ok := (*mainCtx).Value(constants.IncomingAuthorizationEventChId).(*chan *entities.Event)
// 	if !ok {
// 		logger.Errorf("incomingAuthorizationC closed")
// 		return
// 	}
// 	for {
// 		event, ok :=  <-*incomingAuthorizationC
// 		if !ok {
// 			logger.Fatal("incomingAuthorizationC closed")
// 			return
// 		}
// 		go service.HandleNewPubSubAuthEvent(event, ctx)
// 	}
// }
