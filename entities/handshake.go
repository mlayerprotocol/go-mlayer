package entities

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
)



type ClientHandshake struct {
	Signature string          `json:"sig"`
	Signer    string          `json:"sigr"`
	Message   string          `json:"m"`
	Protocol  constants.Protocol `json:"proto"`
	ClientSocket    *interface{} `json:"ws"`
}

type ServerIdentity struct {
	Signature string          `json:"sig"`
	Signer    string          `json:"sigr"`
	Message   string          `json:"m"`
	Address string			`json:"addr"`
}

func (sub *ClientHandshake) ToJSON() []byte {
	m, e := json.Marshal(sub)
	if e != nil {
		logger.Errorf("Unable to parse subscription to []byte")
	}
	return m
}
func (hs *ClientHandshake) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(hs)
	return b
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
/*
*
NODE ANDSHAKE MESSAGE
*
*/
type HandshakeData struct {
	Timestamp  int    `json:"ts"`
	ProtocolId string `json:"proId"`
	Name       string `json:"n"`
	NodeType   uint   `json:"nT"`
}

type Handshake struct {
	Data      HandshakeData `json:"data"`
	Signature string        `json:"s"`
	Signer    string        `json:"sigr"`
}

func (hs *Handshake) ToJSON() []byte {
	h, _ := json.Marshal(hs)
	return h
}
func (hs *Handshake) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(hs)
	return b
}
func (hsd HandshakeData) ToString() string {
	return fmt.Sprintf("%s,%d,%s,%d", hsd.Name, hsd.NodeType, hsd.ProtocolId,  hsd.Timestamp)
}

func (hsd HandshakeData) EncodeBytes() ([]byte, error) {
	return encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: hsd.Name},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: hsd.NodeType},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: hsd.ProtocolId},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: hsd.Timestamp},
	)
}

func UnpackHandshake(b []byte) (Handshake, error) {
	var message Handshake
	err := encoder.MsgPackUnpackStruct(b, &message)
	return message, err
}

func (hs *Handshake) Init(jsonString string) error {
	er := json.Unmarshal([]byte(jsonString), &hs)
	return er
}

func (hsd *HandshakeData) ToJSON() []byte {
	h, _ := json.Marshal(hsd)
	return h
}
func HandshakeFromJSON(json string) (Handshake, error) {
	data := Handshake{}
	er := data.Init(json)
	return data, er
}

func HandshakeFromBytes(b []byte) Handshake {
	var handshake Handshake
	if err := json.Unmarshal(b, &handshake); err != nil {
		panic(err)
	}
	return handshake
}

func HandshakeFromString(hs string) Handshake {
	return HandshakeFromBytes([]byte(hs))
}

func CreateHandshake(name string, network string, privateKey string, nodeType uint) (Handshake, error) {
	pubKey := crypto.GetPublicKeySECP(privateKey)
	data := HandshakeData{Name: name, ProtocolId: network, NodeType: nodeType, Timestamp: int(time.Now().Unix())}
	b, err := data.EncodeBytes();
	if(err != nil) {
		return Handshake{}, err
	}
	_, signature := crypto.SignSECP(b, privateKey)
	return Handshake{Data: data, Signature: signature, Signer: pubKey}, nil
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
