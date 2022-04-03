package utils

import (
	"encoding/json"
)

const (
	PlatformEthereum string = "ethereum"
	PlatformBitcoin         = "bitcoin"
	PlatformSolana          = "solana"
	PlatformCosmos          = "cosmos"
)

type MessageHeader struct {
	Length    int    `json:"length"`
	Timestamp int    `json:"timestamp"`
	Sender    string `json:"from"`
	Receiver  string `json:"to"`
	ChainId   string `json:"chainId"`
	Platform  string `json:"platform"`
}
type MessageBody struct {
	Text string `json:"text"`
	Html string `json:"html"`
}
type MessageAction struct {
	Contract   string `json:"contract"`
	Abi        string `json:"abi"`
	Action     string `json:"action"`
	Parameters string `json:"parameters"`
}

type Message struct {
	Header    MessageHeader `json:"header"`
	Body      MessageBody   `json:"body"`
	Action    MessageAction `json:"action"`
	Signature string        `json:"signature"`
}

type HandshakeData struct {
	Timestamp  int    `json:"timestamp"`
	ProtocolId string `json:"protocalId"`
}
type Handshake struct {
	Data      HandshakeData `json:"data"`
	Signature string        `json:"signature"`
}

func (hs *Handshake) toString() string {
	h, _ := json.Marshal(hs)
	return string(h)
}

func (msg *Message) toString() string {
	m, _ := json.Marshal(msg)
	return string(m)
}

func MessageFromBytes(b []byte) Message {
	var message Message
	if err := json.Unmarshal(b, &message); err != nil {
		panic(err)
	}
	return message
}

func MessageFromString(msg string) Message {
	return MessageFromBytes([]byte(msg))
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
