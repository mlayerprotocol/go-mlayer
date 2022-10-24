package services

import (
	// "errors"
	"context"
	"errors"

	utils "github.com/ByteGum/go-icms/utils"
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

func (p *ChannelService) ChannelSubscription(sub *utils.Subscription) error {
	subscribersc, ok := p.Ctx.Value(utils.SubscribeCh).(*chan *utils.Subscription)
	if !ok {
		return errors.New("Subscription chanel not found")
	}
	sub.Broadcast = true
	*subscribersc <- sub
	return nil
}
