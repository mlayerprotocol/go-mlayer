package entities

import (
	// "errors"

	"encoding/binary"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lib/pq"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
)

type Subnet struct {
	ID            string        `json:"id" gorm:"type:uuid;primaryKey;not null"`
	Meta          string        `json:"meta,omitempty"`
	Ref           string        `json:"ref,omitempty"  gorm:"unique;type:varchar(64);default:null"`
	Categories    pq.Int32Array `gorm:"type:integer[]"`
	Owner         DIDString     `json:"own,omitempty" binding:"required"  gorm:"not null;type:varchar(100);default:did"`
	SignatureData SignatureData `json:"sigD" gorm:"json;"`
	Status        *uint8         `json:"st" gorm:"boolean;default:0"`
	Timestamp     uint64        `json:"ts,omitempty" binding:"required"`
	Balance       uint64        `json:"bal" gorm:"default:0"`
	// Readonly
	Account DIDString    `json:"acct,omitempty" binding:"required"  gorm:"not null;type:varchar(100)"`
	Agent   DeviceString `json:"_"  gorm:"_"`

	// CreateTopicPrivilege   *constants.AuthorizationPrivilege `json:"cTopPriv"` // 
	DefaultAuthPrivilege   *constants.AuthorizationPrivilege `json:"dAuthPriv"` // privilege for external users who joins the subnet. 0 indicates people cant join

	// Derived
	Event EventPath `json:"e,omitempty" gorm:"index;varchar;"`
	Hash  string    `json:"h,omitempty" gorm:"type:char(64)"`
}

func (item *Subnet) Key() string {
	return fmt.Sprintf("/%s/%s", item.Account, item.Hash)
}

func (item *Subnet) ToJSON() []byte {
	m, e := json.Marshal(item)
	if e != nil {
		logger.Errorf("Unable to parse subscription to []byte")
	}
	return m
}

func (item *Subnet) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(item)
	return b
}

func SubnetToByte(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)

	fmt.Println(b)
	return b
}

func SubnetFromBytes(b []byte) (Subnet, error) {
	var item Subnet
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &item)
	return item, err
}
func UnpackSubnet(b []byte) (Subnet, error) {
	var item Subnet
	err := encoder.MsgPackUnpackStruct(b, item)
	return item, err
}

func (p *Subnet) CanSend(channel string, sender DIDString) bool {
	// check if user can send
	return true
}

func (p *Subnet) IsMember(channel string, sender DIDString) bool {
	// check if user can send
	return true
}

func (item Subnet) GetHash() ([]byte, error) {
	b, err := item.EncodeBytes()
	if err != nil {
		return []byte(""), err
	}
	logger.Infof("GetHash crypto.Sha256(b) : %v", crypto.Sha256(b))
	return crypto.Sha256(b), nil
}

func (item Subnet) ToString() string {
	values := []string{}
	values = append(values, item.Hash)
	values = append(values, item.Meta)
	// values = append(values, fmt.Sprintf("%d", item.Timestamp))
	// values = append(values, fmt.Sprintf("%d", item.SubscriberCount))
	values = append(values, string(item.Account))
	// values = append(values, fmt.Sprintf("%s", item.Signature))
	return strings.Join(values, ",")
}

func (entity Subnet) GetEvent() EventPath {
	return entity.Event
}
func (entity Subnet) GetAgent() DeviceString {
	return entity.Agent
}

func (item Subnet) EncodeBytes() ([]byte, error) {
	cats := []byte{}
	for _, d := range item.Categories {
		b, err := encoder.EncodeBytes(encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: d})
		if err != nil {
			return cats, err
		}
		cats = append(cats, b...)
	}
	return encoder.EncodeBytes(

		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: utils.SafePointerValue(item.DefaultAuthPrivilege, 0)},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: item.Meta},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: item.Owner},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: item.Ref},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: utils.SafePointerValue(item.Status, 0)},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: item.Timestamp},
	)
}
