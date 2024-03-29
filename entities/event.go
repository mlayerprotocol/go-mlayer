package entities

import (
	// "errors"

	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


var synced = "sync"


type EventModel string
const (
	AuthEventModel EventModel  = "auth"
	TopicEventModel EventModel = "top"
	SubscriptionEventModel EventModel = "sub"
)

/**
Event paths define the unique path to an event and its relation to the entitie

**/
type EventPath struct {
	Model EventModel
	Hash string
}


func (e *EventPath) ToString() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%s/%s",  e.Model, e.Hash)
}

func NewEventPath(model EventModel, hash string) (*EventPath) {
	return &EventPath{Model: model, Hash: hash}
}

func EventPathFromString(path string) (*EventPath) {
	parts := strings.Split(path, "/")
	// assoc, err := strconv.Atoi(parts[0])
	// if err != nil {
	// 	return nil, err
	// }
	switch len(parts) {
	case 0:
		return &EventPath{}
	case 1:
		return &EventPath{
			//Relationship: EventAssoc(assoc),
			Model: EventModel(""), 
			Hash: parts[0],
			}
		default:
			return &EventPath{
				//Relationship: EventAssoc(assoc),
				Model: EventModel(parts[0]), 
				Hash: parts[1],
				}
	}
}

func (eP EventPath) GormDataType() string {
	return "varchar"
}
func (eP EventPath) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	
	asString := eP.ToString()
	return clause.Expr{
	  SQL:  "?",
	  Vars: []interface{}{asString},
	  WithoutParentheses: false,
	}
  }
  
  func (sD *EventPath) Scan(value interface{}) error {
	data, ok := value.(string)
	if !ok {
	  return errors.New(fmt.Sprint("Value not instance of string:", value))
	}
  
	*sD = *EventPathFromString(data)
	return nil
  }
  
  func (sD *EventPath) Value() (driver.Value, error) {
	logger.Infof("CONVERTING2 %s", sD.ToString())
	return sD.ToString(), nil
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
	Nonce string `json:"nonce" gorm:"type:varchar(80);unique;;default:null" msgpack:",noinline"`
	Timestamp   uint64       `json:"ts"`
	EventType      uint16 `json:"t"`
	Associations   []string      `json:"assoc" gorm:"type:text[]"`
	PreviousEventHash   EventPath      `json:"preE" gorm:"type:varchar;default:null"`
	AuthEventHash   EventPath      `json:"authE" gorm:"type:varchar;default:null"`
	PayloadHash string `json:"pH" gorm:"type:char(64);unique,"`
	// StateHash string `json:"sh"`
	// Secondary
	Error string `json:"err"`
	Hash   string    `json:"h" gorm:"unique,type:char(64)"`
	Signature   string    `json:"sig"`
	Broadcasted   bool      `json:"br"`
	BlockNumber  uint64  `json:"blk"`
	IsValid   bool      `json:"isVal" gorm:"default:false"`
	Synced bool      `json:"sync" gorm:"default:false"`
	Validator PublicKeyString `json:"val"`
	InternalEvents []interface{} `json:"iEs" gorm:"_"`
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
	if d.Payload.Nonce > 0 {
		d.Nonce = fmt.Sprintf("%s:%d", string(d.Payload.Account), d.Payload.Nonce)
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
	values = append(values, fmt.Sprintf("%d", d))
	values = append(values,  e.ID)
	values = append(values, fmt.Sprintf("%d", e.BlockNumber))
	values = append(values, fmt.Sprintf("%d", e.EventType))
	values = append(values, strings.Join(e.Associations, ","))
	values = append(values, e.PreviousEventHash.ToString())
	values = append(values,  e.AuthEventHash.ToString())
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
		encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: e.PreviousEventHash.Hash},
		encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: e.AuthEventHash.Hash},
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

