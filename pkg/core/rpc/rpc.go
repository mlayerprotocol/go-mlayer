package rpc

import (
	// "errors"
	"context"
	"fmt"

	utils "github.com/ByteGum/go-icms/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Flag string

// !sign web3 m
// type msgError struct {
// 	code int
// 	message string
// }

type MessageService struct {
	Ctx context.Context
	Cfg utils.Configuration
}

func NewMessageService(mainCtx *context.Context) *MessageService {
	ctx, _ := context.WithCancel(*mainCtx)
	cfg, _ := ctx.Value("Config").(utils.Configuration)
	return &MessageService{
		Ctx: ctx,
		Cfg: cfg,
	}
}

func (p *MessageService) Send(request utils.MessageJsonInput, reply *utils.ClientMessage) error {
	chatMsg := utils.CreateMessageFromJson(request)
	if utils.IsValidMessage(chatMsg, request.Signature) {
		privateKey := p.Cfg.PrivateKey
		utils.Logger.Infof("private key %s", privateKey)

		// TODO
		// get public key from private key
		// check if chatMsg.Origin == crypto.GetPublicKey(privateKey)

		signature, _ := utils.Sign(request.Signature, privateKey)
		reply.Message = chatMsg
		reply.SenderSignature = request.Signature
		reply.NodeSignature = hexutil.Encode(signature)
		outgoingMessageC, ok := p.Ctx.Value("OutgoingMessageC").(chan utils.ClientMessage)
		if !ok {
			utils.Logger.Error("Could not connect to outgoing channel")
			panic("outgoing channel fail")
		}
		outgoingMessageC <- *reply
		fmt.Printf("Testing my function%s, %s", chatMsg.ToString(), chatMsg.Body.Subject)
		return nil
	}
	return nil
}

//! create valid outgoing channel
//! listen into incoming outgoing
//! store in the db
//! create a copy and broadcast to the network
