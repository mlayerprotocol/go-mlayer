package entities

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	// "math"

	"github.com/google/uuid"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
)

type Payload interface {
	GetHash() ([]byte, error)
	ToString() string
	EncodeBytes() ([]byte, error)
	GetEvent() EventPath
}

func GetId(d Payload) (string, error) {
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
	Data      interface{}   `json:"d"`
	Timestamp uint64           `json:"ts"`
	EventType uint16        `json:"ty"`
	Nonce     uint64        `json:"nonce"`
	Account   DIDString `json:"acct,omitempty"` // optional public key of sender
	ChainId   configs.ChainId `json:"chId"` // optional public key of sender

	Validator string `json:"val,omitempty"`
	// Secondary																								 	AA	`							qaZAA	`q1aZaswq21``		`	`
	Signature string       `json:"sig"`
	Hash      string       `json:"h,omitempty"`
	Agent     DeviceString `gorm:"-" json:"agt"`
	Subnet    string       `json:"snet" gorm:"index;"`
	Page      uint16       `json:"page,omitempty" gorm:"_"`
	PerPage   uint16       `json:"perPage,omitempty" gorm:"_"`
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
	err := encoder.MsgPackUnpackStruct(b, &p)
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
	logger.Infof("BYTESSSS: %s", hex.EncodeToString(b))
	bs := crypto.Keccak256Hash(b)
	return bs, nil
}

func (msg ClientPayload) GetSigner() (DeviceString, error) {

	//if len(msg.Agent) == 0 {
		b, err := msg.EncodeBytes()
		logger.Info("ENCODEDBBBBB", " ", hex.EncodeToString(b), " ", hex.EncodeToString(crypto.Keccak256Hash(b)), " Err: ", err)
		if err != nil {
			return "", err
		}
		agent, _ := crypto.GetSignerECC(&b, &msg.Signature)
		msg.Agent = AddressFromString(agent).ToDeviceString()
		return msg.Agent, nil
	//}
	// return msg.Agent, nil
}
// 0000000000014a34f22033dbd9823243a3ae6ab8b42bacec84688a267d750a028e51d46e16d3f4ea00000000000005156469643a307833466436454344434432323563334445306530373342333337433463424143353334326532414338ddb466a5dd4a5c0835614c7a46e18943ef750a9d000000000000000000000191bdd35250
// 0000000000014a34f22033dbd9823243a3ae6ab8b42bacec84688a267d750a028e51d46e16d3f4ea00000000000005156469643a307833466436454344434432323563334445306530373342333337433463424143353334326532414338ddb466a5dd4a5c0835614c7a46e18943ef750a9d000000000000000000000191bdd35250
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
	hashed := []byte("")
	if msg.Data != nil {
		b, err := msg.Data.(Payload).EncodeBytes()
		if err != nil {
			return []byte(""), err
		}
		hashed = crypto.Keccak256Hash(b)
		logger.Info("ENCODED==== ", hex.EncodeToString(b), " HASHED==== ", hex.EncodeToString(hashed))
	}

	var params []encoder.EncoderParam
	logger.Infof("ChainID %s", hex.EncodeToString(msg.ChainId.Bytes()))
	params = append(params, encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: msg.ChainId.Bytes()})
	params = append(params, encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: hashed})
	params = append(params, encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: msg.EventType})
	if msg.Subnet != "" {
		params = append(params, encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: utils.UuidToBytes(msg.Subnet)})
	}
	if msg.Account != "" {
		params = append(params, encoder.EncoderParam{Type: encoder.AddressEncoderDataType, Value: msg.Account})
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

/** SYNC REQUEST PAYLOAD **/
type SyncRequest struct {
	Interval ResponseInterval `json:"inter"`
	TopicIds string           `json:"topIds"`
}


// 0000000000007a69 684f9c7f9ba3f01e94c4599362a3266b7fa9d92b569dd2b53127f676b525ca36 00000000000004b133c981003bd204eb9fdfb17cb32a08416469643a636f736d6f73317a3770757836706574663666766e67646b6170306370796e657a746a3577776d6c76377a39662bd13f459be8acf7750609db22c4c3daa55fa9509f43e54ad59a0c38107db372000000000000000000000191227711b8
// 0000000000007a69 793e5e986a8a3e29ca7b75d04df840740d948c2fa37d9e57c2ed8ed520021899 00000000000004b133c981003bd204eb9fdfb17cb32a08416469643a636f736d6f73317a3770757836706574663666766e67646b6170306370796e657a746a3577776d6c76377a39662bd13f459be8acf7750609db22c4c3daa55fa9509f43e54ad59a0c38107db372000000000000000000000191227711b8