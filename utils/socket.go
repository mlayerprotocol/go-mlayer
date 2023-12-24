package utils

import (
	// "errors"

	"encoding/json"
	"log"
)


type ClientHandshake struct {
	Signature string          `json:"sig"`
	Signer    string          `json:"sigr"`
	Message   string          `json:"m"`
	Protocol  Protocol `json:"proto"`
	ClientSocket    *interface{} `json:"ws"`
}

func (sub *ClientHandshake) ToJSON() []byte {
	m, e := json.Marshal(sub)
	if e != nil {
		logger.Errorf("Unable to parse subscription to []byte")
	}
	return m
}
func (hs *ClientHandshake) Pack() []byte {
	b, _ := MsgPackStruct(hs)
	return b
}

func UnpackClientHandshake(b []byte) (ClientHandshake, error) {
	var message ClientHandshake
	err := MsgPackUnpackStruct(b, &message)
	return message, err
}
// func ClientHandshakeFromBytes(b []byte) (ClientHandshake, error) {
// 	var verMsg ClientHandshake
// 	// if err := json.Unmarshal(b, &message); err != nil {
// 	// 	panic(err)
// 	// }
// 	err := json.Unmarshal(b, &verMsg)
// 	return verMsg, err
// }

func ConnectClient(message []byte, protocol Protocol,  client interface{}, ch *chan *ClientHandshake) bool {
		verifiedRequest, _ := UnpackClientHandshake(message)
		verifiedRequest.ClientSocket = &client
		verifiedRequest.Protocol = protocol;
		log.Println("verifiedRequest.Message: ", verifiedRequest.Message)
		hasVerified := false
		if VerifySignature(verifiedRequest.Signer, verifiedRequest.Message, verifiedRequest.Signature) {
			// verifiedConn = append(verifiedConn, c)
			hasVerified = true
			log.Println("Verification was successful: ", verifiedRequest)
			*ch <- &verifiedRequest
		}
		
		return hasVerified
}
