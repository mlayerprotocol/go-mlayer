package services

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
	ctx := *mainCtx
	cfg, _ := ctx.Value(utils.ConfigKey).(utils.Configuration)
	return &MessageService{
		Ctx: ctx,
		Cfg: cfg,
	}
}

func (p *MessageService) Send(chatMsg utils.ChatMessage, senderSignature string) (message *utils.ClientMessage, e error) {
	if utils.IsValidMessage(chatMsg, senderSignature) {
		privateKey := p.Cfg.EvmPrivateKey
		utils.Logger.Infof("private key %s", privateKey)

		// TODO:
		// get public key from private key
		if chatMsg.Origin == utils.GetPublicKey(privateKey) {
			panic("Invalid origin")
		}

		signature, _ := utils.Sign(senderSignature, privateKey)
		message.Message = chatMsg
		message.SenderSignature = senderSignature
		message.NodeSignature = hexutil.Encode(signature)
		outgoingMessageC, ok := p.Ctx.Value(utils.OutgoingMessageCh).(chan *utils.ClientMessage)
		if !ok {
			utils.Logger.Error("Could not connect to outgoing channel")
			panic("outgoing channel fail")
		}
		outgoingMessageC <- message
		fmt.Printf("Testing my function%s, %s", chatMsg.ToString(), chatMsg.Body.Subject)
		return message, nil
	}
	return nil, fmt.Errorf("Invalid message signer")
}
