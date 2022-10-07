package utils

import (
	"crypto/ecdsa"
	"crypto/sha256"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var logger = Logger

func GetPublicKey(privKey string) string {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		logger.Fatalf("Invalid private key %s %w", privKey, err)
	}
	return crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
}

func EvmPrivateKeyFromString(privKey string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		logger.Fatalf("Invalid private key %s %w", privKey, err)
		return nil, err
	}
	return privateKey, nil
}

func hashMessage(message string) []byte {
	messageHash := crypto.Keccak256Hash([]byte(message))
	_bytes := []byte{0x19}
	_bytes = append(_bytes, []byte("Ethereum Signed Message:\n32")...)
	_bytes = append(_bytes, messageHash.Bytes()...)
	return crypto.Keccak256Hash(_bytes).Bytes()
}

func CreateHash256(message string) string {
	h := sha256.New()

	h.Write([]byte(message))

	bs := h.Sum(nil)
	return string(bs)
}

func Sign(message string, privKey string) ([]byte, string) {

	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		logger.Fatalf("Invalid private key %s %w", privKey, err)
	}

	hash := hashMessage(message)

	// logger.WithFields(logrus.Fields{"action": "crypto.Sign", "message": message}).Infof("Message hash: %s", hash.Hex())

	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		logger.Fatal(err)
	}
	// signer, err := crypto.Ecrecover(hash.Bytes(), signature[:len(signature)-1])

	return signature, hexutil.Encode(signature)
}

func GetSigner(message string, signature string) (string, error) {
	decoded, err := hexutil.Decode(signature)
	if err != nil {
		logger.Debug(err)
		return "", err
	}
	hash := hashMessage(message)
	if decoded[crypto.RecoveryIDOffset] == 27 || decoded[crypto.RecoveryIDOffset] == 28 {
		decoded[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}
	signer, err := crypto.SigToPub(hash, decoded)
	if err != nil {
		return "", err
	}
	return crypto.PubkeyToAddress(*signer).Hex(), nil
}

func VerifySignature(signer string, message string, signature string) bool {
	// hash := crypto.Keccak256Hash([]byte(message))
	// signatureBytes := hexutil.MustDecode(signature)
	// signatureNoRecoverID := signatureBytes[:len(signature)-1]
	// publicKey, err := hexutil.Decode(signer)
	// if err != nil {
	// 	logger.Fatal(err)
	// }
	// return crypto.VerifySignature(publicKey, hash.Bytes(), signatureNoRecoverID)
	decodedSigner, err := GetSigner(message, signature)
	if err != nil {
		return false
	}
	return decodedSigner == signer
}
