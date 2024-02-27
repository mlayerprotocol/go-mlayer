package client

import (
	// "errors"

	"github.com/mlayerprotocol/go-mlayer/entities"
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

func GetTopic(id string) (*models.TopicState, error) {
	topicState := models.TopicState{}

	err := query.GetOne(models.TopicState{
		Topic: entities.Topic{Hash: id},
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

// func ListenForNewTopicEventFromPubSub (mainCtx *context.Context) {
// 	ctx, cancel := context.WithCancel(*mainCtx)
// 	defer cancel()

// 	incomingTopicC, ok := (*mainCtx).Value(constants.IncomingTopicEventChId).(*chan *entities.Event)
// 	if !ok {
// 		logger.Errorf("incomingTopicC closed")
// 		return
// 	}
// 	for {
// 		event, ok :=  <-*incomingTopicC
// 		if !ok {
// 			logger.Fatal("incomingTopicC closed for read")
// 			return
// 		}
// 		go service.HandleNewPubSubTopicEvent(event, ctx)
// 	}
// }
