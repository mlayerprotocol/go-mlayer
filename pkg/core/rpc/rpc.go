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
	ChannelService *services.ChannelService
}

type RpcResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func NewRpcService(mainCtx *context.Context) *RpcService {
	cfg, _ := (*mainCtx).Value(utils.ConfigKey).(*utils.Configuration)
	return &RpcService{
		Ctx:            mainCtx,
		Cfg:            cfg,
		MessageService: services.NewMessageService(mainCtx),
		ChannelService: services.NewChannelService(mainCtx),
	}
}

func newResponse(status string, data interface{}) *RpcResponse {
	d := RpcResponse{
		Status: status,
		Data:   data,
	}
	return &d
}

func (p *RpcService) SendMessage(request utils.MessageJsonInput, reply *RpcResponse) error {
	utils.Logger.Info("SendMessage request:::", request)
	chatMsg, err := utils.CreateMessageFromJson(request)
	if err != nil {
		return err
	}
	c, err := (*p.MessageService).Send(chatMsg, request.Signature)
	if err != nil {
		return err
	}
	reply = newResponse("success", c)
	return nil
}

func (p *RpcService) Subscription(request utils.Subscription, reply *RpcResponse) error {
	utils.Logger.Debug("Subscription request:::", request)
	err := (*p.ChannelService).ChannelSubscription(&request)
	if err != nil {
		return err
	}
	reply = newResponse("success", request)
	return nil
}
