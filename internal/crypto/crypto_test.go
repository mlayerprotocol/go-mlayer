package crypto

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/mlayerprotocol/go-mlayer/common/encoder"
)

// func GetPublicKeyECC(privKey string) string {
// 	privateKey, err := crypto.HexToECDSA(privKey)
// 	if err != nil {
// 		logger.Fatalf("Invalid private key %o", err)
// 	}
// 	return crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
// }

// func GetPublicKeySECP(privKey string) string {
// 	privateKey, err := hex.DecodeString(privKey)
// 	if err != nil  {
// 		logger.Errorf("Invlaid node network key %v", err)
// 	}
// 	_, pub := btcec.PrivKeyFromBytes(btcec.S256(), privateKey)

// 	return hex.EncodeToString(pub.SerializeCompressed())
// }

// func GetPublicKeyEDD(privKey string) string {
// 	if len(privKey) != 128{
// 		logger.Fatal("Invalid private key length")
//     }
// 	return privKey[64:]
// }

// func PrivateKeyFromString(privKey string) (*ecdsa.PrivateKey, error) {
// 	privateKey, err := crypto.HexToECDSA(privKey)
// 	if err != nil {
// 		logger.Fatalf("Invalid private key %o", err)
// 		return nil, err
// 	}
// 	return privateKey, nil
// }

// func Keccak256Hash(message []byte) []byte {
// 	messageHash := crypto.Keccak256Hash(message)
// 	return messageHash.Bytes()
// }

// func SignECC(message []byte, privKey string) ([]byte, string) {
// 	privateKey, err := crypto.HexToECDSA(privKey)
// 	if err != nil {
// 		logger.Fatalf("Invalid private key %o",  err)
// 	}

// 	hash := Keccak256Hash(message)
// 	// logger.WithFields(logrus.Fields{"action": "crypto.Sign", "message": message}).Infof("Message hash: %s", hash.Hex())
// 	signature, err := crypto.Sign(hash, privateKey)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}
// 	// signer, err := crypto.Ecrecover(hash.Bytes(), signature[:len(signature)-1])
// 	return signature, hexutil.Encode(signature)
// }

// func SignEDD(message []byte, privKey string) ([]byte, string) {
// 	key, err := hex.DecodeString(privKey)
// 	if err != nil {
// 		logger.Fatalf("Invalid private key string %o", err)
// 	}
// 	pKey := ed25519.PrivateKey(key)
// 	hash := Sha256(message)
// 	// logger.WithFields(logrus.Fields{"action": "crypto.Sign", "message": message}).Infof("Message hash: %s", hash.Hex())
// 	signature := ed25519.Sign(pKey, hash)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}
// 	// signer, err := crypto.Ecrecover(hash.Bytes(), signature[:len(signature)-1])
// 	return signature, hex.EncodeToString(signature)
// }

// func SignSECP(message []byte, privKey string) ([]byte, string) {
// 	privateKeyByte, err := hex.DecodeString(privKey)
// 	if err != nil {
// 		logger.Fatalf("Invalid private key %o",  err)
// 	}
// 	privateKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privateKeyByte)
// 	signature, err := privateKey.Sign(Sha256(message)[:])
//     if err != nil {
//         logger.Fatal(err)
//     }
// 	b := signature.Serialize();
// 	return b, hex.EncodeToString(b)
// }

// func Sha256 (s []byte)  []byte {
//     h := cryptoSha256.New()
//     h.Write(s)
//     return h.Sum(nil)
// }

// func GetSignerECC(message *[]byte, signature *string) (string, error) {
// 	decoded, err := hexutil.Decode(*signature)
// 	if err != nil {
// 		logger.Debug(err)
// 		return "", err
// 	}
// 	hash := Keccak256Hash(*message)
// 	if decoded[crypto.RecoveryIDOffset] == 27 || decoded[crypto.RecoveryIDOffset] == 28 {
// 		decoded[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
// 	}

// 	signer, err := crypto.SigToPub(hash, decoded)
// 	if err != nil {
// 		return "", err
// 	}
// 	return crypto.PubkeyToAddress(*signer).Hex(), nil
// }

// func VerifySignatureECC(signer string, message *[]byte, signature string) bool {
// 	decodedSigner, err := GetSignerECC(message, &signature)
// 	if err != nil {
// 		return false
// 	}
// 	println("signer decoded signer %s %s : %T", decodedSigner, signer, (decodedSigner == signer))
// 	return strings.EqualFold(decodedSigner, signer)
// }

// func VerifySignatureEDD(signer string, message *[]byte, signature string) (bool, error) {
// 	signatureByte, err := hex.DecodeString(signature)
// 	if err != nil {
// 		logger.WithFields(logrus.Fields{"signature": signature}).Infof("Unable to decode signature %v", err)
// 		return false, err
// 	}
// 	publicKeyBytes, err := hex.DecodeString(signer)
// 	if err != nil {
// 		logger.WithFields(logrus.Fields{"signer": signer}).Infof("Unable to decode signer %v", err)
// 		return false, err
// 	}

// 	msg := Sha256(*message)
// 	return  ed25519.Verify(publicKeyBytes, msg[:], signatureByte), nil
// }

// func VerifySignatureSECP(signer string, message *[]byte, signature string) (bool, error) {
// 	signatureByte, err := hex.DecodeString(signature)
// 	if err != nil {
// 		logger.WithFields(logrus.Fields{"signature": signature}).Infof("Unable to decode signature %v", err)
// 		return false, err
// 	}
// 	publicKeyBytes, err := hex.DecodeString(signer)
// 	if err != nil {
// 		logger.WithFields(logrus.Fields{"signer": signer}).Infof("Unable to decode signer %v", err)
// 		return false, err
// 	}
// 	publicKey, err := btcec.ParsePubKey(publicKeyBytes, btcec.S256())
//     if err != nil {
// 		logger.WithFields(logrus.Fields{"key": string(publicKeyBytes)}).Infof("Unable to parse pubkey %v", err)
// 		return false, err
//     }
// 	parsedSign, err := btcec.ParseSignature(signatureByte, btcec.S256())
// 	if err != nil {
// 		logger.WithFields(logrus.Fields{"signature": string(signatureByte)}).Infof("Unable to parse signature %v", err)
// 		return false, err
// 	}
// 	msg := Sha256(*message)
// 	return parsedSign.Verify(msg[:], publicKey), nil
// }

// func Bech32AddressFromPrivateKeyEDD(privateKey string) string {
// 	return ToBech32Address(GetPublicKeyEDD(privateKey))
// }

// func ToBech32Address(publicKey string) string {
// 	b, err := hex.DecodeString(publicKey)
// 	if err != nil {
// 		logger.Fatal("Failed decoding public key", err)
// 	}
//     shaHash := Sha256(b)
// 	ripemd160Hasher := ripemd160.New()
//     ripemd160Hasher.Write(shaHash)
//     publicHash := ripemd160Hasher.Sum(nil)

//     // Convert to Bech32
//     bech32Address, err := bech32.ConvertAndEncode("ml:", publicHash)
//     if err != nil {
//         logger.Fatal("Error converting to Bech32:", err)
//     }

//     return bech32Address
// }


func TestGenerateAddress(t *testing.T) {
	account := "cosomos1kx4as0xqy9um0dwnxes39mrvtp0980szyxeweh"
	pubKey := "2MuHyTejCchvad6jcwsKhiJGK6csFl1QEZ/v/w4diCw="
	
	pubkey, err :=  base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		t.Fatalf("Unable to decode pubkey: %v", err )
	}
	hrp, _ :=  encoder.ExtractHRP(account)
	address  := ToBech32Address(pubkey, hrp)
	
	
    if account != address {
        t.Fatalf("Invalid address: %v", err )
    }
}

func TestVerifySignatureAmino(t *testing.T) {
	account := "cosmos1z7pux6petf6fvngdkap0cpyneztj5wwmlv7z9f"
	// pk := "2MuHyTejCchvad6jcwsKhiJGK6csFl1QEZ/v/w4diCw="
	// privKey := "bc3d5a5a6bb5024b1a96fccb677f065985d8e65d8054095eb6468244fb5ea4a9"
	pk := "AtLE+hi6ROU+ENXsJbauhDnz/K+WERg823eF3+Lwx6tz"
	msg:="Approve 0xe652d28F89A28adb89e674a6b51852D0C341Ebe9 for tml: ofrPvSGZK5UtKcDTHBACd9PttDjjRKnHMQqNhjOH2yA="
	sig:="juYiOV/ZOIS3AEBunyl5FLGTTTHOzliZKJeQHW8ZMCEpbHJMecWHWTD612D0kHO5m/BRTUPSSZwJgmFp6wb+gg=="
	plainSign, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {	
	t.Fatalf("Unable to decode signature: %v", err )
	}
	
	// computedPubK, _ := hex.DecodeString(GetPublicKeySECP(privKey))
	// print("PUBLIKEY ====>", encoder.ToBase64Padded(computedPubK) )
	pubkey, err :=  base64.StdEncoding.DecodeString(pk)
	if err != nil {
		t.Fatalf("Unable to decode pubkey: %v", err )
	}
	
	logger.Debugf("signature %s; publicKey %s", string(plainSign), hex.EncodeToString(pubkey))
	logger.Debug("msg: %", []byte(msg))
	
	valid, err := VerifySignatureAmino(encoder.ToBase64Padded([]byte(msg)), plainSign, account, pubkey)
    if !valid {
        t.Fatalf("Invalid signature signer: %v", err )
    }
}
//{"account_number":"0","chain_id":"","fee":{"amount":[],"gas":"0"},"memo":"","msgs":[{"type":"sign/MsgSignData","value":{"data":"aGVsbG93b3JsZA==","signer":"cosomos1z7pux6petf6fvngdkap0cpyneztj5wwm2maten"}}],"sequence":"0"}