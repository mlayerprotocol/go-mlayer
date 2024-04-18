package client

import (
	// "errors"
	"context"
	"encoding/hex"
	"time"

	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/chain"
	"github.com/mlayerprotocol/go-mlayer/internal/channelpool"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"gorm.io/gorm"
)

func CreateEvent[S *models.EventInterface](payload entities.ClientPayload, ctx *context.Context) (model S, err error) {
	cfg, _ := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)

	// check if client payload is valid
	// if err := payload.Validate(entities.PublicKeyString(cfg.NetworkPublicKey)); err != nil {
	// 	return  err
	// }

	payload.Agent, err = payload.GetSigner()
	if err != nil {
		return nil, apperror.Internal(err.Error())
	}
	authState, err := ValidateClientPayload(&payload)
	if err != nil {
		return nil, apperror.Internal(err.Error())
	}
	if authState == nil && payload.EventType != uint16(constants.AuthorizationEvent) {
		// agent not authorized
		return nil, apperror.Unauthorized("Agent not authorized to perform this action")
	}

	var assocPrevEvent *entities.EventPath
	var assocAuthEvent *entities.EventPath
	var eventPayloadType constants.EventPayloadType

	//Perfom checks base on event types
	switch payload.EventType {
	case uint16(constants.CreateTopicEvent), uint16(constants.UpdateNameEvent), uint16(constants.UpdateTopicEvent), uint16(constants.LeaveEvent):
		eventPayloadType = constants.TopicPayloadType
		// if authState.Authorization.Priviledge < constants.AdminPriviledge {
		// 	return nil, apperror.Forbidden("Agent not authorized to perform this action")
		// }
		assocPrevEvent, assocAuthEvent, err = ValidateTopicPayload(payload, authState)
		if err != nil {
			return nil, err
		}
	// case uint16(constants.SubscribeTopicEvent):
	// 	if authState.Authorization.Priviledge < constants.AdminPriviledge {
	// 		return nil, apperror.Forbidden("Agent not authorized to perform this action")
	// 	}
	// 	eventPayloadType = constants.SubscriptionPayloadType
	// 	assocPrevEvent, assocAuthEvent,  err = ValidateSubscriptionPayload(payload, authState)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	case uint16(constants.SubscribeTopicEvent), uint16(constants.ApprovedEvent), uint16(constants.BanMemberEvent), uint16(constants.UnbanMemberEvent):
		if authState.Authorization.Priviledge < constants.WritePriviledge {
			return nil, apperror.Forbidden("Agent not authorized to perform this action")
		}
		eventPayloadType = constants.SubscriptionPayloadType
		assocPrevEvent, assocAuthEvent, err = ValidateSubscriptionPayload(payload, authState)
		if err != nil {
			return nil, err
		}
	case uint16(constants.SendMessageEvent):
		if authState.Authorization.Priviledge < constants.WritePriviledge {
			return nil, apperror.Forbidden("Agent not authorized to perform this action")
		}
		eventPayloadType = constants.MessagePayloadType
		assocPrevEvent, assocAuthEvent, err = ValidateMessagePayload(payload, authState)
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
		PreviousEventHash: *utils.IfThenElse(assocPrevEvent == nil, entities.EventPathFromString(""), assocPrevEvent),
		AuthEventHash:     *utils.IfThenElse(assocAuthEvent == nil, entities.EventPathFromString(""), assocAuthEvent),
		Synced:            false,
		PayloadHash:       hex.EncodeToString(payloadHash),
		Broadcasted:       false,
		BlockNumber:       chain.MLChainApi.GetCurrentBlockNumber(),
		Validator:         entities.PublicKeyString(cfg.NetworkPublicKey),
	}

	logger.Infof("NewEvent: %v", event)
	b, err := event.EncodeBytes()
	if err != nil {
		return nil, apperror.Internal(err.Error())
	}
	logger.Infof("Validator 2: %s", event.Validator)
	logger.Infof("eventPayloadType 2: %s", eventPayloadType)

	event.Hash = hex.EncodeToString(crypto.Sha256(b))
	_, event.Signature = crypto.SignEDD(b, cfg.NetworkPrivateKey)

	switch eventPayloadType {
	case constants.TopicPayloadType:
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

	case constants.SubscriptionPayloadType:
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

	case constants.MessagePayloadType:
		eModel, created, err := query.SaveRecord(
			models.MessageEvent{
				Event: entities.Event{Hash: event.Hash},
			},
			models.MessageEvent{
				Event: event,
			}, false, nil)

		if err != nil {
			return nil, err
		}

		// channelpool.TopicEventPublishC <- &(eModel.Event)

		if created {
			channelpool.MessageEventPublishC <- &(eModel.Event)
		}
		var returnModel = models.EventInterface(*eModel)
		model = &returnModel
	}

	_, _, blockStatErr := query.IncrementBlockStat(models.BlockStat{
		Stats: entities.Stats{
			BlockNumber: event.BlockNumber,
			EventType:   event.EventType,
		},
	})

	if blockStatErr != nil {
		return nil, blockStatErr
	}

	return model, nil

}

func GetEvent(eventId string, eventType int) (model interface{}, err error) {
	// cfg, _ := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)

	// check if client payload is valid
	// if err := payload.Validate(entities.PublicKeyString(cfg.NetworkPublicKey)); err != nil {
	// 	return  err
	// }

	//Perfom checks base on event types
	switch uint16(eventType) {
	case uint16(constants.CreateTopicEvent), uint16(constants.UpdateNameEvent), uint16(constants.UpdateTopicEvent), uint16(constants.LeaveEvent):
		topic, err1 := GetTopicEventById(eventId)

		if err1 != nil {
			logger.Error(err)
			return nil, err1
		}
		return topic, nil

	case uint16(constants.SubscribeTopicEvent), uint16(constants.ApprovedEvent), uint16(constants.BanMemberEvent), uint16(constants.UnbanMemberEvent):
		topic, err1 := GetSubEventById(eventId)

		if err1 != nil {
			logger.Error(err)
			return nil, err1
		}
		return topic, nil
	case uint16(constants.SendMessageEvent), uint16(constants.DeleteMessageEvent):
		topic, err1 := GetMessageEventById(eventId)

		if err1 != nil {
			logger.Error(err)
			return nil, err1
		}
		return topic, nil

	case uint16(constants.AuthorizationEvent), uint16(constants.UnauthorizationEvent):
		topic, err1 := GetAuthorizationEventById(eventId)

		if err1 != nil {
			logger.Error(err)
			return nil, err1
		}
		return topic, nil
	default:
	}

	return model, nil

}

func GetTopicEventById(id string) (*models.TopicEvent, error) {
	nEvent := models.TopicEvent{}

	err := query.GetOne(models.TopicEvent{
		Event: entities.Event{ID: id},
	}, &nEvent)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &nEvent, nil

}

func GetSubEventById(id string) (*models.SubscriptionEvent, error) {
	nEvent := models.SubscriptionEvent{}

	err := query.GetOne(models.SubscriptionEvent{
		Event: entities.Event{ID: id},
	}, &nEvent)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &nEvent, nil
}

func GetMessageEventById(id string) (*models.MessageEvent, error) {
	nEvent := models.MessageEvent{}

	err := query.GetOne(models.MessageEvent{
		Event: entities.Event{ID: id},
	}, &nEvent)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &nEvent, nil
}

func GetAuthorizationEventById(id string) (*models.AuthorizationEvent, error) {
	nEvent := models.AuthorizationEvent{}

	err := query.GetOne(models.AuthorizationEvent{
		Event: entities.Event{ID: id},
	}, &nEvent)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &nEvent, nil
}
