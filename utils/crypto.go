package utils

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
)

var logger = Logger()

func GetPublicKey(privKey string) string {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		logger.Fatalf("Invalid private key %s %w", privKey, err)
	}
	return crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
}

func PrivateKeyFromString(privKey string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		logger.Fatalf("Invalid private key %s %w", privKey, err)
		return nil, err
	}
	return privateKey, nil
}

func Sign(message string, privKey string) ([]byte, string) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		logger.Fatalf("Invalid private key %s %w", privKey, err)
	}

	hash := crypto.Keccak256Hash([]byte(message))
	logger.WithFields(logrus.Fields{"action": "crypto.Sign", "message": message}).Infof("Message hash: %s", hash.Hex())

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
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
	hash := crypto.Keccak256Hash([]byte(message))
	signer, err := crypto.SigToPub(hash.Bytes(), decoded)
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
