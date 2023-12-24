package services

import (
	// "errors"
	"context"

	utils "github.com/mlayerprotocol/go-mlayer/utils"
)

// the channel_interface interacts with the concensus layer


type ChannelInterface struct {
	Ctx context.Context
	Cfg utils.Configuration
}

type StateChannel struct {
	Name string
	Signature []byte
	Owner string
}

func (p *StateChannel) CanSend(channel string, sender []byte) bool {
	// check if user can send
	return true
}

func (p *StateChannel) IsMember(channel string, sender []byte) bool {
	// check if user can send
	return true
}

func NewStateChannelInterface(mainCtx *context.Context) *ChannelInterface {
	ctx := *mainCtx
	cfg, _ := ctx.Value(utils.ConfigKey).(*utils.Configuration)
	return &ChannelInterface{
		Ctx: ctx,
		Cfg: *cfg,
	}
}

func (p *ChannelInterface) GetChannel(channel string) StateChannel {
	return StateChannel{Name: "/test", Signature: []byte("0x20939092039023"), Owner: "0x0920902930232"}
}
