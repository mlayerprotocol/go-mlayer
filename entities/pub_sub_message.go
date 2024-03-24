package entities

import (
	"encoding/json"
	"fmt"

	// "math"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/vmihailenco/msgpack/v5"
)

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

