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
		case uint16(constants.CreateTopicEvent), uint16(constants.UpdateNameEvent), uint16(constants.LeaveEvent):
			eventPayloadType = constants.TopicPayloadType
			if authState.Authorization.Priviledge < constants.AdminPriviledge {
				return nil, apperror.Forbidden("Agent not authorized to perform this action")
			}
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
			if authState.Authorization.Priviledge < constants.AdminPriviledge {
				return nil, apperror.Forbidden("Agent not authorized to perform this action")
			}
			eventPayloadType = constants.SubscriptionPayloadType
			assocPrevEvent, assocAuthEvent,  err = ValidateSubscriptionPayload(payload, authState)
			if err != nil {
				return nil, err
			}
		case uint16(constants.SendMessageEvent):
			if authState.Authorization.Priviledge < constants.WritePriviledge {
				return nil, apperror.Forbidden("Agent not authorized to perform this action")
			}
			eventPayloadType = constants.MessagePayloadType
			assocPrevEvent, assocAuthEvent,  err = ValidateMessagePayload(payload, authState)
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
		PreviousEventHash: utils.IfThenElse(assocPrevEvent == nil, *entities.EventPathFromString(""), *assocPrevEvent),
		AuthEventHash:     utils.IfThenElse(assocAuthEvent == nil, *entities.EventPathFromString(""), *assocAuthEvent),
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

	event.Hash = hex.EncodeToString(crypto.Sha256(b))
	_, event.Signature = crypto.SignEDD(b, cfg.NetworkPrivateKey)

	switch(eventPayloadType) {
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

		case  constants.SubscriptionPayloadType:
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
