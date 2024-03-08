package entities

import (
	"encoding/json"
	"fmt"

	// "math"

	"github.com/google/uuid"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
)

type Payload interface {
	GetHash() ([]byte, error)
	ToString() string
	EncodeBytes() ([]byte, error)
}

func GetId (d Payload) (string, error) {
	hash, err := d.GetHash()
	if err != nil {
		return "", err
	 }
	 u, err := uuid.FromBytes(hash[:16])
	 if err != nil {
	   return "", err
	 }
	 return u.String(), nil
  }

type ClientPayload struct {
	// Primary
	Data      interface{}     `json:"d"`
	Timestamp int             `json:"ts"`
	EventType uint16          `json:"ty"`
	Nonce     uint64          `json:"nonce"`
	Account   PublicKeyString `json:"acct,omitempty"` // optional public key of sender
	// Authorization *Authorization `json:"auth"`
	// AuthHash string `json:"auth"` // optional hash of
	Validator PublicKeyString `json:"val,omitempty"`
	// Secondary																								 	AA	`							qaZAA	`q1aZaswq21``		`	`
	Signature string `json:"sig"`
	Hash      string `json:"h,omitempty"`
	Agent string `gorm:"-" json:"-"`
}

func (msg ClientPayload) ToJSON() []byte {
	m, _ := json.Marshal(&msg)
	return m
}


// func (msg ClientPayload) EventNonce() string {
// 	return fmt.Sprintf("%s:%s", string(msg.Account), msg.Nonce)
// 	// d, err := msg.EncodeBytes()
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// agent, _ := crypto.GetSignerECC(&d, &msg.Signature)
// 	// return hex.EncodeToString(crypto.Keccak256Hash([]byte(fmt.Sprintf("%s:%d",  agent, msg.Nonce))))
// 	return 
// }
// func (s *ClientPayload) Encode() []byte {
// 	b, _ := s.Data.ToString()
// 	return b
// }

func (s *ClientPayload) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(s)
	return b
}

func MsgUnpackClientPayload(b []byte) (ClientPayload, error) {
	var p ClientPayload
	err := encoder.MsgPackUnpackStruct(b, p)
	return p, err
}

func (msg ClientPayload) ToString() string {
	return fmt.Sprintf("Data: %s, EventType: %d, Authority: %s", (msg.Data).(Payload).ToString(), msg.EventType, msg.Account)
}

func (msg ClientPayload) GetHash() ([]byte, error) {
	b, err := msg.EncodeBytes()
	if err != nil {
		return []byte(""), err
	}
	bs := crypto.Keccak256Hash(b)
	return bs, nil
}

func (msg ClientPayload) GetSigner() (string, error) {
	
	if len(msg.Agent) == 0 {
		b, err := msg.EncodeBytes()
		if err != nil {
			return "", err
		}
		msg.Agent, _ = crypto.GetSignerECC(&b, &msg.Signature)
		return msg.Agent, nil
	}
	return msg.Agent, nil
}

// func (msg *ClientPayload) Validate(pubKey PublicKeyString) error {
// 	if string(msg.Validator)  != string(pubKey) {
// 		// logger.Infof("VALIDIATOR %s %s, %s", msg.Validator, crypto.GetPublicKeyEDD(privateKey), crypto.ToBech32Address(crypto.GetPublicKeyEDD(privateKey)))
// 		return errors.New("Invalid message. Message not registered to this validator")
// 	}
// 	_, err := msg.EncodeBytes()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (msg *ClientPayload) Key() string {
// 	hash, _  := msg.GetHash()
// 	return fmt.Sprintf("/%s", hex.EncodeToString(hash))
// }

func (msg ClientPayload) EncodeBytes() ([]byte, error) {

	b, err := msg.Data.(Payload).EncodeBytes()
	if err != nil {
		return []byte(""), err
	}
	
	hashed := crypto.Keccak256Hash(b)
	// logger.Info("ENCODED==== ", hex.EncodeToString(b), " ", hex.EncodeToString(hashed))
	var params []encoder.EncoderParam
	params = append(params, encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: hashed})
	params = append(params, encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: msg.EventType})
	if msg.Account != "" {
		params = append(params, encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: msg.Account})
	}
	params = append(params, encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: msg.Validator})
	params = append(params, encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: msg.Nonce})
	params = append(params, encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: msg.Timestamp})

	return encoder.EncodeBytes(
		params...,
	)
}

func ClientPayloadFromBytes(b []byte) (ClientPayload, error) {
	var message ClientPayload
	err := json.Unmarshal(b, &message)
	return message, err
}
