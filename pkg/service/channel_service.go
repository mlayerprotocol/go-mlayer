package services

import (
	// "errors"
	"context"
	"errors"

	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/db"
	utils "github.com/mlayerprotocol/go-mlayer/utils"
)

type ChannelService struct {
	Ctx context.Context
	Cfg utils.Configuration
}

func NewChannelService(mainCtx *context.Context) *ChannelService {
	ctx := *mainCtx
	cfg, _ := ctx.Value(utils.ConfigKey).(*utils.Configuration)
	return &ChannelService{
		Ctx: ctx,
		Cfg: *cfg,
	}
}

func (p *ChannelService) NewChannelSubscription(sub *entities.SubscriptionEvent) error {
	// subscribersc, ok := p.Ctx.Value(utils.SubscribeChId).(*chan *entities.SubscriptionEvent)

	// validate before storing
	if entities.IsValidSubscription(*sub, true) {
		channelSubscriberStore, ok := p.Ctx.Value(utils.NewChannelSubscriptionStore).(*db.Datastore)
		if !ok {
			return errors.New("Could not connect to subscription datastore")
		}
		error := channelSubscriberStore.Set(p.Ctx, db.Key(sub.Key()), sub.Pack(), false)
		if error != nil {
			return error
		}
	}
	return nil
}
