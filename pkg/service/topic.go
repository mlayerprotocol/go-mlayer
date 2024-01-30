package db

import (
	// "errors"
	"context"
	"errors"

	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/db"
	"github.com/mlayerprotocol/go-mlayer/utils/constants"
)

type TopicService struct {
	Ctx context.Context
	Cfg configs.MainConfiguration
}

func NewTopicService(mainCtx *context.Context) *TopicService {
	ctx := *mainCtx
	cfg, _ := ctx.Value(constants.ConfigKey).(*configs.MainConfiguration)
	return &TopicService{
		Ctx: ctx,
		Cfg: *cfg,
	}
}

func (p *TopicService) NewTopicSubscription(sub *entities.Subscription) error {
	// subscribersc, ok := p.Ctx.Value(utils.SubscribeChId).(*chan *entities.Subscription)

	// validate before storing
	if entities.IsValidSubscription(*sub, true) {
		topicSubscriberStore, ok := p.Ctx.Value(constants.NewTopicSubscriptionStore).(*db.Datastore)
		if !ok {
			return errors.New("Could not connect to subscription datastore")
		}
		error := topicSubscriberStore.Set(p.Ctx, db.Key(sub.Key()), sub.MsgPack(), false)
		if error != nil {
			return error
		}
	}
	return nil
}
