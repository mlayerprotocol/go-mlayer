package utils

import (
	"encoding/json"
	"fmt"
)

const (
	PlatformEthereum string = "ethereum"
	PlatformBitcoin         = "bitcoin"
	PlatformSolana          = "solana"
	PlatformCosmos          = "cosmos"
)

/**
CHAT MESSAGE
**/
type NodeMessage struct {
	Timestamp uint   `json:"timestamp"`
	ChainId   string `json:"chainId"`
	Platform  string `json:"platform"`
	Message   string `json:"text"`
	Signature string `json:"signature"`
}

func (msg *NodeMessage) ToJSON() string {
	m, _ := json.Marshal(msg)
	return string(m)
}

func NodeMessageFromBytes(b []byte) NodeMessage {
	var message NodeMessage
	if err := json.Unmarshal(b, &message); err != nil {
		panic(err)
	}
	return message
}

func NodeMessageFromString(msg string) NodeMessage {
	return NodeMessageFromBytes([]byte(msg))
}

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

/**
CHAT MESSAGE
**/
type ChatMessageHeader struct {
	Length    int    `json:"length"`
	Timestamp int    `json:"timestamp"`
	Sender    string `json:"from"`
	Receiver  string `json:"receiver"`
	Channel   string `json:"channel"`
	ChainId   string `json:"chainId"`
	Platform  string `json:"platform"`
}
type ChatMessageBody struct {
	Text string `json:"text"`
	Html string `json:"html"`
}
type ChatMessageAction struct {
	Contract   string `json:"contract"`
	Abi        string `json:"abi"`
	Action     string `json:"action"`
	Parameters string `json:"parameters"`
}

type ChatMessage struct {
	Header    ChatMessageHeader `json:"header"`
	Body      ChatMessageBody   `json:"body"`
	Action    ChatMessageAction `json:"action"`
	Signature string            `json:"signature"`
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
