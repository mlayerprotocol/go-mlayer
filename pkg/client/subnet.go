package client

import (
	// "errors"

	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/service"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"gorm.io/gorm"
)

// type SubnetService struct {
// 	Ctx context.Context
// 	Cfg configs.MainConfiguration
// }

// func NewSubnetService(mainCtx *context.Context) *SubnetService {
// 	ctx := *mainCtx
// 	cfg, _ := ctx.Value(constants.ConfigKey).(*configs.MainConfiguration)
// 	return &SubnetService{
// 		Ctx: ctx,
// 		Cfg: *cfg,
// 	}
// }

// func (p *SubnetService) NewSubnetSubscription(sub *entities.Subscription) error {
// 	// subscribersc, ok := p.Ctx.Value(utils.SubscribeChId).(*chan *entities.Subscription)

// 	// validate before storing
// 	if entities.IsValidSubscription(*sub, true) {
// 		SubnetSubscriberStore, ok := p.Ctx.Value(constants.NewSubnetSubscriptionStore).(*db.Datastore)
// 		if !ok {
// 			return errors.New("Could not connect to subscription datastore")
// 		}
// 		error := SubnetSubscriberStore.Set(p.Ctx, db.Key(sub.Key()), sub.MsgPack(), false)
// 		if error != nil {
// 			return error
// 		}
// 	}
// 	return nil
// }

/*
Validate and Process the Subnet request
*/

func GetSubnetById(id string) (*models.SubnetState, error) {
	SubnetState := models.SubnetState{}

	err := query.GetOne(models.SubnetState{
		Subnet: entities.Subnet{ID: id},
	}, &SubnetState)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &SubnetState, nil

}
func GetSubnetByHash(hash string) (*models.SubnetState, error) {
	SubnetState := models.SubnetState{}

	err := query.GetOne(models.SubnetState{
		Subnet: entities.Subnet{Hash: hash},
	}, &SubnetState)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &SubnetState, nil

}

func GetSubnets(item models.SubnetState) (*[]models.SubnetState, error) {
	var SubnetStates []models.SubnetState

	err := query.GetMany(item, &SubnetStates, nil)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &SubnetStates, nil
}

func GetSubnetEvents() (*[]models.SubnetEvent, error) {
	var SubnetEvents []models.SubnetEvent

	err := query.GetMany(models.SubnetEvent{
		Event: entities.Event{
			BlockNumber: 1,
		},
	}, &SubnetEvents, nil)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &SubnetEvents, nil
}

// func ListenForNewSubnetEventFromPubSub (mainCtx *context.Context) {
// 	ctx, cancel := context.WithCancel(*mainCtx)
// 	defer cancel()

//		incomingSubnetC, ok := (*mainCtx).Value(constants.IncomingSubnetEventChId).(*chan *entities.Event)
//		if !ok {
//			logger.Errorf("incomingSubnetC closed")
//			return
//		}
//		for {
//			event, ok :=  <-*incomingSubnetC
//			if !ok {
//				logger.Fatal("incomingSubnetC closed for read")
//				return
//			}
//			go service.HandleNewPubSubSubnetEvent(event, ctx)
//		}
//	}
func ValidateSubnetPayload(payload entities.ClientPayload, authState *models.AuthorizationState, ctx *context.Context) (assocPrevEvent *entities.EventPath, assocAuthEvent *entities.EventPath, err error) {

	payloadData := entities.Subnet{}
	d, _ := json.Marshal(payload.Data)
	e := json.Unmarshal(d, &payloadData)
	if e != nil {
		logger.Errorf("UnmarshalError %v", e)
	}

	payload.Data = payloadData
	if payload.EventType == uint16(constants.CreateSubnetEvent) {
		// dont worry validating the AuthHash for Authorization requests
		if uint64(payloadData.Timestamp) > uint64(time.Now().UnixMilli())+15000 {
			return nil, nil, apperror.BadRequest("Event timestamp exceeded")
		}
		if payloadData.ID != "" {
			return nil, nil, apperror.BadRequest("You cannot set an id when creating a subnet")
		}

	}
	if payload.EventType == uint16(constants.UpdateSubnetEvent) {
		if payloadData.ID == "" {
			return nil, nil, apperror.BadRequest("Subnet ID must be provided")
		}
	}
	cfg, _ := (*ctx).Value(constants.ConfigKey).(*configs.MainConfiguration)

	currentState, err := service.ValidateSubnetData(&payloadData, cfg.AddressPrefix)
	if err != nil {
		return nil, nil, err
	}

	// generate associations
	if currentState != nil {
		if strings.EqualFold(currentState.Account.ToString(), payloadData.Account.ToString()) {
			return nil, nil, apperror.BadRequest("Subnet account do not match")
		}
		assocPrevEvent = &currentState.Event
	}
	if authState != nil {
		assocAuthEvent = &authState.Event
	}
	return assocPrevEvent, assocAuthEvent, nil
}
