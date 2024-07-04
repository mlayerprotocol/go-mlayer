package entities

import (
	// "errors"

	"encoding/json"
	"fmt"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
)

var logger = &log.Logger

type Subscription struct {

	ID         string    `gorm:"primaryKey;type:char(36);not null"  json:"id,omitempty"`
	Topic      string    `json:"top" binding:"required"  gorm:"not null;uniqueIndex:idx_sub_topic;type:char(36);index"`
	Ref        string    `json:"ref" gorm:"unique;type:varchar(100);default:null"`
	Meta       string    `json:"meta"  gorm:"type:varchar(100);"`
	Subnet     string    `json:"snet"  binding:"required" gorm:"not null;type:varchar(36)"`
	Subscriber DIDString `json:"sub"  gorm:"not null;uniqueIndex:idx_sub_topic;type:varchar(100);index"`
	// Device     DeviceString                  `json:"dev,omitempty" binding:"required"  gorm:"not null;uniqueIndex:idx_acct_dev_topic;type:varchar(100);index"`
	Status *constants.SubscriptionStatus `json:"st"  gorm:"not null;type:smallint;default:2"`
	Role   *constants.SubscriberRole  `json:"rol" gorm:"default:0"`

	//Signature string                         `json:"sig"`
	Timestamp uint64       `json:"ts"`
	Hash      string       `json:"h" gorm:"unique" `
	Event     EventPath    `json:"e" gorm:"index;char(64);"`
	Agent     DeviceString `json:"agt,omitempty" binding:"required"  gorm:"not null;type:varchar(100);index"`

}

func (sub *Subscription) Key() string {
	return fmt.Sprintf("/%s/%s", sub.Subscriber, sub.Topic)
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
	values = append(values, string(subscription.Subscriber))
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

func (sub Subscription) GetEvent() EventPath {
	return sub.Event
}
func (sub Subscription) GetAgent() DeviceString {
	return sub.Agent
}

func (sub Subscription) EncodeBytes() ([]byte, error) {
	return encoder.EncodeBytes(

		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: sub.Meta},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: sub.Ref},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: utils.SafePointerValue(sub.Role, 0)},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: utils.SafePointerValue(sub.Status, 0)},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: sub.Subscriber.ToString()},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: sub.Topic},
	)
}
