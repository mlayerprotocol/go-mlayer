package entities

import (
	// "errors"

	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/utils/encoder"
)


var synced = "sync"


type EventInterface interface {
	EncodeBytes()  ([]byte, error)
	GetNode() string
	GetSignature() string
}

type Event struct {
	// Primary
	ID string `gorm:"primaryKey" json:"ID,omitempty"`

	Payload  ClientPayload    `json:"pld" gorm:"-"`
	Timestamp   uint64       `json:"ts"`
	EventType      uint16 `json:"t"`
	Parents   []string      `json:"p" gorm:"type:text[]"`
	// StateHash string `json:"sh"`
	// Secondary
	Hash   string    `json:"h" gorm:"unique"`
	Signature   string    `json:"sig"`
	Broadcasted   bool      `json:"br"`
	Node string		`json:"node"`
	BlockNumber  uint64  `json:"blk"`
	IsValid   bool      `json:"isVal" gorm:"default:false"`
	Synced bool      `json:"sync" gorm:"default:false"`
}


func (e *Event) Key() string {
	return fmt.Sprintf("/%s", e.Hash)
}

func (e *Event) ToJSON() []byte {
	m, err := json.Marshal(e)
	if err != nil {
		logger.Errorf("Unable to parse event to []byte")
	}
	return m
}

func (e *Event) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(e)
	return b
}

func UnpackEvent[DataType any] (b []byte, data *DataType) (*Event, error) {
	// e.Payload = payload
	e := Event{}
	err := encoder.MsgPackUnpackStruct(b, &e)
	c, err := json.Marshal(e.Payload);
	if err != nil {
		return  nil, err
	}
	pl := &ClientPayload{
		Data: data,
	}
	err = json.Unmarshal(c, &pl)
	_, err = pl.EncodeBytes()
	if err != nil {
		logger.Errorf("Unmarshal--> ERROR %v",err )
	}
	copier.Copy(e.Payload, &pl)
	newEvent := Event{
		Payload: *pl,
	}
	copier.Copy(&e.Payload, &newEvent.Payload)
	

	return &e, err
}


func EventFromJSON(b []byte) (Event, error) {
	var e Event
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &e)
	return e, err
}



func (e Event) GetHash() string {
	b, err := e.EncodeBytes()
	if err  != nil {

	}
	return hex.EncodeToString(crypto.Sha256(b))
}

func (e Event) ToString() string {
	values := []string{}
	d, _ := json.Marshal(e.Payload)
	values = append(values, fmt.Sprintf("%s", e.ID))
	values = append(values, fmt.Sprintf("%s", d))
	values = append(values, fmt.Sprintf("%d", e.EventType))
	values = append(values, fmt.Sprintf("%s", strings.Join(e.Parents, ",")))
	values = append(values, fmt.Sprintf("%d", e.Timestamp))
	return strings.Join(values, "")
}


func (e Event) EncodeBytes() ([]byte, error) {

	d, err := e.Payload.EncodeBytes()
	if err  != nil {
		return []byte(""), err
	}
	return encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: d},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: e.EventType},
		encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: strings.Join(e.Parents, "")},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: e.Timestamp},
	)
}

func (e Event) GetNode() string {
	return e.Node
}
func (e Event) 	GetSignature() string {
	return e.Signature
}

