package rpc

import (
	// "errors"
	"context"

	services "github.com/ByteGum/go-icms/pkg/service"
	utils "github.com/ByteGum/go-icms/utils"
)

type Flag string

// !sign web3 m
// type msgError struct {
// 	code int
// 	message string
// }

type RpcService struct {
	Ctx            *context.Context
	Cfg            *utils.Configuration
	MessageService *services.MessageService
}

func NewRpcService(mainCtx *context.Context) *RpcService {
	cfg, _ := (*mainCtx).Value(utils.ConfigKey).(*utils.Configuration)
	return &RpcService{
		Ctx:            mainCtx,
		Cfg:            cfg,
		MessageService: services.NewMessageService(mainCtx),
	}
}

func (p *RpcService) SendMessage(request utils.MessageJsonInput, reply *utils.ClientMessage) error {
	chatMsg := utils.CreateMessageFromJson(request)
	reply, err := (*p.MessageService).Send(chatMsg, request.Signature)
	if err != nil {
		return err
	}
	return nil
}

//! create valid outgoing channel
//! listen into incoming outgoing
//! store in the db
//! create a copy and broadcast to the network
