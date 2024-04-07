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

// type TopicService struct {
// 	Ctx context.Context
// 	Cfg configs.MainConfiguration
// }

// func NewTopicService(mainCtx *context.Context) *TopicService {
// 	ctx := *mainCtx
// 	cfg, _ := ctx.Value(constants.ConfigKey).(*configs.MainConfiguration)
// 	return &TopicService{
// 		Ctx: ctx,
// 		Cfg: *cfg,
// 	}
// }

// func (p *TopicService) NewTopicSubscription(sub *entities.Subscription) error {
// 	// subscribersc, ok := p.Ctx.Value(utils.SubscribeChId).(*chan *entities.Subscription)

// 	// validate before storing
// 	if entities.IsValidSubscription(*sub, true) {
// 		topicSubscriberStore, ok := p.Ctx.Value(constants.NewTopicSubscriptionStore).(*db.Datastore)
// 		if !ok {
// 			return errors.New("Could not connect to subscription datastore")
// 		}
// 		error := topicSubscriberStore.Set(p.Ctx, db.Key(sub.Key()), sub.MsgPack(), false)
// 		if error != nil {
// 			return error
// 		}
// 	}
// 	return nil
// }

/*
Validate and Process the topic request
*/

func GetTopicById(id string) (*models.TopicState, error) {
	topicState := models.TopicState{}

	err := query.GetOne(models.TopicState{
		Topic: entities.Topic{ID: id},
	}, &topicState)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &topicState, nil

}
func GetTopicByHash(hash string) (*models.TopicState, error) {
	topicState := models.TopicState{}

	err := query.GetOne(models.TopicState{
		Topic: entities.Topic{Hash: hash},
	}, &topicState)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &topicState, nil

}

func GetTopics() (*[]models.TopicState, error) {
	var topicStates []models.TopicState

	err := query.GetMany(models.TopicState{}, &topicStates)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &topicStates, nil
}

func GetTopicEvents() (*[]models.TopicEvent, error) {
	var topicEvents []models.TopicEvent

	err := query.GetMany(models.TopicEvent{
		Event: entities.Event{
			BlockNumber: 1,
		},
	}, &topicEvents)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &topicEvents, nil
}

// func ListenForNewTopicEventFromPubSub (mainCtx *context.Context) {
// 	ctx, cancel := context.WithCancel(*mainCtx)
// 	defer cancel()

//		incomingTopicC, ok := (*mainCtx).Value(constants.IncomingTopicEventChId).(*chan *entities.Event)
//		if !ok {
//			logger.Errorf("incomingTopicC closed")
//			return
//		}
//		for {
//			event, ok :=  <-*incomingTopicC
//			if !ok {
//				logger.Fatal("incomingTopicC closed for read")
//				return
//			}
//			go service.HandleNewPubSubTopicEvent(event, ctx)
//		}
//	}
func ValidateTopicPayload(payload entities.ClientPayload, authState *models.AuthorizationState) (assocPrevEvent *entities.EventPath, assocAuthEvent *entities.EventPath, err error) {

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
			return nil, nil, errors.New("Authorization timestamp exceeded")
		}

	}

	currentState, err := service.ValidateTopicData(&payloadData)
	if err != nil {
		return nil, nil, err
	}

	// generate associations
	if currentState != nil {
		assocPrevEvent = &currentState.Event

	}
	if authState != nil {
		assocAuthEvent =&authState.Event
	}
	return assocPrevEvent, assocAuthEvent, nil
}
