package entities

import (
	// "errors"

	"encoding/hex"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto/schnorr"
)


type RegisterationData struct {
	ChainId configs.ChainId  `json:"cId"`
	Timestamp uint64 `json:"ts"`
}

func (regData RegisterationData) EncodeBytes() ([]byte, error) {
	return encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: regData.ChainId.Bytes()},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: regData.Timestamp},
	)
}

func (regData *RegisterationData) Sign(privateKey string) ([]byte, schnorr.EthAddress, error) {
	if regData.Timestamp == 0 {
		regData.Timestamp = uint64(time.Now().UnixMilli())
	}
	privkBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, nil, err
	}
	_, p := btcec.PrivKeyFromBytes(btcec.S256(), privkBytes)
	logger.Infof("PUBKEY_X %d | %d", p.X, p.Y )
	signature, commitment, _, _ := schnorr.SignSingle(privkBytes, [32]byte(regData.GetHash()))
	return signature, commitment, err
}

func (regData *RegisterationData) GetHash() []byte {
	d, _ := regData.EncodeBytes()
	return crypto.Keccak256Hash(d)
}

