package utils

import (
	"encoding/json"
	"fmt"
	"time"
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

/**
NODE MESSAGE
**/
type NodeMessage struct {
	Message   ChatMessage `json:"message"`
	Signature string      `json:"signature"`
}

func (msg *NodeMessage) ToJSON() ([]byte, error) {
	m, err := json.Marshal(msg)
	return m, err
}

func NodeMessageFromBytes(b []byte) (NodeMessage, error) {
	var message NodeMessage
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &message)
	return message, err
}

func NodeMessageFromString(msg string) (NodeMessage, error) {
	return NodeMessageFromBytes([]byte(msg))
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
