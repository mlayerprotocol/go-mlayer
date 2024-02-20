package client

import (
	// "errors"
	"context"
	"encoding/hex"
	"encoding/json"
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
	"gorm.io/gorm"
)

func GetSubscriptions() (*[]models.SubscriptionState, error) {
	var subscriptionStates []models.SubscriptionState

	err := query.GetMany(models.SubscriptionState{}, &subscriptionStates)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &subscriptionStates, nil
}

func GetSubscription(id string) (*models.SubscriptionState, error) {
	subscriptionState := models.SubscriptionState{}

	err := query.GetOne(models.SubscriptionState{
		Subscription: entities.Subscription{ID: id},
	}, &subscriptionState)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &subscriptionState, nil

}

func CreateSubscriptionEvent(payload entities.ClientPayload, ctx *context.Context) (*models.SubscriptionEvent, error) {
	cfg, _ := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)

	// check if client payload is valid

	// if err := payload.Validate(entities.PublicKeyString(cfg.NetworkPublicKey)); err != nil {
	// 	return nil, err
	// }

	authState, err := ValidateClientPayload(&payload)
	if authState == nil {
		// agent not authorized
		return nil, apperror.Unauthorized("Agent not authorized to update this topic")
	}

	payloadData := entities.Subscription{}
	d, _ := json.Marshal(payload.Data)
	e := json.Unmarshal(d, &payloadData)
	if e != nil {
		logger.Errorf("UnmarshalError %v", e)
	}
	payload.Data = payloadData

	topicData, err := GetTopic(payloadData.Topic)
	if err != nil {
		return nil, err
	}

	if topicData == nil {
		return nil, apperror.BadRequest("Invalid topic")
	}

	logger.Info("Before Assoc Event")
	var assocPrevEvent string
	var assocAuthEvent string
	// var pool chan *entities.Event

	if payload.EventType == uint16(constants.JoinTopicEvent) {
		var currentState models.SubscriptionState

		// generate associations
		if currentState.ID != "" {
			assocPrevEvent = currentState.EventHash

		}

		// pool = channelpool.SubscriptionEventPublishC

	}

	if payload.EventType == uint16(constants.LeaveEvent) {
		// dont worry validating the AuthHash for Authorization requests
		// if uint64(payloadData.Timestamp) > uint64(time.Now().UnixMilli())+15000 {
		// 	return nil, errors.New("Authorization timestamp exceeded")
		// }
		currentState, grantorAuthState, err := service.ValidateSubscriptionData(&payloadData)
		if err != nil {
			return nil, err
		}

		if currentState == nil {
			return nil, apperror.BadRequest("Account not subscribed")
		}

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

		// pool = channelpool.UnSubscribeEventPublishC

	}

	if payload.EventType == uint16(constants.ApprovedEvent) {
		if authState.Priviledge != 2 {
			return nil, apperror.Unauthorized("Agent does not have sufficient privileges to update this topic")
		}
		// dont worry validating the AuthHash for Authorization requests
		// if uint64(payloadData.Timestamp) > uint64(time.Now().UnixMilli())+15000 {
		// 	return nil, errors.New("Authorization timestamp exceeded")
		// }
		currentState, grantorAuthState, err := service.ValidateApproveSubscriptionData(&payloadData)
		if err != nil {
			return nil, err
		}

		if currentState == nil {
			return nil, apperror.BadRequest("Subscription not found")
		}

		if currentState.Status == string(constants.UNSUBSCRIBED) || currentState.Status == string(constants.APPROVED) {
			return nil, apperror.BadRequest("Invalid request")
		}

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
		// pool = channelpool.ApproveSubscribeEventPublishC

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
	return eModel, nil

}
