package client

import (
	// "errors"
	"context"
	"encoding/hex"
	"fmt"
	"slices"
	"strings"
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
	"github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"gorm.io/gorm"
)

func CreateEvent[S *models.EventInterface](payload entities.ClientPayload, ctx *context.Context) (model S, err error) {
	cfg, _ := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)
	if !strings.EqualFold(utils.AddressToHex(payload.Validator), utils.AddressToHex(cfg.OwnerAddress.String())) {
		return nil, apperror.Forbidden(fmt.Sprintf("Validator (%s) not authorized to procces this request", cfg.OwnerAddress.String()))
	}

	
	var authState *models.AuthorizationState
	var agent *entities.DeviceString
	excludedEvents := []constants.EventType{constants.CreateSubnetEvent, constants.UpdateSubnetEvent, constants.DeleteSubnetEvent, constants.AuthorizationEvent}
	if !slices.Contains(excludedEvents, constants.EventType(payload.EventType)) {
		authState, agent, err = ValidateClientPayload(&payload, true, cfg.ChainId)

		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		if authState == nil || *authState.Authorization.Priviledge == constants.UnauthorizedPriviledge {
			// agent not authorized
			return nil, apperror.Unauthorized("agent unauthorized")
		}

		if *authState.Duration != 0 && uint64(time.Now().UnixMilli()) >
			(uint64(*authState.Timestamp)+uint64(*authState.Duration)) {
			return nil, apperror.Unauthorized("Agent authorization expired")
		}
		payload.Agent = *agent
	}

	var assocPrevEvent *entities.EventPath
	var assocAuthEvent *entities.EventPath
	var eventPayloadType constants.EventPayloadType
	var subnetState = models.SubnetState{}

	if payload.Subnet != "" {
		query.GetOneState(entities.Subnet{ID: payload.Subnet}, &subnetState)
	}
	//Perfom checks base on event types
	logger.Debugf("authState****** 2: %v ", authState)
	switch payload.EventType {
	case uint16(constants.AuthorizationEvent):
		// authData := entities.Authorization{}
		// d, _ := json.Marshal(payload.Data)
		// e := json.Unmarshal(d, &authData)
		// if e != nil {
		// 	logger.Errorf("UnmarshalError %v", e)
		// }
		// payload.Data = authData
		assocPrevEvent, assocAuthEvent, err = ValidateAuthPayloadData(cfg, payload)
		if err != nil {
			return nil, err
		}
		eventPayloadType = constants.AuthorizationPayloadType
		
	case uint16(constants.CreateTopicEvent), uint16(constants.UpdateNameEvent), uint16(constants.UpdateTopicEvent), uint16(constants.LeaveEvent):
		eventPayloadType = constants.TopicPayloadType
		// if authState.Authorization.Priviledge < constants.AdminPriviledge {
		// 	return nil, apperror.Forbidden("Agent not authorized to perform this action")
		// }
		assocPrevEvent, assocAuthEvent, err = ValidateTopicPayload(payload, authState)
		if err != nil {
			return nil, err
		}
		if *authState.Authorization.Priviledge < constants.StandardPriviledge {
			return nil, apperror.Forbidden("Agent not authorized to perform this action")
		}
		if assocPrevEvent == nil {
			assocPrevEvent = &subnetState.Event
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
	case uint16(constants.CreateSubnetEvent), uint16(constants.UpdateSubnetEvent):
		eventPayloadType = constants.SubnetPayloadType
		// if authState.Authorization.Priviledge < constants.AdminPriviledge {
		// 	return nil, apperror.Forbidden("Agent not authorized to perform this action")
		// }

		assocPrevEvent, assocAuthEvent, err = ValidateSubnetPayload(payload, authState, ctx)
		if err != nil {
			return nil, err
		}
		
		
	case uint16(constants.CreateWalletEvent), uint16(constants.UpdateWalletEvent):
		eventPayloadType = constants.WalletPayloadType
		// if authState.Authorization.Priviledge < constants.AdminPriviledge {
		// 	return nil, apperror.Forbidden("Agent not authorized to perform this action")
		// }
		assocPrevEvent, assocAuthEvent, err = ValidateWalletPayload(payload, authState)
		if err != nil {
			return nil, err
		}
	case uint16(constants.SubscribeTopicEvent), uint16(constants.ApprovedEvent), uint16(constants.BanMemberEvent), uint16(constants.UnbanMemberEvent):
		if *authState.Authorization.Priviledge < constants.StandardPriviledge {
			return nil, apperror.Forbidden("Agent not authorized to perform this action")
		}
		eventPayloadType = constants.SubscriptionPayloadType

		assocPrevEvent, assocAuthEvent, err = ValidateSubscriptionPayload(payload, authState)
		if err != nil {
			return nil, err
		}
	case uint16(constants.SendMessageEvent):
		logger.Debugf("authState 2: %d ", *authState.Authorization.Priviledge)
		// 1. Agent message
		// if *authState.Authorization.Priviledge < constants.StandardPriviledge {
		// 	return nil, apperror.Forbidden("Agent not authorized to perform this action")
		// }
		eventPayloadType = constants.MessagePayloadType
		assocPrevEvent, assocAuthEvent, err = ValidateMessagePayload(payload, authState)
		if err != nil {
			logger.Error("ERRRRRRR:::", err)
			return nil, err
		}
	default:
	}
	logger.Debugf("UPDATINGSUBNE1: %v", err)
	payloadHash, err := payload.GetHash()
	
	if err != nil {
		logger.Errorf("UPDATINGSUBNET2 %v", err)
	}
	logger.Debugf("UPDATINGSUBNE_HASH: %v",payloadHash)
	chainInfo, err := chain.DefaultProvider(cfg).GetChainInfo()
	// bNum, err := chain.DefaultProvider(cfg).GetCurrentBlockNumber()
	// cycle, err := chain.DefaultProvider(cfg).GetCycle(bNum)
	logger.Debugf("UPDATINGSUBNET2: %v", chainInfo.ChainId)
	if err != nil {
		return nil, err
	}
	
	subnet := payload.Subnet
	if payload.EventType == uint16(constants.CreateSubnetEvent) {
		subnet, err = entities.GetId(payload.Data.(entities.Payload))
		if err != nil {
			logger.Debugf("Subnet error: %v", err)
			return nil, err
		}
		
	}
	if payload.EventType == uint16(constants.UpdateAvatarEvent) {
		subnet = payload.Data.(entities.Subnet).ID
	}
	if payload.EventType == uint16(constants.AuthorizationEvent) {
		subnet = payload.Data.(entities.Authorization).Subnet
	}
	logger.Info("UPDATINGSUBNET2")
	event := entities.Event{
		Payload:           payload,
		Timestamp:         uint64(time.Now().UnixMilli()),
		EventType:         uint16(payload.EventType),
		Associations:      []string{},
		PreviousEvent: *utils.IfThenElse(assocPrevEvent == nil, entities.EventPathFromString(""), assocPrevEvent),
		AuthEvent:     *utils.IfThenElse(assocAuthEvent == nil, entities.EventPathFromString(""), assocAuthEvent),
		Synced:            false,
		PayloadHash:       hex.EncodeToString(payloadHash),
		Broadcasted:       false,
		BlockNumber:       chainInfo.CurrentBlock.Uint64(),
		Cycle: 				chainInfo.CurrentCycle.Uint64(),
		Epoch: 				chainInfo.CurrentEpoch.Uint64(),		
		Validator:         entities.PublicKeyString(cfg.PublicKeyEDDHex),
		Subnet: subnet,
	}
	logger.Info("UPDATINGSUBNET3")
	logger.Debugf("NewEvent: %v", event)
	b, err := event.EncodeBytes()
	if err != nil {
		return nil, apperror.Internal(err.Error())
	}
	logger.Debugf("Validator 2: %s", event.Validator)
	logger.Debugf("eventPayloadType 2: %s", eventPayloadType)

	event.Hash = hex.EncodeToString(crypto.Sha256(b))
	_, event.Signature = crypto.SignEDD(b, cfg.PrivateKeyEDD)

	switch eventPayloadType {
	
	case constants.SubnetPayloadType:
		eModel, created, err := query.SaveRecord(
			models.SubnetEvent{
				Event: entities.Event{Hash: event.Hash},
			},
			&models.SubnetEvent{
				Event: event,
			}, nil, nil)

		if err != nil {
			return nil, err
		}

		// channelpool.SubnetEventPublishC <- &(eModel.Event)

		if created {
			channelpool.SubnetEventPublishC <- &(eModel.Event)
		}
		var returnModel = models.EventInterface(*eModel)
		model = &returnModel
	case constants.AuthorizationPayloadType:
		eModel, created, err := query.SaveAuthorizationEvent(&event, false, nil)
		if err != nil {
			return nil, err
		}
		// channelpool.AuthorizationEventPublishC <- &(eModel.Event)
		if created {
			channelpool.AuthorizationEventPublishC <- &(eModel.Event)
		}
		var returnModel = models.EventInterface(*eModel)
		model = &returnModel
	case constants.TopicPayloadType:
		eModel, created, err := query.SaveRecord(
			models.TopicEvent{
				Event: entities.Event{Hash: event.Hash},
			},
			&models.TopicEvent{
				Event: event,
			}, nil, nil)

		if err != nil {
			return nil, err
		}

		// channelpool.TopicEventPublishC <- &(eModel.Event)

		if created {
			channelpool.TopicEventPublishC <- &(eModel.Event)
		}
		var returnModel = models.EventInterface(*eModel)
		model = &returnModel
	case constants.WalletPayloadType:
		eModel, created, err := query.SaveRecord(
			models.WalletEvent{
				Event: entities.Event{Hash: event.Hash},
			},
			&models.WalletEvent{
				Event: event,
			}, nil, nil)

		if err != nil {
			return nil, err
		}

		// channelpool.WalletEventPublishC <- &(eModel.Event)

		if created {
			channelpool.WalletEventPublishC <- &(eModel.Event)
		}
		var returnModel = models.EventInterface(*eModel)
		model = &returnModel

	case constants.SubscriptionPayloadType:
		eModel, created, err := query.SaveRecord(
			models.SubscriptionEvent{
				Event: entities.Event{Hash: event.Hash},
			},
			&models.SubscriptionEvent{
				Event: event,
			}, nil, nil)

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
			&models.MessageEvent{
				Event: event,
			}, nil, nil)

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
	// logger.Debugf("Pushed event to channel %v", *model)
	//query.IncrementBlockStat(event.BlockNumber, (*constants.EventType)(&event.EventType) )
	//_, _, blockStatErr := query.IncrementBlockStat(event.BlockNumber, (*constants.EventType)(&event.EventType))

	// if blockStatErr != nil {
	// 	return nil, blockStatErr
	// }

	return model, nil

}

func GetEventTypeFromModel(eventType entities.EntityModel) constants.EventType {
	// cfg, _ := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)

	// check if client payload is valid
	// if err := payload.Validate(entities.PublicKeyString(cfg.NetworkPublicKey)); err != nil {
	// 	return  err
	// }

	//Perfom checks base on event types
	switch eventType {
	case entities.AuthModel:

		return constants.AuthorizationEvent

	case entities.TopicModel:
		return constants.CreateTopicEvent

	case entities.SubscriptionModel:
		return constants.SubscribeTopicEvent

	case entities.MessageModel:
		return constants.SendMessageEvent

	case entities.SubnetModel:
		return constants.CreateSubnetEvent

	case entities.WalletModel:
		return constants.CreateWalletEvent
	}

	return 0

}



func GetEvent(eventId string, eventType int) (model interface{}, err error) {

	if eventType < 1000 {
		event, err1 := GetAuthorizationEventById(eventId)

		if err1 != nil {
			logger.Error(err)
			return nil, err1
		}
		return event, nil
	}
	if eventType < 1100 {
		event, err1 := GetTopicEventById(eventId)

		if err1 != nil {
			logger.Error(err)
			return nil, err1
		}
		return event, nil
	}
	if eventType < 1200 {
		event, err1 := GetSubEventById(eventId)

		if err1 != nil {
			logger.Error(err)
			return nil, err1
		}
		return event, nil
	}
	if eventType < 1300 {
		event, err1 := GetMessageEventById(eventId)

		if err1 != nil {
			logger.Error(err)
			return nil, err1
		}
		return event, nil
	}
	if eventType < 1400 {
		event, err1 := GetSubnetEventById(eventId)

		if err1 != nil {
			logger.Error(err)
			return nil, err1
		}
		return event, nil
	}
	if eventType < 1400 {
		event, err1 := GetWalletEventById(eventId)

		if err1 != nil {
			logger.Error(err)
			return nil, err1
		}
		return event, nil
	}
	return nil, fmt.Errorf("invalid event type")
}

func GetEventByHash(eventHash string, eventType int) (model interface{}, err error) {

	switch uint16(eventType) {
	case uint16(constants.CreateTopicEvent), uint16(constants.UpdateNameEvent), uint16(constants.UpdateTopicEvent), uint16(constants.LeaveEvent):
		event, err1 := GetTopicEventByHash(eventHash)

		if err1 != nil {
			logger.Error(err)
			return nil, err1
		}
		return event, nil

	case uint16(constants.CreateSubnetEvent):
		event, err1 := GetSubnetEventByHash(eventHash)

		if err1 != nil {
			logger.Error(err)
			return nil, err1
		}
		return event, nil

	case uint16(constants.SubscribeTopicEvent), uint16(constants.ApprovedEvent), uint16(constants.BanMemberEvent), uint16(constants.UnbanMemberEvent):
		event, err1 := GetSubEventByHash(eventHash)

		if err1 != nil {
			logger.Error(err)
			return nil, err1
		}
		return event, nil
	case uint16(constants.SendMessageEvent), uint16(constants.DeleteMessageEvent):
		event, err1 := GetMessageEventByHash(eventHash)

		if err1 != nil {
			logger.Error(err)
			return nil, err1
		}
		return event, nil

	case uint16(constants.AuthorizationEvent), uint16(constants.UnauthorizationEvent):
		event, err1 := GetAuthorizationEventByHash(eventHash)

		if err1 != nil {
			logger.Error(err)
			return nil, err1
		}
		return event, nil
	default:
	}

	return model, nil

}

func GetTopicEventById(id string) (*models.TopicEvent, error) {
	nEvent := models.TopicEvent{}

	// err := query.GetOne(&models.TopicEvent{
	// 	Event: entities.Event{ID: id},
	// }, &nEvent)
	err := sql.SqlDb.Where(&models.TopicEvent{
			Event: entities.Event{ID: id},
		}).Take(&nEvent).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &nEvent, nil

}

func GetTopicEventByHash(hash string) (*models.TopicEvent, error) {
	nEvent := models.TopicEvent{}

	err := query.GetOne(models.TopicEvent{
		Event: entities.Event{Hash: hash},
	}, &nEvent)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &nEvent, nil

}

func GetSubnetEventById(id string) (*models.SubnetEvent, error) {
	nEvent := models.SubnetEvent{}

	err := query.GetOne(models.SubnetEvent{
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

func GetSubnetEventByHash(hash string) (*models.SubnetEvent, error) {
	nEvent := models.SubnetEvent{}

	err := query.GetOne(models.SubnetEvent{
		Event: entities.Event{Hash: hash},
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

func GetSubEventByHash(hash string) (*models.SubscriptionEvent, error) {
	nEvent := models.SubscriptionEvent{}

	err := query.GetOne(models.SubscriptionEvent{
		Event: entities.Event{Hash: hash},
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

func GetMessageEventByHash(hash string) (*models.MessageEvent, error) {
	nEvent := models.MessageEvent{}

	err := query.GetOne(models.MessageEvent{
		Event: entities.Event{Hash: hash},
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

func GetAuthorizationEventByHash(hash string) (*models.AuthorizationEvent, error) {
	nEvent := models.AuthorizationEvent{}

	err := query.GetOne(models.AuthorizationEvent{
		Event: entities.Event{Hash: hash},
	}, &nEvent)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &nEvent, nil
}



func GetEventModelFromEventType(eventType constants.EventType)  any {
	if eventType < 1000 {
		return &models.AuthorizationEvent{}
	}
	if eventType < 1100 {
		return &models.TopicEvent{}
	}
	if eventType < 1200 {
		return &models.SubscriptionEvent{}
	}
	if eventType < 1300 {
		return &models.MessageEvent{}
	}
	if eventType < 1400 {
		return &models.SubnetEvent{}
	}
	if eventType < 1400 {
		return &models.WalletEvent{}
	}
	return nil
}