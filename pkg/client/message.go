package client

import (
	// "errors"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
)

var logger = &log.Logger

type Flag string

// !sign web3 m
// type msgError struct {
// 	code int
// 	message string
// }

type MessageService struct {
	Ctx context.Context
	Cfg configs.MainConfiguration
}

type Subscribe struct {
	channel   string
	timestamp string
}

func NewMessageService(mainCtx *context.Context) *MessageService {
	ctx := *mainCtx
	cfg, _ := ctx.Value(constants.ConfigKey).(*configs.MainConfiguration)
	return &MessageService{
		Ctx: ctx,
		Cfg: *cfg,
	}
}

func (p *MessageService) Send(chatMsg entities.ChatMessage, senderSignature string) (*entities.Event, error) {
	if strings.ToLower(chatMsg.Validator) != strings.ToLower(crypto.GetPublicKeyEDD(p.Cfg.NetworkPrivateKey)) {
		return nil, errors.New("Invalid Origin node address: " + chatMsg.Validator + " is not")
	}
	if entities.IsValidMessage(chatMsg, senderSignature) {
		channel := strings.Split(chatMsg.Header.Receiver, ":")

		if utils.Contains(chatMsg.Header.Channels, "*") || utils.Contains(chatMsg.Header.Channels, strings.ToLower(channel[0])) {

			privateKey := p.Cfg.NetworkPrivateKey

			// TODO:
			// if its an array check the channels .. if its * allow
			// message server::: store messages, require receiver to request message through an endpoint
			hash, _ := chatMsg.GetHash()
			signature, _ := crypto.SignECC(hash, privateKey)
			message := entities.Event{}
			message.Payload.Data = &chatMsg
			message.Signature = hexutil.Encode(signature)
			outgoingMessageC, ok := p.Ctx.Value(constants.OutgoingMessageChId).(*chan *entities.Event)
			if !ok {
				logger.Error("Could not connect to outgoing channel")
				panic("outgoing channel fail")
			}
			*outgoingMessageC <- &message
			fmt.Printf("Testing my function%s, %s", chatMsg.ToString(), chatMsg.Body.SubjectHash)
			return &message, nil
		}
	}
	return nil, fmt.Errorf("Invalid message signer")
}
