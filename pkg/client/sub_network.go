package client

import (
	// "errors"

	"encoding/json"
	"errors"
	"time"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/service"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	query "github.com/mlayerprotocol/go-mlayer/internal/sql/query"
	"gorm.io/gorm"
)

// type SubNetworkService struct {
// 	Ctx context.Context
// 	Cfg configs.MainConfiguration
// }

// func NewSubNetworkService(mainCtx *context.Context) *SubNetworkService {
// 	ctx := *mainCtx
// 	cfg, _ := ctx.Value(constants.ConfigKey).(*configs.MainConfiguration)
// 	return &SubNetworkService{
// 		Ctx: ctx,
// 		Cfg: *cfg,
// 	}
// }

// func (p *SubNetworkService) NewSubNetworkSubscription(sub *entities.Subscription) error {
// 	// subscribersc, ok := p.Ctx.Value(utils.SubscribeChId).(*chan *entities.Subscription)

// 	// validate before storing
// 	if entities.IsValidSubscription(*sub, true) {
// 		subNetworkSubscriberStore, ok := p.Ctx.Value(constants.NewSubNetworkSubscriptionStore).(*db.Datastore)
// 		if !ok {
// 			return errors.New("Could not connect to subscription datastore")
// 		}
// 		error := subNetworkSubscriberStore.Set(p.Ctx, db.Key(sub.Key()), sub.MsgPack(), false)
// 		if error != nil {
// 			return error
// 		}
// 	}
// 	return nil
// }

/*
Validate and Process the subNetwork request
*/

func GetSubNetworkById(id string) (*models.SubNetworkState, error) {
	subNetworkState := models.SubNetworkState{}

	err := query.GetOne(models.SubNetworkState{
		SubNetwork: entities.SubNetwork{ID: id},
	}, &subNetworkState)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &subNetworkState, nil

}
func GetSubNetworkByHash(hash string) (*models.SubNetworkState, error) {
	subNetworkState := models.SubNetworkState{}

	err := query.GetOne(models.SubNetworkState{
		SubNetwork: entities.SubNetwork{Hash: hash},
	}, &subNetworkState)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &subNetworkState, nil

}

func GetSubNetworks() (*[]models.SubNetworkState, error) {
	var subNetworkStates []models.SubNetworkState

	err := query.GetMany(models.SubNetworkState{}, &subNetworkStates)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &subNetworkStates, nil
}

func GetSubNetworkEvents() (*[]models.SubNetworkEvent, error) {
	var subNetworkEvents []models.SubNetworkEvent

	err := query.GetMany(models.SubNetworkEvent{
		Event: entities.Event{
			BlockNumber: 1,
		},
	}, &subNetworkEvents)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &subNetworkEvents, nil
}

// func ListenForNewSubNetworkEventFromPubSub (mainCtx *context.Context) {
// 	ctx, cancel := context.WithCancel(*mainCtx)
// 	defer cancel()

//		incomingSubNetworkC, ok := (*mainCtx).Value(constants.IncomingSubNetworkEventChId).(*chan *entities.Event)
//		if !ok {
//			logger.Errorf("incomingSubNetworkC closed")
//			return
//		}
//		for {
//			event, ok :=  <-*incomingSubNetworkC
//			if !ok {
//				logger.Fatal("incomingSubNetworkC closed for read")
//				return
//			}
//			go service.HandleNewPubSubSubNetworkEvent(event, ctx)
//		}
//	}
func ValidateSubNetworkPayload(payload entities.ClientPayload, authState *models.AuthorizationState) (assocPrevEvent *entities.EventPath, assocAuthEvent *entities.EventPath, err error) {

	payloadData := entities.SubNetwork{}
	d, _ := json.Marshal(payload.Data)
	e := json.Unmarshal(d, &payloadData)
	if e != nil {
		logger.Errorf("UnmarshalError %v", e)
	}

	payload.Data = payloadData
	if payload.EventType == uint16(constants.CreateSubNetworkEvent) {
		// dont worry validating the AuthHash for Authorization requests
		if uint64(payloadData.Timestamp) > uint64(time.Now().UnixMilli())+15000 {
			return nil, nil, errors.New("Authorization timestamp exceeded")
		}

	}

	currentState, err := service.ValidateSubNetworkData(&payloadData)
	if err != nil {
		return nil, nil, err
	}

	// generate associations
	if currentState != nil {
		assocPrevEvent = &currentState.Event

	}
	if authState != nil {
		assocAuthEvent = &authState.Event
	}
	return assocPrevEvent, assocAuthEvent, nil
}
