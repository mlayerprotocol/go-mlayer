package crypto

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	cryptoSha256 "crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
	"github.com/sirupsen/logrus"
)

var logger = &log.Logger

func GetPublicKeyECC(privKey string) string {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		logger.Fatalf("Invalid private key %o", err)
	}
	return crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
}

func GetPublicKeySECP(privKey string) string {
	privateKey, err := hex.DecodeString(privKey)
	if err != nil  {
		logger.Errorf("Invlaid node network key %v", err)
	}
	_, pub := btcec.PrivKeyFromBytes(btcec.S256(), privateKey)
	
	return hex.EncodeToString(pub.SerializeCompressed())
}

func GetPublicKeyEDD(privKey string) string {
	if len(privKey) != 128{
		logger.Fatal("Invalid private key length")
    }
	return privKey[64:]
}

func NetworkPrivateKeyFromString(privKey string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		logger.Fatalf("Invalid private key %o", err)
		return nil, err
	}
	return privateKey, nil
}

func Keccak256Hash(message []byte) []byte {
	messageHash := crypto.Keccak256Hash(message)
	_bytes := []byte{0x19}
	_bytes = append(_bytes, []byte("Ethereum Signed Message:\n32")...)
	_bytes = append(_bytes, messageHash.Bytes()...)
	return crypto.Keccak256Hash(_bytes).Bytes()
}


func SignECC(message []byte, privKey string) ([]byte, string) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		logger.Fatalf("Invalid private key %o",  err)
	}
	
	hash := Keccak256Hash(message)
	// logger.WithFields(logrus.Fields{"action": "crypto.Sign", "message": message}).Infof("Message hash: %s", hash.Hex())
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		logger.Fatal(err)
	}
	// signer, err := crypto.Ecrecover(hash.Bytes(), signature[:len(signature)-1])
	return signature, hexutil.Encode(signature)
}

func SignEDD(message []byte, privKey string) ([]byte, string) {
	key, err := hex.DecodeString(privKey)
	if err != nil {
		logger.Fatalf("Invalid private key string %o", err)
	}
	pKey := ed25519.PrivateKey(key)
	hash := Sha256(message)
	// logger.WithFields(logrus.Fields{"action": "crypto.Sign", "message": message}).Infof("Message hash: %s", hash.Hex())
	signature := ed25519.Sign(pKey, hash)
	if err != nil {
		logger.Fatal(err)
	}
	// signer, err := crypto.Ecrecover(hash.Bytes(), signature[:len(signature)-1])
	return signature, hex.EncodeToString(signature)
}

func SignSECP(message []byte, privKey string) ([]byte, string) {
	privateKeyByte, err := hex.DecodeString(privKey)
	if err != nil {
		logger.Fatalf("Invalid private key %o",  err)
	}
	privateKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privateKeyByte)
	signature, err := privateKey.Sign(Sha256(message)[:])
    if err != nil {
        logger.Fatal(err)
    }
	b := signature.Serialize();
	return b, hex.EncodeToString(b)
}

func Sha256 (s []byte)  []byte {
    h := cryptoSha256.New()
    h.Write(s)
    return h.Sum(nil)
}


func GetSignerECC(message *[]byte, signature *string) (string, error) {
	decoded, err := hexutil.Decode(*signature)
	if err != nil {
		logger.Debug(err)
		return "", err
	}
	hash := Keccak256Hash(*message)
	if decoded[crypto.RecoveryIDOffset] == 27 || decoded[crypto.RecoveryIDOffset] == 28 {
		decoded[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}
	signer, err := crypto.SigToPub(hash, decoded)
	if err != nil {
		return "", err
	}
	return crypto.PubkeyToAddress(*signer).Hex(), nil
}

func VerifySignatureECC(signer *string, message *[]byte, signature *string) bool {
	decodedSigner, err := GetSignerECC(message, signature)
	if err != nil {
		return false
	}
	println("signer decoded signer %s %s : %T", decodedSigner, signer, (decodedSigner == *signer))
	return strings.EqualFold(decodedSigner, *signer)
}

func VerifySignatureEDD(signer string, message *[]byte, signature string) (bool, error) {
	signatureByte, err := hex.DecodeString(signature)
	if err != nil {
		logger.WithFields(logrus.Fields{"signature": signature}).Infof("Unable to decode signature %v", err)
		return false, err
	}
	publicKeyBytes, err := hex.DecodeString(signer)
	if err != nil {
		logger.WithFields(logrus.Fields{"signer": signer}).Infof("Unable to decode signer %v", err)
		return false, err
	}
	
	msg := Sha256(*message)
	return  ed25519.Verify(publicKeyBytes, msg[:], signatureByte), nil
}

func VerifySignatureSECP(signer string, message *[]byte, signature string) (bool, error) {
	signatureByte, err := hex.DecodeString(signature)
	if err != nil {
		logger.WithFields(logrus.Fields{"signature": signature}).Infof("Unable to decode signature %v", err)
		return false, err
	}
	publicKeyBytes, err := hex.DecodeString(signer)
	if err != nil {
		logger.WithFields(logrus.Fields{"signer": signer}).Infof("Unable to decode signer %v", err)
		return false, err
	}
	publicKey, err := btcec.ParsePubKey(publicKeyBytes, btcec.S256())
    if err != nil {
		logger.WithFields(logrus.Fields{"key": string(publicKeyBytes)}).Infof("Unable to parse pubkey %v", err)
		return false, err
    }
	parsedSign, err := btcec.ParseSignature(signatureByte, btcec.S256())
	if err != nil {
		logger.WithFields(logrus.Fields{"signature": string(signatureByte)}).Infof("Unable to parse signature %v", err)
		return false, err
	}
	msg := Sha256(*message)
	return parsedSign.Verify(msg[:], publicKey), nil
}

func ToBech32Address(publicKey string) string {
	parts := strings.Split(publicKey, ":")
	if len(parts) > 1 {
		publicKey = parts[1]
	}
	b, err := hex.DecodeString(publicKey)
	if err != nil {
		logger.Fatal("Failed decoding public key", err)
	}
    shaHash := Sha256(b)
    publicHash := btcutil.Hash160(shaHash)

    // Convert to Bech32
    bech32Address, err := bech32.ConvertAndEncode("ml:", publicHash)
    if err != nil {
        logger.Fatal("Error converting to Bech32:", err)
    }

    return fmt.Sprintf("%s", bech32Address)
}

