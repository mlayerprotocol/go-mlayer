package entities

import (
	"fmt"

	// "math"
	"strings"

	cryptoEth "github.com/ethereum/go-ethereum/crypto"

	"github.com/mlayerprotocol/go-mlayer/common/encoder"
)

type Stats struct {
	ID                 string `json:"id" gorm:"type:uuid;primaryKey;not null"`
	BlockNumber        uint64 `json:"blk"`
	EventType          uint16 `json:"t"`
	EventCount         uint64 `json:"ec"`
	MessageCount       uint64 `json:"mc"`
	TopicCount         uint64 `json:"tc"`
	AuthorizationCount uint64 `json:"ac"`
	Count              uint64 `json:"c" gorm:"default:1"`
}

func (chatStats Stats) ToString() string {
	values := []string{}

	values = append(values, string(chatStats.BlockNumber))
	values = append(values, fmt.Sprintf("%d", chatStats.EventType))
	values = append(values, fmt.Sprintf("%d", chatStats.EventCount))
	values = append(values, fmt.Sprintf("%d", chatStats.MessageCount))
	values = append(values, fmt.Sprintf("%d", chatStats.TopicCount))
	values = append(values, fmt.Sprintf("%d", chatStats.AuthorizationCount))
	// values = append(values, fmt.Sprintf("%d", chatStats.ApprovalExpiry))
	// values = append(values, fmt.Sprintf("%s", chatStats.ChainId))
	// values = append(values, fmt.Sprintf("%s", chatStats.Platform))
	// values = append(values, fmt.Sprintf("%d", chatStats.Timestamp))

	// values = append(values, fmt.Sprintf("%s", chatStats.SubjectHash))

	return strings.Join(values, "")
}

func (msg Stats) GetHash() ([]byte, error) {
	b, err := msg.EncodeBytes()
	if err != nil {
		return []byte(""), err
	}
	return cryptoEth.Keccak256Hash(b).Bytes(), nil
}

func (msg Stats) EncodeBytes() ([]byte, error) {
	var attachments []byte
	var actions []byte

	// for _, at := range msg.Actions {
	// 	attachments = append(actions, at.EncodeBytes()...)
	// }
	// for _, ac := range msg.Actions {
	// 	actions = append(actions, ac.EncodeBytes()...)
	// }

	// logger.Info("Mesage....", string(msg.Data))
	// dataByte, _ := hex.DecodeString(msg.Data)
	return encoder.EncodeBytes(
		// encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: msg.TopicId},
		// encoder.EncoderParam{Type: encoder.AddressEncoderDataType, Value: msg.Sender},
		// encoder.EncoderParam{Type: encoder.AddressEncoderDataType, Value: msg.Receiver},
		// encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: crypto.Keccak256Hash(dataByte)},
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: attachments},
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: actions},
	)
}
