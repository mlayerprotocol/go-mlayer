package entities

import (
	// "errors"

	"encoding/binary"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/internal/crypto"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
)

type Topic struct {
	ID string `json:"id" gorm:"type:uuid;primaryKey;not null"`
	// Name            string        `json:"n,omitempty" binding:"required"`
	Ref             string        `json:"ref,omitempty" binding:"required" gorm:"uniqueIndex:idx_unique_subnet_ref;type:varchar(64);default:null"`
	Meta            string        `json:"meta,omitempty"`
	ParentTopicHash string        `json:"pTH,omitempty" gorm:"type:char(64)"`
	SubscriberCount uint64        `json:"sC,omitempty"`
	Account         DIDString `json:"acct,omitempty" binding:"required"  gorm:"not null;type:varchar(100)"`

	Agent DeviceString `json:"agt,omitempty" binding:"required"  gorm:"not null;type:varchar(100)"`
	//
	Public   *bool `json:"pub,omitempty" gorm:"default:false"`

	DefaultSubscriberRole   *constants.SubscriberRole `json:"dSubRol,omitempty"`

	ReadOnly *bool `json:"rO,omitempty" gorm:"default:false"`
	// InviteOnly bool `json:"invO" gorm:"default:false"`

	// Derived
	Event   EventPath `json:"e,omitempty" gorm:"index;varchar;"`
	Hash    string    `json:"h,omitempty" gorm:"type:char(64)"`
	// Signature   string    `json:"sig,omitempty" binding:"required"  gorm:"non null;"`
	// Broadcasted   bool      `json:"br,omitempty"  gorm:"default:false;"`
	Timestamp uint64 `json:"ts,omitempty" binding:"required"`
	Subnet    string `json:"snet" gorm:"uniqueIndex:idx_unique_subnet_ref;type:char(36);"`
	// Subnet string `json:"snet" gorm:"index;varchar(36)"`
	BlockNumber uint64          `json:"blk"`
	Cycle   	uint64			`json:"cy"`
	Epoch		uint64			`json:"ep"`
}

func (topic *Topic) Key() string {
	return fmt.Sprintf("/%s/%s", topic.Account, topic.Hash)
}

func (topic *Topic) ToJSON() []byte {
	m, e := json.Marshal(topic)
	if e != nil {
		logger.Errorf("Unable to parse subscription to []byte")
	}
	return m
}

func (topic *Topic) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(topic)
	return b
}

func TopicToByte(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)

	return b
}

func TopicFromBytes(b []byte) (Topic, error) {
	var topic Topic
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &topic)
	return topic, err
}
func UnpackTopic(b []byte) (Topic, error) {
	var topic Topic
	err := encoder.MsgPackUnpackStruct(b, &topic)
	return topic, err
}

func (p *Topic) CanSend(channel string, sender DIDString) bool {
	// check if user can send
	return true
}

func (p *Topic) IsMember(channel string, sender DIDString) bool {
	// check if user can send
	return true
}

func (topic Topic) GetHash() ([]byte, error) {
	// if topic.Hash != "" {
	// 	return hex.DecodeString(topic.Hash)
	// }
	b, err := topic.EncodeBytes()
	if err != nil {
		return []byte(""), err
	}
	return crypto.Sha256(b), nil
}

func (topic Topic) ToString() string {
	values := []string{}
	values = append(values, topic.Hash)
	values = append(values, topic.Meta)
	// values = append(values, fmt.Sprintf("%d", topic.Timestamp))
	values = append(values, fmt.Sprintf("%d", topic.SubscriberCount))
	values = append(values, string(topic.Account))
	values = append(values, fmt.Sprintf("%t", topic.Public))
	// values = append(values, fmt.Sprintf("%s", topic.Signature))
	return strings.Join(values, ",")
}

func (topic Topic) GetEvent() EventPath {
	return topic.Event
}
func (topic Topic) GetAgent() DeviceString {
	return topic.Agent
}

func (topic Topic) EncodeBytes() ([]byte, error) {
	return encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: utils.SafePointerValue(topic.DefaultSubscriberRole, 0)},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: topic.ID},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: topic.Meta},
		encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: topic.ParentTopicHash},
		encoder.EncoderParam{Type: encoder.BoolEncoderDataType, Value: utils.SafePointerValue(topic.Public, false)},
		encoder.EncoderParam{Type: encoder.BoolEncoderDataType, Value:  utils.SafePointerValue(topic.ReadOnly, false)},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: topic.Ref},
		// encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: *topic.DefaultSubscriptionStatus},
		// encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: topic.Subnet},
	)
}

