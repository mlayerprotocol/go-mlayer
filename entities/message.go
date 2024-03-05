package entities

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	// "math"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	cryptoEth "github.com/ethereum/go-ethereum/crypto"

	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
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
	Length   int    `json:"l"`
	Sender   AddressString `json:"s"`
	Receiver string `json:"r"`
	// ChainId       string `json:"cId"`
	// Platform      string `json:"p"`
	Timestamp      uint64 `json:"ts"`
	// ApprovalExpiry uint64 `json:"apExp"`
	// Wildcard      bool   `json:"wildcard"`
	// Channels      []string `json:"chs"`
	// SenderAddress string   `json:"sA"`
	// OwnerAddress  string `json:"oA"`
}

func (h ChatMessageHeader) EncodeBytes() []byte {
	b, _ := encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: h.Length},
		encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: h.Sender},
		encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: h.Receiver},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: h.Timestamp},
	)
	return b
}

// TODO! platform enum channel
// ! receiver field is name of channel u are sending to
// ! look for all subscribers to the channel
// ! channel subscribers store
type ChatMessageBody struct {
	DataHash string `json:"mH"`
	Url         string `json:"url"`
	Data 		json.RawMessage `json:"d,omitempty"`
}
func (b ChatMessageBody) EncodeBytes() []byte {
	e, _ := encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: b.DataHash},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: b.Url},
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: b.Data},
	)
	return e
}

type ChatMessageAttachment struct {
	CID string `json:"cid"`
	Hash         string `json:"h"`
}
func (b ChatMessageAttachment) EncodeBytes() []byte {
	e, _ := encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: b.CID},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: b.Hash},
	)
	return e
}
type ChatMessageAction struct {
	Contract   string   `json:"c"`
	Abi        string   `json:"abi"`
	Action     string   `json:"a"`
	Parameters []string `json:"pa"`
}

func (a ChatMessageAction) EncodeBytes() []byte {
	var b []byte
	for _, d := range a.Parameters {
		data, _ := encoder.EncodeBytes(encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: d},)
		b = append(b, data...)
	}
	encoded, _ := encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: a.Contract},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: a.Abi},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: a.Action},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: a.Action},
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: b},
	)
	return encoded
}

type ChatMessage struct {
	Header    ChatMessageHeader   `json:"h"`
	Body      ChatMessageBody     `json:"b"`
	Actions   []ChatMessageAction `json:"as"`
	Attachments []ChatMessageAttachment `json:"atts"`
}

func (chatMessage ChatMessage) ToString() string {
	values := []string{}

	values = append(values, fmt.Sprintf("%s", chatMessage.Header.Receiver))
	// values = append(values, fmt.Sprintf("%s", chatMessage.Header.Approval))
	// values = append(values, fmt.Sprintf("%d", chatMessage.Header.ApprovalExpiry))
	// values = append(values, fmt.Sprintf("%s", chatMessage.Header.ChainId))
	// values = append(values, fmt.Sprintf("%s", chatMessage.Header.Platform))
	// values = append(values, fmt.Sprintf("%d", chatMessage.Header.Timestamp))

	// values = append(values, fmt.Sprintf("%s", chatMessage.Body.SubjectHash))
	values = append(values, fmt.Sprintf("%s", chatMessage.Body.DataHash))
	values = append(values, fmt.Sprintf("%s", chatMessage.Body.Url))
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
	var attachments []byte
	var actions []byte
	
	for _, at := range msg.Actions {
		attachments = append(actions, at.EncodeBytes()...)
	}
	for _, ac := range msg.Actions {
		actions = append(actions, ac.EncodeBytes()...)
	}
	return encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: msg.Header.EncodeBytes()},
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: msg.Body.EncodeBytes()},
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: attachments},
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: actions},
	)
}

// func (channel ChatMessageHeader) ToApprovalBytes() ([]byte, error) {
// 	values := []string{}
// 	values = append(values, fmt.Sprintf("%d", channel.ApprovalExpiry))
// 	// values = append(values, fmt.Sprintf("%s", channel.Wildcard))
// 	values = append(values, fmt.Sprintf("%s", channel.Channels))
// 	values = append(values, fmt.Sprintf("%s", channel.SenderAddress))
// 	// values = append(values, fmt.Sprintf("%s", channel.OwnerAddress))
// 	b, err := encoder.EncodeBytes(
// 		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: channel.ApprovalExpiry},
// 		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: channel.Channels},
// 		encoder.EncoderParam{Type: encoder.AddressEncoderDataType, Value: channel.SenderAddress},
// 	)
// 	if err != nil {
// 		return []byte(""), err
// 	}
// 	return b, nil
// }

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
type MessageJsonInputAttachments struct {
	File []json.RawMessage `json:"f"`
	Type string `json:"ty"`
}
type MessageJsonInput struct {
	Timestamp      uint64   `json:"ts"`
	// Approval       string   `json:"ap"`
	// ApprovalExpiry uint64   `json:"apExp"`
	// Channels       []string `json:"c"`
	TopicId  string `json:"topId"`
	Sender  AddressString   `json:"s"`
	// OwnerAddress  string              `json:"oA"`
	Receiver    string              `json:"r"`
	// Platform    string              `json:"p"`
	// ChainId     string              `json:"cI"`
	Type        string              `json:"t"`
	Data    []byte 				`json:"d"`
	// Subject     string              `json:"s"`
	Signature   string              `json:"sig"`
	Actions     []ChatMessageAction `json:"a"`
	// Origin      string              `json:"o"`
	DataHash string              `json:"dH"`
	Url string              `json:"url"`
	Length int `json:"len"`
	Attachments []MessageJsonInputAttachments `json:"atts"`
}

// PubSubMessage
type PubSubMessage struct {
	Data msgpack.RawMessage `json:"d"`
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

func (msg *PubSubMessage) EncodeBytes() ([]byte, error) {
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


// func IsValidSubscription(
// 	subscription Subscription,
// 	verifyTimestamp bool,
// ) bool {
// 	if verifyTimestamp {
// 		if math.Abs(float64(int(subscription.Timestamp)-int(time.Now().UnixMilli()))) > constants.VALID_HANDSHAKE_SECONDS {
// 			logger.Info("Invalid Subscription, invalid handshake duration")
// 			return false
// 		}
// 	}
// 	b, err := subscription.EncodeBytes()
// 	if err != nil {
// 		return false
// 	}
// 	return crypto.VerifySignatureECC(string(subscription.Subscriber), &b, subscription.Signature)
// }

func (msg MessageJsonInput) ToChatMessage() (*ChatMessage, error) {

	if len(msg.Data) > 0 {
		msgHash := hexutil.Encode(crypto.Keccak256Hash(msg.Data))
		if msg.DataHash != msgHash {
			return nil, errors.New("INVALID MESSAGE")
		}
	}
	// if len(msg.Subject) > 0 {
	// 	subHash := hexutil.Encode(crypto.Keccak256Hash([]byte(msg.Subject)))
	// 	if msg.SubjectHash != subHash {
	// 		return ChatMessage{}, errors.New("Invalid Subject")
	// 	}
	// }
	chatMessage := ChatMessageHeader{
		Timestamp: uint64(msg.Timestamp),
		Receiver:  msg.Receiver,
		// ChainId:       msg.ChainId,
		// Platform:      msg.Platform,
		Length:         msg.Length,
		Sender:  msg.Sender,
		// OwnerAddress:  msg.OwnerAddress,
	}

	bodyMessage := ChatMessageBody{
		DataHash: msg.DataHash,
		Data: msg.Data,
		Url: msg.Url,
	}
	_chatMessage := ChatMessage{Header: chatMessage, Body: bodyMessage, Actions: msg.Actions,}
	return &_chatMessage, nil
}
