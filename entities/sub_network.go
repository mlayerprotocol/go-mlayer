package entities

import (
	// "errors"

	"encoding/binary"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/internal/crypto"

	"github.com/mlayerprotocol/go-mlayer/common/encoder"
)

type Subnet struct {
	ID          string        `json:"id" gorm:"type:uuid;primaryKey;not null"`
	Ref         string        `json:"ref,omitempty"`
	Name        string        `json:"n,omitempty" binding:"required"`
	Handle      string        `json:"hand,omitempty" binding:"required" gorm:"unique;type:char(64);default:null"`
	Description string        `json:"desc,omitempty"`
	Account     AddressString `json:"acct,omitempty" binding:"required"  gorm:"not null;type:varchar(100)"`

	Agent AddressString `json:"agt,omitempty" binding:"required"  gorm:"not null;type:varchar(100)"`
	//
	Public   *bool `json:"pub,omitempty" gorm:"default:false"`
	ReadOnly *bool `json:"rO,omitempty" gorm:"default:false"`
	// InviteOnly bool `json:"invO" gorm:"default:false"`

	// Derived
	Event   EventPath `json:"e,omitempty" gorm:"index;varchar;"`
	Hash    string    `json:"h,omitempty" gorm:"type:char(64)"`
	Balance float64   `json:"bal" gorm:"default:0"`
	// Signature   string    `json:"sig,omitempty" binding:"required"  gorm:"non null;"`
	// Broadcasted   bool      `json:"br,omitempty"  gorm:"default:false;"`
	Timestamp uint64 `json:"ts,omitempty" binding:"required"`
}

func (topic *Subnet) Key() string {
	return fmt.Sprintf("/%s/%s", topic.Account, topic.Hash)
}

func (topic *Subnet) ToJSON() []byte {
	m, e := json.Marshal(topic)
	if e != nil {
		logger.Errorf("Unable to parse subscription to []byte")
	}
	return m
}

func (topic *Subnet) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(topic)
	return b
}

func SubnetToByte(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)

	fmt.Println(b)
	return b
}

func SubnetFromBytes(b []byte) (Subnet, error) {
	var topic Subnet
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &topic)
	return topic, err
}
func UnpackSubnet(b []byte) (Subnet, error) {
	var topic Subnet
	err := encoder.MsgPackUnpackStruct(b, topic)
	return topic, err
}

func (p *Subnet) CanSend(channel string, sender AddressString) bool {
	// check if user can send
	return true
}

func (p *Subnet) IsMember(channel string, sender AddressString) bool {
	// check if user can send
	return true
}

func (topic Subnet) GetHash() ([]byte, error) {
	b, err := topic.EncodeBytes()
	if err != nil {
		return []byte(""), err
	}
	return crypto.Keccak256Hash(b), nil
}

func (topic Subnet) ToString() string {
	values := []string{}
	values = append(values, topic.Hash)
	values = append(values, topic.Name)
	// values = append(values, fmt.Sprintf("%d", topic.Timestamp))
	// values = append(values, fmt.Sprintf("%d", topic.SubscriberCount))
	values = append(values, string(topic.Account))
	values = append(values, fmt.Sprintf("%t", topic.Public))
	// values = append(values, fmt.Sprintf("%s", topic.Signature))
	return strings.Join(values, ",")
}

func (topic Subnet) EncodeBytes() ([]byte, error) {
	return encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: topic.ID},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: topic.Ref},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: topic.Name},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: topic.Handle},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: topic.Description},
		// encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: topic.SubscriberCount},
		// encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: topic.Account},
		encoder.EncoderParam{Type: encoder.BoolEncoderDataType, Value: *topic.Public},
		encoder.EncoderParam{Type: encoder.BoolEncoderDataType, Value: *topic.ReadOnly},
		// encoder.EncoderParam{Type: encoder.BoolEncoderDataType, Value: topic.InviteOnly},
	)
}
