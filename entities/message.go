package entities

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"

	// "math"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	cryptoEth "github.com/ethereum/go-ethereum/crypto"

	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/utils/constants"
	"github.com/mlayerprotocol/go-mlayer/utils/encoder"
	"github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack/v5"
)

const (
	PlatformEthereum string = "ethereum"
	PlatformBitcoin         = "bitcoin"
	PlatformSolana          = "solana"
	PlatformCosmos          = "cosmos"
)

var buf bytes.Buffer
var msgPackEncoder = msgpack.NewEncoder(&buf)

func init() {
	msgPackEncoder.SetCustomStructTag("json")
}

/*
*
CHAT MESSAGE
*
*/
type ChatMessageHeader struct {
	Length        int    `json:"l"`
	Sender        string `json:"s"`
	Approval      string `json:"ap"`
	Receiver      string `json:"r"`
	// ChainId       string `json:"cId"`
	// Platform      string `json:"p"`
	Timestamp     uint64  `json:"ts"`
	ApprovalExpiry uint64    `json:"apExp"`
	// Wildcard      bool   `json:"wildcard"`
	Channels      []string `json:"chs"`
	SenderAddress string   `json:"sA"`
	// OwnerAddress  string `json:"oA"`
}

// TODO! platform enum channel
// ! receiver field is name of channel u are sending to
// ! look for all subscribers to the channel
// ! channel subscribers store
type ChatMessageBody struct {
	SubjectHash string `json:"subH"`
	MessageHash string `json:"mH"`
	CID         string `json:"cid"`
}
type ChatMessageAction struct {
	Contract   string   `json:"c"`
	Abi        string   `json:"abi"`
	Action     string   `json:"a"`
	Parameters []string `json:"pa"`
}

type ChatMessage struct {
	Header  ChatMessageHeader   `json:"h"`
	Body    ChatMessageBody     `json:"b"`
	Actions []ChatMessageAction `json:"as"`
	Validator  string              `json:"v"`
}

func (chatMessage ChatMessage) ToString() string {
	values := []string{}

	values = append(values, fmt.Sprintf("%s", chatMessage.Header.Receiver))
	// values = append(values, fmt.Sprintf("%s", chatMessage.Header.Approval))
	// values = append(values, fmt.Sprintf("%d", chatMessage.Header.ApprovalExpiry))
	// values = append(values, fmt.Sprintf("%s", chatMessage.Header.ChainId))
	// values = append(values, fmt.Sprintf("%s", chatMessage.Header.Platform))
	values = append(values, fmt.Sprintf("%d", chatMessage.Header.Timestamp))

	values = append(values, fmt.Sprintf("%s", chatMessage.Body.SubjectHash))
	values = append(values, fmt.Sprintf("%s", chatMessage.Body.MessageHash))
	values = append(values, fmt.Sprintf("%s", chatMessage.Body.CID))
	_action := []string{}
	for i := 0; i < len(chatMessage.Actions); i++ {
		_action = append(_action, fmt.Sprintf("[%d]:%s", i, chatMessage.Actions[i].Contract))
		_action = append(_action, fmt.Sprintf("[%d]:%s", i, chatMessage.Actions[i].Abi))
		_action = append(_action, fmt.Sprintf("[%d]:%s", i, chatMessage.Actions[i].Action))

		_parameter := []string{}
		for j := 0; j < len(chatMessage.Actions[i].Parameters); j++ {
			_parameter = append(_parameter, fmt.Sprintf("[%d][%d]:%s", i, j, chatMessage.Actions[i].Parameters[j]))
		}

		_action = append(_action, fmt.Sprintf("[%d]:%s", i, _parameter))
	}

	values = append(values, fmt.Sprintf("%s", _action))
	values = append(values, fmt.Sprintf("%s", chatMessage.Validator))

	return strings.Join(values, "")
}

func (msg ChatMessage) GetHash() ([]byte, error) {
	b, err := msg.EncodeBytes()
	if err != nil {
		return []byte(""), err
	}
	return cryptoEth.Keccak256Hash(b).Bytes(), nil
}

func (msg ChatMessage) EncodeBytes() ([]byte, error) {
	return encoder.EncodeBytes(encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: msg.Validator})
}

// func (channel *ChatMessageHeader) ToApprovalString() string {
// 	values := []string{}
// 	values = append(values, fmt.Sprintf("%d", channel.ApprovalExpiry))
// 	// values = append(values, fmt.Sprintf("%s", channel.Wildcard))
// 	values = append(values, fmt.Sprintf("%s", channel.Channels))
// 	values = append(values, fmt.Sprintf("%s", channel.SenderAddress))
// 	// values = append(values, fmt.Sprintf("%s", channel.OwnerAddress))
// 	return strings.Join(values, ",")
// }
func (channel *ChatMessageHeader) ToApprovalBytes() ([]byte, error) {
	values := []string{}
	values = append(values, fmt.Sprintf("%d", channel.ApprovalExpiry))
	// values = append(values, fmt.Sprintf("%s", channel.Wildcard))
	values = append(values, fmt.Sprintf("%s", channel.Channels))
	values = append(values, fmt.Sprintf("%s", channel.SenderAddress))
	// values = append(values, fmt.Sprintf("%s", channel.OwnerAddress))
	b, err := encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: channel.ApprovalExpiry},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: channel.Channels},
		encoder.EncoderParam{Type: encoder.AddressEncoderDataType, Value: channel.SenderAddress},
	)
	if(err != nil) {
		return []byte(""), err 
	}
	return b, nil
}

type SuccessResponse struct {
	Body ClientPayload
	Meta Meta
}

type ErrorResponse struct {
	statusCode int
	meta       Meta
}

type Meta struct {
	statusCode int
	success    bool
}

func ReturnError(msg string, code int) *ErrorResponse {
	meta := Meta{statusCode: code}
	meta.success = false
	e := ErrorResponse{statusCode: code}
	e.meta = meta
	return &e
}


func JsonMessageFromBytes(b []byte) (MessageJsonInput, error) {
	var message MessageJsonInput
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &message)
	return message, err
}

func UnpackJsonMessage(b []byte) (MessageJsonInput, error) {
	var message MessageJsonInput
	err := encoder.MsgPackUnpackStruct(b, message)
	return message, err
}



func (msg *ChatMessage) ToJSON() string {
	m, _ := json.Marshal(msg)
	return string(m)
}

func (msg *ChatMessage) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(msg)
	return b
}

func ChatMessageFromBytes(b []byte) ChatMessage {
	var message ChatMessage
	if err := json.Unmarshal(b, &message); err != nil {
		panic(err)
	}
	return message
}

func ChatMessageFromString(msg string) ChatMessage {
	return ChatMessageFromBytes([]byte(msg))
}

type MessageJsonInput struct {
	Timestamp     uint64      `json:"ts"`
	Approval      string   `json:"ap"`
	ApprovalExpiry	uint64      `json:"apExp"`
	Channels      []string `json:"c"`
	SenderAddress string   `json:"sA"`
	// OwnerAddress  string              `json:"oA"`
	Receiver    string              `json:"r"`
	Platform    string              `json:"p"`
	ChainId     string              `json:"cI"`
	Type        string              `json:"t"`
	Message     string              `json:"m"`
	Subject     string              `json:"s"`
	Signature   string              `json:"sig"`
	Actions     []ChatMessageAction `json:"a"`
	Origin      string              `json:"o"`
	MessageHash string              `json:"mH"`
	SubjectHash string              `json:"subH"`
}

// PubSubMessage
type PubSubMessage struct {
	Data      msgpack.RawMessage `json:"d"`
	// Timestamp uint64          `json:"ts"`
	// Signature string          `json:"sig"`
}

func (msg *PubSubMessage) ToJSON() []byte {
	m, _ := json.Marshal(msg)
	return m
}

func (msg *PubSubMessage) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(msg)
	return b
}

func (msg *PubSubMessage) ToString() string {
	values := []string{}
	values = append(values, fmt.Sprintf("Data:%s", string(msg.Data)))
	//values = append(values, fmt.Sprintf("Timestmap%d", msg.Timestamp))
	return strings.Join(values, "")
}

func (msg *PubSubMessage) EncodeBytes()  ([]byte, error) {
	return encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: []byte(msg.Data)},
		//encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: msg.Timestamp},
	)
}

func NewPubSubMessage(data []byte) PubSubMessage {
	message := PubSubMessage{Data: data}
	return message
}
func PubSubMessageFromBytes(b []byte) (PubSubMessage, error) {
	var message PubSubMessage
	err := json.Unmarshal(b, &message)
	return message, err
}

func UnpackPubSubMessage(b []byte) (PubSubMessage, error) {
	var message PubSubMessage
	err := encoder.MsgPackUnpackStruct(b, &message)
	return message, err
}

func IsValidTopic(ch ChatMessageHeader, signature string, channelOwner string) bool {
	approval, err := ch.ToApprovalBytes()
	if(err != nil) {
		return false
	}
	signer, _ := crypto.GetSignerECC(&approval, &signature)
	if strings.ToLower(channelOwner) != strings.ToLower(signer) {
		return false
	}
	if math.Abs(float64(int(ch.ApprovalExpiry)-int(time.Now().Unix()))) > constants.VALID_HANDSHAKE_SECONDS {
		logger.WithFields(logrus.Fields{"data": ch}).Warnf("Channel Expired: %d", ch.ApprovalExpiry)
		return false
	}
	isValid := crypto.VerifySignatureECC(&signer, &approval, &signature)
	if !isValid {
		logger.WithFields(logrus.Fields{"message": approval, "signature": signature}).Warnf("Invalid signer %s", signer)
		return false
	} else {

	}
	return true
}



func IsValidMessage(msg ChatMessage, signature string) bool {
	chatMessage := msg.ToJSON()
	msgByte := []byte(msg.ToString())
	signer, _ := crypto.GetSignerECC(&msgByte, &signature)
	channel := strings.Split(msg.Header.Receiver, ":")
	chaByte := []byte(strings.ToLower(channel[0]))
	channelOwner, _ := crypto.GetSignerECC(&chaByte, &(channel[1]))
	if strings.ToLower(channelOwner) != strings.ToLower(signer) {
		return false
	}
	if !IsValidTopic(msg.Header, channel[1], channelOwner) {
		return false
	}
	if math.Abs(float64(int(msg.Header.Timestamp)-int(time.Now().Unix()))) > constants.VALID_HANDSHAKE_SECONDS {
		logger.WithFields(logrus.Fields{"data": chatMessage}).Warnf("ChatMessage Expired: %s", chatMessage)
		return false
	}
	message := []byte(msg.ToString())
	isValid := crypto.VerifySignatureECC(&signer, &message, &signature)
	if !isValid {
		logger.WithFields(logrus.Fields{"message": string(message), "signature": signature}).Warnf("Invalid signer %s", signer)
		return false
	} else {

	}
	return true
}

func IsValidSubscription(
	subscription Subscription,
	verifyTimestamp bool,
) bool {
	if verifyTimestamp {
		if math.Abs(float64(int(subscription.Timestamp)-int(time.Now().Unix()))) > constants.VALID_HANDSHAKE_SECONDS {
			logger.Info("Invalid Subscription, invalid handshake duration")
			return false
		}
	}
	b, err := subscription.EncodeBytes();
	if(err != nil) {
		return false
	}
	return crypto.VerifySignatureECC(&subscription.Subscriber, &b, &subscription.Signature)
}

func CreateMessageFromJson(msg MessageJsonInput) (ChatMessage, error) {

	if len(msg.Message) > 0 {
		msgHash := hexutil.Encode(crypto.Keccak256Hash([]byte(msg.Message)))
		if msg.MessageHash != msgHash {
			return ChatMessage{}, errors.New("Invalid Message")
		}
	}
	if len(msg.Subject) > 0 {
		subHash := hexutil.Encode(crypto.Keccak256Hash([]byte(msg.Subject)))
		if msg.SubjectHash != subHash {
			return ChatMessage{}, errors.New("Invalid Subject")
		}
	}
	chatMessage := ChatMessageHeader{
		Timestamp:     uint64(msg.Timestamp),
		Approval:      msg.Approval,
		Receiver:      msg.Receiver,
		// ChainId:       msg.ChainId,
		// Platform:      msg.Platform,
		Length:        100,
		ApprovalExpiry: msg.ApprovalExpiry,
		Channels:      msg.Channels,
		SenderAddress: msg.SenderAddress,
		// OwnerAddress:  msg.OwnerAddress,
	}

	bodyMessage := ChatMessageBody{
		SubjectHash: msg.SubjectHash,
		MessageHash: msg.MessageHash,
	}
	_chatMessage := ChatMessage{chatMessage, bodyMessage, msg.Actions, msg.Origin}
	return _chatMessage, nil
}
