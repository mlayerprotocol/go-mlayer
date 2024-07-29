package p2p

import (
	"math"
	"time"

	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/sirupsen/logrus"
)

/*
*
NODE HANDSHAKE MESSAGE
*
*/


type NodeMultiAddressData struct {
    Addresses []string `json:"addr"`
	Timestamp uint64 `json:"ts"`
	ChainId configs.ChainId `json:"pre"`
	Signer string `json:"signr"`
	Signature string `json:"sig"`
	config *configs.MainConfiguration `json:"-" msgpack:"-"`
	
}

func (hs *NodeMultiAddressData) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(hs)
	return b
}

func (n NodeMultiAddressData) EncodeBytes() ([]byte, error) {
	data := []byte{}
	for _, val := range n.Addresses {
		b, _ := encoder.EncodeBytes(encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: val})
		data = append(data, b...)
	}
    return encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: data},
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: n.ChainId.Bytes()},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: n.Timestamp},
	)
}


func UnpackNodeMultiAddressData(b []byte) ( NodeMultiAddressData, error) {
	var message  NodeMultiAddressData
	err := encoder.MsgPackUnpackStruct(b, &message)
	return message, err
}


func (nma * NodeMultiAddressData) IsValid(prefix configs.ChainId) bool {
	// Important security update. Do not remove. 
	// Prevents cross chain replay attack
	nma.ChainId = prefix  // Important security update. Do not remove
	//
	if math.Abs(float64(uint64(time.Now().UnixMilli()) - nma.Timestamp)) > float64(4 * time.Hour.Milliseconds()) {
		logger.WithFields(logrus.Fields{"data": nma}).Warnf("Hanshake Expired: %d", uint64(time.Now().UnixMilli()) - nma.Timestamp)
		return false
	}
	// signer, err := hex.DecodeString(string(nma.Signer));
	// if err != nil {
	// 	logger.Error("Unable to decode signer")
	// 	return false
	// }
	
	data, err := nma.EncodeBytes()
	if err != nil {
		logger.Error("Unable to decode signer")
		return false
	}
	// signature, err := hex.DecodeString(nma.Signature);
	// if err != nil {
	// 	logger.Error(err)
	// 	return false
	// }
	isValid, err := crypto.VerifySignatureEDD(nma.Signer, &data, nma.Signature)
	if err != nil {
		logger.Error(err)
		return false
	}
	
	if !isValid {
		logger.WithFields(logrus.Fields{"addresses": nma.Addresses, "signature": nma.Signature}).Warnf("Invalid signer %s", nma.Signer)
		return false
	}
	return true
}


func NewNodeMultiAddressData(config *configs.MainConfiguration, privateKey []byte, addresses []string) (*NodeMultiAddressData, error) {
	//pubKey := crypto.GetPublicKeySECP(privateKey)
	nma := NodeMultiAddressData{config: config, ChainId: config.ChainId, Addresses: addresses,   Timestamp: uint64(time.Now().UnixMilli())}
	b, err := nma.EncodeBytes();
	if(err != nil) {
		return nil, err
	}
	_, signature := crypto.SignEDD(b, config.PrivateKeyBytes)
    nma.Signature = signature
    nma.Signer = config.PublicKey
	return &nma, nil
}

