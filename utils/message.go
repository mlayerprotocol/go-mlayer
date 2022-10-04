package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	PlatformEthereum string = "ethereum"
	PlatformBitcoin         = "bitcoin"
	PlatformSolana          = "solana"
	PlatformCosmos          = "cosmos"
)

/**
HANDSHAKE MESSAGE
**/
type HandshakeData struct {
	Timestamp  int    `json:"timestamp"`
	ProtocolId string `json:"protocolId"`
	Name       string `json:"name"`
}

type Handshake struct {
	Data      HandshakeData `json:"data"`
	Signature string        `json:"signature"`
	Signer    string        `json:"signer"`
}

func (hs *Handshake) ToJSON() []byte {
	h, _ := json.Marshal(hs)
	return h
}
func (hs *Handshake) Init(jsonString string) error {
	er := json.Unmarshal([]byte(jsonString), &hs)
	return er
}
func (hsd *HandshakeData) ToString() string {
	return fmt.Sprintf("name:%s,timestamp:%d,protocolId:%s", hsd.Name, hsd.Timestamp, hsd.ProtocolId)
}

func (hsd *HandshakeData) ToJSON() []byte {
	h, _ := json.Marshal(hsd)
	return h
}
func HandshakeFromJSON(json string) (Handshake, error) {
	data := Handshake{}
	er := data.Init(json)
	return data, er
}

func HandshakeFromBytes(b []byte) Handshake {
	var handshake Handshake
	if err := json.Unmarshal(b, &handshake); err != nil {
		panic(err)
	}
	return handshake
}

func HandshakeFromString(hs string) Handshake {
	return HandshakeFromBytes([]byte(hs))
}

func CreateHandshake(name string, network string, privateKey string) Handshake {
	pubKey := GetPublicKey(privateKey)
	data := HandshakeData{Name: name, ProtocolId: network, Timestamp: int(time.Now().Unix())}
	_, signature := Sign((&data).ToString(), privateKey)
	return Handshake{Data: data, Signature: signature, Signer: pubKey}
}

/**
CHAT MESSAGE
**/
type ChatMessageHeader struct {
	Length    int    `json:"length"`
	Sender    string `json:"from"`
	Receiver  string `json:"reciever"`
	ChainId   string `json:"chainId"`
	Platform  string `json:"platform"`
	Timestamp uint   `json:"timestamp"`
}
type ChatMessageBody struct {
	Subject string `json:"subject"`
	Text    string `json:"text"`
	Html    string `json:"html"`
}
type ChatMessageAction struct {
	Contract   string   `json:"contract"`
	Abi        string   `json:"abi"`
	Action     string   `json:"action"`
	Parameters []string `json:"parameters"`
}

type ChatMessage struct {
	Header  ChatMessageHeader   `json:"header"`
	Body    ChatMessageBody     `json:"body"`
	Actions []ChatMessageAction `json:"actions"`
	Origin  string              `json:"origin"`
}

func (chatMessage *ChatMessage) ToString() string {
	values := []string{}
	// values = append(values, fmt.Sprintf("Header.Length:%d", chatMessage.Header.Length))
	values = append(values, fmt.Sprintf("Header.Sender:%s", chatMessage.Header.Sender))
	values = append(values, fmt.Sprintf("Header.Receiver:%s", chatMessage.Header.Receiver))
	values = append(values, fmt.Sprintf("Header.ChainId:%s", chatMessage.Header.ChainId))
	values = append(values, fmt.Sprintf("Header.Platform:%s", chatMessage.Header.Platform))
	values = append(values, fmt.Sprintf("Header.Timestamp:%d", chatMessage.Header.Timestamp))

	values = append(values, fmt.Sprintf("Body.Subject:%s", chatMessage.Body.Subject))
	values = append(values, fmt.Sprintf("Body.Text:%s", chatMessage.Body.Text))
	values = append(values, fmt.Sprintf("Body.Html:%s", chatMessage.Body.Html))
	_action := []string{}
	for i := 0; i < len(chatMessage.Actions); i++ {
		_action = append(_action, fmt.Sprintf("Actions[%d].Contract:%s", i, chatMessage.Actions[i].Contract))
		_action = append(_action, fmt.Sprintf("Actions[%d].Abi:%s", i, chatMessage.Actions[i].Abi))
		_action = append(_action, fmt.Sprintf("Actions[%d].Action:%s", i, chatMessage.Actions[i].Action))

		_parameter := []string{}
		for j := 0; j < len(chatMessage.Actions[i].Parameters); j++ {
			_parameter = append(_parameter, fmt.Sprintf("Actions[%d].Parameters[%d]:%s", i, j, chatMessage.Actions[i].Parameters[j]))
		}

		_action = append(_action, fmt.Sprintf("Actions[%d].Parameters:%s", i, _parameter))
	}

	values = append(values, fmt.Sprintf("Actions:%s", _action))
	values = append(values, fmt.Sprintf("Origin:%s", chatMessage.Origin))

	return strings.Join(values, ",")
}

/**
NODE MESSAGE
**/
type ClientMessage struct {
	Message         ChatMessage `json:"message"`
	SenderSignature string      `json:"senderSignature"`
	NodeSignature   string      `json:"nodeSignature"`
}

type SuccessResponse struct {
	body ClientMessage
	meta Meta
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

func (msg *ClientMessage) ToJSON() ([]byte, error) {
	m, err := json.Marshal(msg)
	return m, err
}

func ClientMessageFromBytes(b []byte) (ClientMessage, error) {
	var message ClientMessage
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &message)
	return message, err
}

func ClientMessageFromString(msg string) (ClientMessage, error) {
	return ClientMessageFromBytes([]byte(msg))
}

func (msg *ChatMessage) ToJSON() string {
	m, _ := json.Marshal(msg)
	return string(m)
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
	Timestamp int                 `json:"timestamp"`
	From      string              `json:"from"`
	Receiver  string              `json:"receiver"`
	Platform  string              `json:"platform"`
	ChainId   string              `json:"chainId"`
	Type      string              `json:"type"`
	Message   string              `json:"message"`
	Subject   string              `json:"subject"`
	Signature string              `json:"signature"`
	Actions   []ChatMessageAction `json:"actions"`
}

func CreateMessageFromJson(msg MessageJsonInput) ChatMessage {

	chatMessage := ChatMessageHeader{
		Timestamp: uint(msg.Timestamp),
		Sender:    msg.From,
		Receiver:  msg.Receiver,
		ChainId:   msg.ChainId,
		Platform:  msg.Platform,
		Length:    100,
	}
	var bodyMessage ChatMessageBody
	if msg.Type == "html" {
		bodyMessage = ChatMessageBody{
			Subject: msg.Subject,
			Html:    msg.Message,
		}
	} else {
		bodyMessage = ChatMessageBody{
			Subject: msg.Subject,
			Text:    msg.Message,
		}
	}

	Origin := ""
	_chatMessage := ChatMessage{chatMessage, bodyMessage, msg.Actions, Origin}
	return _chatMessage
}

func IsValidMessage(msg ChatMessage, signature string) bool {
	chatMessage := msg.ToJSON()
	signer := msg.Header.Sender
	if math.Abs(float64(int(msg.Header.Timestamp)/1000-int(time.Now().Unix()))) > VALID_HANDSHAKE_SECONDS {
		logger.WithFields(logrus.Fields{"data": chatMessage}).Warnf("ChatMessage Expired: %s", chatMessage)
		return false
	}
	message := msg.ToString()
	logger.Infof("message %s", message)
	isValid := VerifySignature(signer, message, signature)
	if !isValid {
		logger.WithFields(logrus.Fields{"message": message, "signature": signature}).Warnf("Invalid signer %s", signer)
		return false
	} else {

	}
	return true
}
