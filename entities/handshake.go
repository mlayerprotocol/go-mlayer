package entities

import (
	"encoding/json"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/configs"
)



type ClientHandshake struct {
	Signature string          `json:"sig"`
	Signer    DIDString          `json:"sigr"`
	// Message   string          `json:"m"`
	Protocol  constants.Protocol `json:"proto"`
	ClientSocket    *interface{} `json:"ws"`
	ChainId configs.ChainId  `json:"chId"`
	Timestamp int64 `json:"ts"`
}

type ServerIdentity struct {
	Signature string          `json:"sig"`
	Signer    string          `json:"sigr"`
	Message   string          `json:"m"`
	Address string			`json:"addr"`
}

func (cs ClientHandshake) ToJSON() []byte {
	m, e := json.Marshal(cs)
	if e != nil {
		logger.Errorf("Unable to parse subscription to []byte")
	}
	return m
}

func (hs *ClientHandshake) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(hs)
	return b
}

func (cs ClientHandshake) EncodeBytes() ([]byte, error) {
	return encoder.EncodeBytes(encoder.EncoderParam{
		Type: encoder.AddressEncoderDataType,
		Value: cs.Signer,
	}, encoder.EncoderParam{
		Type: encoder.ByteEncoderDataType,
		Value: cs.ChainId.Bytes(),
	}, encoder.EncoderParam{
		Type: encoder.IntEncoderDataType,
		Value: cs.Timestamp,
	})
}

func UnpackClientHandshake(b []byte) (ClientHandshake, error) {
	var message ClientHandshake
	err := encoder.MsgPackUnpackStruct(b, &message)
	return message, err
}

func (hs *ServerIdentity) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(hs)
	return b
}

func UnpackServerIdentity(b []byte) (ServerIdentity, error) {
	var id ServerIdentity
	err := encoder.MsgPackUnpackStruct(b, &id)
	return id, err
}



// func ClientHandshakeFromBytes(b []byte) (ClientHandshake, error) {
// 	var verMsg ClientHandshake
// 	// if err := json.Unmarshal(b, &message); err != nil {
// 	// 	panic(err)
// 	// }
// 	err := json.Unmarshal(b, &verMsg)
// 	return verMsg, err
// }



// func SignIdentity(message []byte) bool {
// 	verifiedRequest, _ := UnpackClientHandshake(message)
// 	verifiedRequest.ClientSocket = &client
// 	verifiedRequest.Protocol = protocol;
// 	log.Println("verifiedRequest.Message: ", verifiedRequest.Message)
// 	hasVerified := false
// 	if VerifySignature(verifiedRequest.Signer, verifiedRequest.Message, verifiedRequest.Signature) {
// 		// verifiedConn = append(verifiedConn, c)
// 		hasVerified = true
// 		log.Println("Verification was successful: ", verifiedRequest)
// 		*ch <- &verifiedRequest
// 	}
	
// 	return hasVerified
// }
