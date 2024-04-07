package entities

import (
	// "errors"

	"encoding/json"
	"fmt"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
)

var logger = &log.Logger

type Subscription struct {
	ID         string                         `gorm:"primaryKey;type:char(36);not null"  json:"id,omitempty"`
	Topic      string                         `json:"top"`
	Account    AddressString                  `json:"sub"`
	Timestamp  uint64                         `json:"ts"`
	Signature  string                         `json:"sig"`
	Hash       string                         `json:"h" gorm:"unique" `
	Event  EventPath                         `json:"e" gorm:"index;char(64);"`
	Agent      AddressString                  `json:"agt,omitempty" binding:"required"  gorm:"not null;type:varchar(100);index"`
	Status     constants.SubscriptionStatuses `json:"st"  gorm:"not null;type:tinyint;index"`
	Role       constants.SubscriberPrivilege  `json:"rol" gorm:"default:0"`
}

func (sub *Subscription) Key() string {
	return fmt.Sprintf("/%s/%s", sub.Account, sub.Topic)
}

func (sub Subscription) ToJSON() []byte {
	m, e := json.Marshal(sub)
	if e != nil {
		logger.Errorf("Unable to parse subscription to []byte")
	}
	return m
}

func (subscription Subscription) ToString() string {
	values := []string{}
	values = append(values, subscription.Hash)
	values = append(values, subscription.ID)
	// values = append(values, fmt.Sprintf("%d", subscription.Timestamp))
	values = append(values, string(subscription.Account))
	values = append(values, fmt.Sprintf("%d", subscription.Timestamp))
	return strings.Join(values, ",")
}

func (sub Subscription) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(sub)
	return b
}

func SubscriptionFromBytes(b []byte) (Subscription, error) {
	var sub Subscription
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &sub)
	return sub, err
}
func UnpackSubscription(b []byte) (Subscription, error) {
	var sub Subscription
	err := encoder.MsgPackUnpackStruct(b, sub)
	return sub, err
}

func (sub Subscription) GetHash() ([]byte, error) {
	b, err := sub.EncodeBytes()
	if err != nil {
		log.Logger.Errorf("Subscription Hashing error, %v", err)
		return []byte(""), err
	}
	return crypto.Keccak256Hash(b), nil
}

// func (sub *Subscription) ToString() string {
// 	values := []string{}
// 	values = append(values, fmt.Sprintf("%s", sub.Topic))
// 	// values = append(values, fmt.Sprintf("%s", sub.ChannelName))
// 	values = append(values, fmt.Sprintf("%d", sub.Timestamp))
// 	values = append(values, fmt.Sprintf("%d", sub.Action))
// 	return strings.Join(values, "")
// }

func (sub Subscription) EncodeBytes() ([]byte, error) {
	// var buffer bytes.Buffer
	// buffer.Write([]byte(sub.Topic))
	// buffer.Write(encoder.NumberToByte(sub.Timestamp))
	// buffer.Write(encoder.NumberToByte(uint64(sub.Action)))
	logger.Info("SUBTOPIC", sub.Topic)
	return encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: sub.Topic},
		// encoder.EncoderParam{Type: encoder.AddressEncoderDataType, Value: sub.Subscriber},
		// encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: sub.Event},
		// encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: sub.Status},
		// encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: sub.Timestamp},
	)
}
