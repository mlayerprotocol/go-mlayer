package entities

import (
	// "errors"

	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"gorm.io/gorm"
)



type WalletBalance struct {
	// Primary
	ID string `gorm:"primaryKey;type:uuid;not null" json:"id,omitempty"`
	Account           AddressString `json:"acct"  gorm:"type:varchar(64);index;not null"`
	Wallet string `json:"wal"  gorm:"type:char(32);index;not null"`
	Balance             big.Int        `json:"bal" gorm:"index;not null"`
	Timestamp            uint64                           `json:"ts"`
}

func (d *WalletBalance) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == "" {
		uuid, err := GetId(*d)
		if err != nil {
			logger.Error(err)
			panic(err)
		}

		d.ID = uuid
	}
	
	return nil
}

// func (e *WalletBalance) Key() string {
// 	return fmt.Sprintf("/%s", e.Hash)
// }

func (e *WalletBalance) ToJSON() []byte {
	m, err := json.Marshal(e)
	if err != nil {
		logger.Errorf("Unable to parse event to []byte")
	}
	return m
}

func (e *WalletBalance) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(e)
	return b
}


func WalletBalanceFromJSON(b []byte) (Event, error) {
	var e Event
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &e)
	return e, err
}

func (e WalletBalance) GetHash() ([]byte, error) {
	b, err := e.EncodeBytes()
	if err != nil {
		return []byte(""), err
	}
	return crypto.Sha256(b), nil
}

func (e WalletBalance) ToString() string {
	values := []string{}
	values = append(values, fmt.Sprintf("%s", e.Account))
	values = append(values, fmt.Sprintf("%s", e.Wallet))
	
	return strings.Join(values, "")
}

func (e WalletBalance) EncodeBytes() ([]byte, error) {

	
	return encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: e.Account},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: e.Wallet},
	)
}

