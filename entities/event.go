package entities

import (
	// "errors"

	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"gorm.io/gorm"
)


var synced = "sync"

type EventAssoc uint8
const (
	AuthorizationEventAssoc EventAssoc  = 1
	PreviousEventAssoc EventAssoc = 2
)

type EventModel string
const (
	AuthorizationEventModel EventModel  = "auth"
	TopicEventModel EventModel = "topic"
)

/**
Event paths define the unique path to an event and its relation to the entitie

**/
type EventPath struct {
	Relationship EventAssoc
	Model EventModel
	Hash string
}

func (e EventPath) ToString() string {
	return fmt.Sprintf("%d/%s/%s", e.Relationship, e.Model, e.Hash)
}

func EventPathFromString(path string) (*EventPath, error) {
	parts := strings.Split(path, "/")
	assoc, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	return &EventPath{
		Relationship: EventAssoc(assoc),
		Model: EventModel(parts[1]), 
		Hash: parts[2],
		}, nil
}

type EventInterface interface {
	EncodeBytes()  ([]byte, error)
	GetValidator() PublicKeyString
	GetSignature() string
}

type Event struct {
	// Primary
	ID string `gorm:"primaryKey;type:uuid;not null" json:"id,omitempty"`

	Payload  ClientPayload    `json:"pld" gorm:"serializer:json" msgpack:",noinline"`
	Timestamp   uint64       `json:"ts"`
	EventType      uint16 `json:"t"`
	Associations   []string      `json:"assoc" gorm:"type:text[]"`
	PreviousEventHash   string      `json:"preE" gorm:"type:char(64)"`
	AuthEventHash   string      `json:"authE" gorm:"type:char(64)"`
	PayloadHash string `json:"pH" gorm:"type:varchar(64);index:,"`
	// StateHash string `json:"sh"`
	// Secondary
	Error string `json:"err"`
	Hash   string    `json:"h" gorm:"unique,type:varchar(64)"`
	Signature   string    `json:"sig"`
	Broadcasted   bool      `json:"br"`
	BlockNumber  uint64  `json:"blk"`
	IsValid   bool      `json:"isVal" gorm:"default:false"`
	Synced bool      `json:"sync" gorm:"default:false"`
	Validator PublicKeyString `json:"val"`
}

func (d *Event) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == ""  {
		uuid, err := GetId(*d)
		if err != nil {
			logger.Error(err)
			panic(err)
		}
		
		d.ID = uuid
	}
	return nil
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
	if err != nil {
		logger.Errorf("UnmarshalError:: %o", err )
	}
	logger.Infof("PL:: %v", pl.Data )
	_, err2 := pl.EncodeBytes()
	if err2 != nil {
		logger.Errorf("EncodeBytesError:: %o", err )
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



func (e Event) GetHash() ([]byte, error) {
	b, err := e.EncodeBytes()
	if err  != nil {
		return []byte(""), err
	}
	return crypto.Sha256(b), nil
}

func (e Event) ToString() string {
	values := []string{}
	d, _ := json.Marshal(e.Payload)
	values = append(values, fmt.Sprintf("%s", e.ID))
	values = append(values, fmt.Sprintf("%s", d))
	values = append(values, fmt.Sprintf("%d", e.EventType))
	values = append(values, fmt.Sprintf("%s", strings.Join(e.Associations, ",")))
	values = append(values, fmt.Sprintf("%s", e.PreviousEventHash))
	values = append(values, fmt.Sprintf("%s", e.AuthEventHash))
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
		encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: strings.Join(e.Associations, "")},
		encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: e.PreviousEventHash},
		encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: e.AuthEventHash},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: e.BlockNumber},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: e.Timestamp},
	)
}

func (e Event) GetValidator() PublicKeyString {
	return e.Validator
}
func (e Event) 	GetSignature() string {
	return e.Signature
}

