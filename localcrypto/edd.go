package localcrypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
)

func GenerateKeyPair() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return publicKey, privateKey, nil
}

func SignMessage(privateKey ed25519.PrivateKey, message []byte) []byte {
	return ed25519.Sign(privateKey, message)
}

func VerifyMessage(publicKey ed25519.PublicKey, message, signature []byte) bool {
	return ed25519.Verify(publicKey, message, signature)
}

func KeyToString(key []byte) string {
	return base64.StdEncoding.EncodeToString(key)
}

func StringToKey(keyStr string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(keyStr)
}
