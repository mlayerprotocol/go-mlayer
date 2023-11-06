package utils

import (
	"encoding/json"
	"fmt"
	"strconv"

	// "math"
	"strings"
)

// DeliveryProof
type DeliveryProof struct {
	MessageHash   string `json:"mH"`
	MessageSender string `json:"mS"`
	NodeAddress   string `json:"nA"`
	Timestamp     int    `json:"ts"`
	Signature     string `json:"sig"`
	Block         string `json:"bl"`
	Index         int    `json:"i"`
}

func (msg *DeliveryProof) ToJSON() []byte {
	m, _ := json.Marshal(msg)
	return m
}
func (msg *DeliveryProof) Pack() []byte {
	b, _ := MsgPackStruct(msg)
	return b
}

func (msg *DeliveryProof) Key() string {
	return fmt.Sprintf("/%s/%s", msg.MessageHash, msg.MessageSender)
}
func (msg *DeliveryProof) BlockKey() string {
	return fmt.Sprintf("/%s", msg.Block)
}

func (msg *DeliveryProof) ToString() string {
	values := []string{}
	values = append(values, fmt.Sprintf("%s", string(msg.MessageHash)))
	values = append(values, fmt.Sprintf("%s", msg.NodeAddress))
	values = append(values, fmt.Sprintf("%s", strconv.Itoa(msg.Timestamp)))
	return strings.Join(values, ",")
}

// func NewSignedDeliveryProof(data []byte, privateKey string) DeliveryProof {
// 	message, _ := DeliveryProofFromBytes(data)
// 	_, sig := Sign(message.ToString(), privateKey)
// 	message.Signature = sig
// 	return message
// }

func DeliveryProofFromBytes(b []byte) (DeliveryProof, error) {
	var message DeliveryProof
	err := json.Unmarshal(b, &message)
	return message, err
}

// DeliveryClaim
type DeliveryClaim struct {
	NodeHeight int      `json:"nh"`
	Signature  string   `json:"sig"`
	Amount     string   `json:"a"`
	Proofs     []string `json:"prs"`
}

func (msg *DeliveryClaim) ToJSON() []byte {
	m, _ := json.Marshal(msg)
	return m
}

func DeliveryClaimFromBytes(b []byte) (DeliveryClaim, error) {
	var message DeliveryClaim
	err := json.Unmarshal(b, &message)
	return message, err
}

func UnpackDelvieryClaim(b []byte) (DeliveryClaim, error) {
	var message DeliveryClaim
	err := MsgPackUnpackStruct(b, &message)
	return message, err
}
