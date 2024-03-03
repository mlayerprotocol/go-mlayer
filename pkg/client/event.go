package client

import (
	// "errors"
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

func validateTopicPayload(payload entities.ClientPayload) (assocPrevEvent string, assocAuthEvent string, eventType string, err error) {

	eventType = "topic"
	payloadData := entities.Topic{}
	d, _ := json.Marshal(payload.Data)
	e := json.Unmarshal(d, &payloadData)
	if e != nil {
		logger.Errorf("UnmarshalError %v", e)
	}

	payload.Data = payloadData
	if payload.EventType == uint16(constants.CreateTopicEvent) {
		// dont worry validating the AuthHash for Authorization requests
		if uint64(payloadData.Timestamp) > uint64(time.Now().UnixMilli())+15000 {
			return "", "", "", errors.New("Authorization timestamp exceeded")
		}

	}

	currentState, grantorAuthState, err := service.ValidateTopicData(&payloadData)
	if err != nil {
		return "", "", "", err
	}

	// generate associations
	if currentState != nil {
		assocPrevEvent = currentState.EventHash

	}
	if grantorAuthState != nil {
		assocAuthEvent = grantorAuthState.EventHash
	}
	return assocPrevEvent, assocAuthEvent, "topic", nil
}

func validateSubscriptionPayload(payload entities.ClientPayload) (assocPrevEvent string, assocAuthEvent string, eventType string, err error) {

	eventType = "subscription"
	payloadData := entities.Subscription{}
	d, _ := json.Marshal(payload.Data)
	e := json.Unmarshal(d, &payloadData)
	if e != nil {
		logger.Errorf("UnmarshalError %v", e)
	}
	payload.Data = payloadData

	topicData, err := GetTopic(payloadData.Topic)
	if err != nil {
		return "", "", "", err
	}

	if topicData == nil {
		return "", "", "", apperror.BadRequest("Invalid topic")
	}

	// pool = channelpool.SubscriptionEventPublishC

	currentState, grantorAuthState, err := service.ValidateSubscriptionData(&payloadData, &payload)
	if err != nil {
		return "", "", "", err
	}

	if currentState == nil {
		return "", "", "", apperror.BadRequest("Account not subscribed")
	}
	if grantorAuthState == nil {
		return "", "", "", apperror.Unauthorized("Agent does not have sufficient privileges to update this event")
	}

	// dont worry validating the AuthHash for Authorization requests
	// if uint64(payloadData.Timestamp) > uint64(time.Now().UnixMilli())+15000 {
	// 	return  errors.New("Authorization timestamp exceeded")
	// }

	// generate associations
	if currentState != nil {
		assocPrevEvent = currentState.EventHash

	}

	if grantorAuthState != nil {
		assocAuthEvent = grantorAuthState.EventHash
		// assocAuthEvent =  entities.EventPath{
		// 	Relationship: entities.AuthorizationEventAssoc,
		// 	Hash: grantorAuthState.EventHash,
		// 	Model: entities.AuthorizationEventModel,
		// }
	}
	return assocPrevEvent, assocAuthEvent, "subscription", nil
}

func CreatTopSubEvent[S *models.EventInterface](payload entities.ClientPayload, ctx *context.Context) (model S, err error) {
	cfg, _ := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)

	// check if client payload is valid
	// if err := payload.Validate(entities.PublicKeyString(cfg.NetworkPublicKey)); err != nil {
	// 	return  err
	// }

	authState, err := ValidateClientPayload(&payload)
	if authState == nil && payload.EventType != uint16(constants.CreateTopicEvent) {
		// agent not authorized
		return nil, apperror.Unauthorized("Agent not authorized to perform this update")
	}

	logger.Info("Before Assoc Event")
	var assocPrevEvent string
	var assocAuthEvent string
	var eventType string

	//Perfom checks base on event types
	switch payload.EventType {
	case uint16(constants.CreateTopicEvent):
		assocPrevEvent, assocAuthEvent, eventType, err = validateTopicPayload(payload)

		if err != nil {
			return nil, err
		}
	case uint16(constants.UpdateNameEvent):
		assocPrevEvent, assocAuthEvent, eventType, err = validateTopicPayload(payload)
		if err != nil {
			return nil, err
		}
	case uint16(constants.JoinTopicEvent):
		eventType = "subscription"
		var currentState models.SubscriptionState
		// generate associations
		if currentState.ID != "" {
			assocPrevEvent = currentState.EventHash
		}
	case uint16(constants.LeaveEvent):
		assocPrevEvent, assocAuthEvent, eventType, err = validateTopicPayload(payload)
		if err != nil {
			return nil, err
		}
	case uint16(constants.ApprovedEvent):
		assocPrevEvent, assocAuthEvent, eventType, err = validateSubscriptionPayload(payload)
		if err != nil {
			return nil, err
		}
	case uint16(constants.BanMemberEvent):
		assocPrevEvent, assocAuthEvent, eventType, err = validateSubscriptionPayload(payload)
		if err != nil {
			return nil, err
		}
	case uint16(constants.UnbanMemberEvent):
		assocPrevEvent, assocAuthEvent, eventType, err = validateSubscriptionPayload(payload)
		if err != nil {
			return nil, err
		}
	default:

	}
	payloadHash, _ := payload.GetHash()

	event := entities.Event{
		Payload:           payload,
		Timestamp:         uint64(time.Now().UnixMilli()),
		EventType:         uint16(payload.EventType),
		Associations:      []string{},
		PreviousEventHash: assocPrevEvent,
		AuthEventHash:     assocAuthEvent,
		Synced:            false,
		PayloadHash:       hex.EncodeToString(payloadHash),
		Broadcasted:       false,
		BlockNumber:       chain.MLChainApi.GetCurrentBlockNumber(),
		Validator:         entities.PublicKeyString(cfg.NetworkPublicKey),
	}

	logger.Infof("Validator: %s", event.Validator)
	b, err := event.EncodeBytes()

	event.Hash = hex.EncodeToString(crypto.Sha256(b))
	_, event.Signature = crypto.SignEDD(b, cfg.NetworkPrivateKey)

	if eventType == "topic" {
		eModel, created, err := query.SaveRecord(
			models.TopicEvent{
				Event: entities.Event{Hash: event.Hash},
			},
			models.TopicEvent{
				Event: event,
			}, false, nil)

		if err != nil {
			return nil, err
		}

		// channelpool.TopicEventPublishC <- &(eModel.Event)

		if created {
			channelpool.TopicEventPublishC <- &(eModel.Event)
		}
		var returnModel = models.EventInterface(*eModel)
		model = &returnModel
	}

	if eventType == "subscription" {
		eModel, created, err := query.SaveRecord(
			models.SubscriptionEvent{
				Event: entities.Event{Hash: event.Hash},
			},
			models.SubscriptionEvent{
				Event: event,
			}, false, nil)

		if err != nil {
			return nil, err
		}

		// channelpool.TopicEventPublishC <- &(eModel.Event)

		if created {
			channelpool.SubscriptionEventPublishC <- &(eModel.Event)
		}
		var returnModel = models.EventInterface(*eModel)
		model = &returnModel
	}

	return model, nil

}
