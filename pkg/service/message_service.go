package services

import (
	// "errors"
	"context"
	"errors"
	"fmt"
	"strings"

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

type Subscribe struct {
	channel   string
	timestamp string
}

func NewMessageService(mainCtx *context.Context) *MessageService {
	ctx := *mainCtx
	cfg, _ := ctx.Value(utils.ConfigKey).(*utils.Configuration)
	return &MessageService{
		Ctx: ctx,
		Cfg: *cfg,
	}
}

func (p *MessageService) Send(chatMsg utils.ChatMessage, senderSignature string) (*utils.ClientMessage, error) {
	if strings.ToLower(chatMsg.Origin) != strings.ToLower(utils.GetPublicKey(p.Cfg.EvmPrivateKey)) {
		return nil, errors.New("Invalid Origin node address")
	}
	if utils.IsValidMessage(chatMsg, senderSignature) {
		privateKey := p.Cfg.EvmPrivateKey

		// TODO:
		// check message timestamp. It must be within a 15 seconds difference from the current server time

		signature, _ := utils.Sign(senderSignature, privateKey)
		message := utils.ClientMessage{}
		message.Message = chatMsg
		message.SenderSignature = senderSignature
		message.NodeSignature = hexutil.Encode(signature)
		outgoingMessageC, ok := p.Ctx.Value(utils.OutgoingMessageCh).(*chan *utils.ClientMessage)
		if !ok {
			utils.Logger.Error("Could not connect to outgoing channel")
			panic("outgoing channel fail")
		}
		*outgoingMessageC <- &message
		fmt.Printf("Testing my function%s, %s", chatMsg.ToString(), chatMsg.Body.Subject)
		return &message, nil
	}
	return nil, fmt.Errorf("Invalid message signer")
}
