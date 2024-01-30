package entities

import (
	"encoding/json"
	"errors"
	"fmt"

	// "math"

	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/utils/encoder"
)


type Payload interface {
	GetHash() ([]byte, error)
	ToString()  string
	ToJSON()  []byte
	EncodeBytes()  ([]byte, error)
}

/*
map[
	d:map[
		agt:0xe652d28F89A28adb89e674a6b51852D0C341Ebe9
		authEventId:
		du:2592000000 
		eventHash: 
		gr:02ebec9d95769bb3d71712f0bf1e7e88b199fc945f67f908bbab81e9b7cb1092d8 
		h: 
		privi:3 
		sig:304402204fe2f0814dd55227d0af214b22dde326574027436fc98445ffd5d5099e28877d02207552b379ca04e9ec7aec69478576b5523fdcdfc041fc7c1038356a2b5035839b 
		topIds:* ts:1705392177894
	] 
	gr: 
	hash: 
	n: 
	sig:0x8af68d73d6296ffe34e7bd082dada6ed92e2d250fce547a73801719d4648279e7037bfc5f5c41198f869275cd97b1a46f22e2d3403dbdf4d0e6729cff41bd16a1c 
	ts:1705392177894 
	ty:100]

*/

type ClientPayload struct {
	// Primary
	Data  interface{}    `json:"d"`
	Timestamp   int       `json:"ts"`
	EventType      uint16 `json:"ty"`
	Grantor  string    `json:"gr,omitempty"` // optional public key of sender
	// Authorization *Authorization `json:"auth"`
	// AuthHash string `json:"auth"` // optional hash of 
	Validator string `json:"val,omitempty"`

	// Secondary																								 	AA	`							qaZAA	`q1aZaswq21``		`	`
	Signature   string    `json:"sig"`
	Hash   string    `json:"h,omitempty"`
	
	
}

func (msg ClientPayload) ToJSON() []byte {
	m, _ := json.Marshal(&msg)
	return m
}



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
	return fmt.Sprintf("Data: %s, EventType: %d, Authority: %s", (msg.Data).(Payload).ToString(), msg.EventType, msg.Grantor)
}

func (msg ClientPayload) GetHash() ([]byte, error) {
	b, err := msg.EncodeBytes()
	if err != nil {
		return []byte(""), err
	}
	bs := crypto.Keccak256Hash(b)
	return bs, nil
}

func (msg *ClientPayload) Validate(privateKey string) error {
	if msg.Validator  != crypto.ToBech32Address(crypto.GetPublicKeyEDD(privateKey)) {
		return errors.New("Invalid message. Message not registered to this validator")
	}
	return nil
}

// func (msg *ClientPayload) Key() string {
// 	hash, _  := msg.GetHash()
// 	return fmt.Sprintf("/%s", hex.EncodeToString(hash))
// }

func (msg ClientPayload) EncodeBytes() ([]byte, error) {
	
	
	b, _ := msg.Data.(Payload).EncodeBytes()
	
	var params []encoder.EncoderParam
	params = append(params, encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: b})
	params = append(params, encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: msg.EventType})
	if msg.Grantor != "" {
		params = append(params, encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: msg.Grantor})
	}
	params = append(params, encoder.EncoderParam{Type: encoder.AddressEncoderDataType, Value: msg.Validator})
	params = append(params, encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: msg.Timestamp})
	
	return encoder.EncodeBytes(
		params...
	)
}


func ClientPayloadFromBytes(b []byte) (ClientPayload, error) {
	var message ClientPayload
	err := json.Unmarshal(b, &message)
	return message, err
}






