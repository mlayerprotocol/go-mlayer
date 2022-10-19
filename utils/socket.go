package utils

import (
	// "errors"

	"encoding/json"

	"github.com/gorilla/websocket"
)

type VerificationRequest struct {
	Signature string          `json:"signature"`
	Signer    string          `json:"signer"`
	Message   string          `json:"message"`
	Socket    *websocket.Conn `json:"socket"`
}

func (sub *VerificationRequest) ToJSON() []byte {
	m, e := json.Marshal(sub)
	if e != nil {
		logger.Errorf("Unable to parse subscription to []byte")
	}
	return m
}

func VerificationRequestFromBytes(b []byte) (VerificationRequest, error) {
	var verMsg VerificationRequest
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &verMsg)
	return verMsg, err
}
