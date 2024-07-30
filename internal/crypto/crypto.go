package crypto

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	cryptoSha256 "crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
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
	if err != nil {
		logger.Errorf("Invlaid node network key %v", err)
	}
	_, pub := btcec.PrivKeyFromBytes(btcec.S256(), privateKey)

	return hex.EncodeToString(pub.SerializeCompressed())
}

func GetPublicKeyEDD(privKey string) string {
	if len(privKey) != 128 {
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
	return messageHash.Bytes()
}

func SignECC(message []byte, privKey string) ([]byte, string) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		logger.Fatalf("Invalid private key %o", err)
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
		logger.Fatalf("Invalid private key %o", err)
	}
	privateKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privateKeyByte)
	signature, err := privateKey.Sign(Sha256(message)[:])
	if err != nil {
		logger.Fatal(err)
	}
	b := signature.Serialize()
	return b, hex.EncodeToString(b)
}

func Sha256(s []byte) []byte {
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
	logger.Info("Keccak256Hash Hash:: ", hex.EncodeToString(hash))
	if decoded[crypto.RecoveryIDOffset] == 27 || decoded[crypto.RecoveryIDOffset] == 28 {
		decoded[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	signer, err := crypto.SigToPub(hash, decoded)
	if err != nil {
		return "", err
	}
	return crypto.PubkeyToAddress(*signer).Hex(), nil
}

func VerifySignatureECC(signer string, message *[]byte, signature string) bool {
	decodedSigner, err := GetSignerECC(message, &signature)
	fmt.Printf("message decoded signer %s %s %s : %v === %s", message, decodedSigner, signer, (decodedSigner == signer), signature)
	if err != nil {
		return false
	}

	return strings.EqualFold(decodedSigner, signer)
}

func VerifySignatureEDD(signer string, message *[]byte, signature string) (bool, error) {
	logger.Infof("NODESIGNER %s; Signature: %s; message: %s", signer, signature, hex.EncodeToString(*message))
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
	return ed25519.Verify(publicKeyBytes, msg[:], signatureByte), nil
}

func VerifySignatureSECP(publicKeyBytes []byte, message []byte, signatureByte []byte) (bool, error) {

	pubKey, err := btcec.ParsePubKey(publicKeyBytes, btcec.S256())
	if err != nil {
		logger.WithFields(logrus.Fields{"key": string(publicKeyBytes)}).Infof("Unable to parse pubkey %v", err)
		return false, err
	}
	parsedSign, err := btcec.ParseSignature(signatureByte, btcec.S256())
	if err != nil {
		logger.WithFields(logrus.Fields{"signature": string(signatureByte)}).Infof("Unable to parse signature %v", err)
		return false, err
	}
	msg := Sha256(message)
	// v, _ := hex.DecodeString(fmt.Sprintf("%s%s",  parsedSign.R.Text(16), parsedSign.S.Text(16)))
	// sig := encoder.ToBase64Padded(v)
	print("HASSEHD ", parsedSign.R.Text(16), parsedSign.S.Text(16))
	return parsedSign.Verify(msg[:], pubKey), nil
}

func Bech32AddressFromPrivateKeyEDD(privateKey string) string {
	decoded, _ := hex.DecodeString(GetPublicKeyEDD(privateKey))
	return ToBech32Address(decoded, "ml")
}

/*
{
  "chain_id": "",
  "account_number": "0",
  "sequence": "0",
  "fee": {
    "gas": "0",
    "amount": []
  },
  "msgs": [
    {
      "type": "sign/MsgSignData",
      "value": {
        "signer": "cosmos1z7pux6petf6fvngdkap0cpyneztj5wwmlv7z9f",
        "data": "QXBwcm92ZSAweGU2NTJkMjhGODlBMjhhZGI4OWU2NzRhNmI1MTg1MkQwQzM0MUViZTkgZm9yIHRtbDogb2ZyUHZTR1pLNVV0S2NEVEhCQUNkOVB0dERqalJLbkhNUXFOaGpPSDJ5QT0="
      }
    }
  ],
  "memo": ""
}

*/

func ToBech32Address(publicKey []byte, prefix string) string {

	shaHash := Sha256(publicKey)
	ripemd160Hasher := ripemd160.New()
	ripemd160Hasher.Write(shaHash)
	publicHash := ripemd160Hasher.Sum(nil)

	// Convert to Bech32
	bech32Address, err := bech32.ConvertAndEncode(prefix, publicHash)
	if err != nil {
		logger.Fatal("Error converting to Bech32:", err)
	}

	return bech32Address
}

// 30440220 8ee622395fd93884b700406e9f297914b1934d31cece58992897901d6f1930210220296c724c79c5875930fad760f49073b99bf0514d43d2499c0982
// 8ee622395fd93884b700406e9f297914b1934d31cece58992897901d6f193021296c724c79c5875930fad760f49073b99bf0514d43d2499c0982 6169eb06fe82
func ToBtcecSignature(sigHex string) (*[]byte, error) {
	signature, err := hex.DecodeString(sigHex)
	if err != nil {
		return nil, err
	}
	R := sigHex[:64]
	logger.Infof("SIGNA %s", R)
	S := sigHex[64:]
	rByte, err := hex.DecodeString(R)
	if err != nil {
		return nil, err
	}
	sByte, err := hex.DecodeString(S)
	if err != nil {
		return nil, err
	}
	signature = append([]byte{byte(32)}, sByte...)
	signature = append([]byte{byte(0x02)}, signature...)
	signature = append(rByte, signature...)
	headerMagic := []byte{byte(0x30)}
	len := []byte{byte(68)}
	intMarker := []byte{byte(0x02)}
	signature = append([]byte{byte(32)}, signature...)
	signature = append(intMarker, signature...)
	signature = append(len, signature...)
	signature = append(headerMagic, signature...)
	return &signature, nil
}

func VerifySignatureAmino(signedData string, signature []byte, account string, pubKey []byte) (bool, error) {
	jsonData := `{"account_number":"0","chain_id":"","fee":{"amount":[],"gas":"0"},"memo":"","msgs":[{"type":"sign/MsgSignData","value":{"data":"%s","signer":"%s"}}],"sequence":"0"}`
	jsonData = fmt.Sprintf(jsonData, signedData, account)

	b := []byte(jsonData)
	sig, _ := ToBtcecSignature(hex.EncodeToString(signature))
	verified, err := VerifySignatureSECP(pubKey, b, *sig)
	if err != nil {
		logger.Error(err)
		return false, err
	}
	return verified, err

}
func HashMessageEth(message []byte) []byte {
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message))
	hashed := sha3.NewLegacyKeccak256()
	hashed.Write([]byte(prefix))
	hashed.Write(message)
	return hashed.Sum(nil)
}

func EthMessage(message []byte) []byte {
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message))
	byteArr := []byte(prefix)
	byteArr = append(byteArr, message...)
	return byteArr
}
