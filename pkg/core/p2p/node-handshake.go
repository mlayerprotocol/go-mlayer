package p2p

import (
	"encoding/hex"
	"encoding/json"
	"math"
	"time"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/sirupsen/logrus"
)

/*
*
NODE HANDSHAKE MESSAGE
*
*/


type NodeHandshake struct {
    Timestamp  uint64    `json:"ts"`
	Protocol string `json:"pro"`
    ChainId configs.ChainId `json:"pre"`
	NodeType   constants.NodeType   `json:"nT"`
	Salt      string `json:"salt"`
	Signature json.RawMessage        `json:"s"`
	Signer    string        `json:"sigr"`
	LastSyncedBlock  json.RawMessage `json:"lSy"`
	config 	*configs.MainConfiguration `json:"-" msgpack:"-"`
}

func (hs *NodeHandshake) ToJSON() []byte {
	h, _ := json.Marshal(hs)
	return h
}
func (hs *NodeHandshake) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(hs)
	return b
}

func (hsd NodeHandshake) EncodeBytes() ([]byte, error) {
    return encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: hsd.NodeType},
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: hsd.ChainId.Bytes()},
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: hsd.LastSyncedBlock},
        encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: hsd.Protocol},
        encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: hsd.Salt},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: hsd.Timestamp},
	)
}


func UnpackNodeHandshake(b []byte) (NodeHandshake, error) {
	var message NodeHandshake
	err := encoder.MsgPackUnpackStruct(b, &message)
	return message, err
}

func (hs *NodeHandshake) Init(jsonString string) error {
	er := json.Unmarshal([]byte(jsonString), &hs)
	return er
}

func (handshake *NodeHandshake) IsValid(chainId configs.ChainId) bool {
	// Important security update. Do not remove. 
	// Prevents cross chain replay attack
	handshake.ChainId = chainId // Important security update. Do not remove
	//
	if math.Abs(float64(uint64(time.Now().UnixMilli()) - handshake.Timestamp)) > constants.VALID_HANDSHAKE_SECONDS {
		logger.WithFields(logrus.Fields{"data": handshake}).Warnf("Node Handshake Expired: %d", uint64(time.Now().UnixMilli()) - handshake.Timestamp)
		return false
	}
	signer, err := hex.DecodeString(string(handshake.Signer));
	if err != nil {
		logger.Error("Unable to decode signer")
		return false
	}

	data, err := handshake.EncodeBytes()
	if err != nil {
		logger.Error("Unable to decode signer")
		return false
	}
	
	isValid, err := crypto.VerifySignatureSECP(signer, data, handshake.Signature)
	if err != nil {
		logger.Error(err)
		return false
	}
	logger.Debugf("Validating handshake data %v", isValid)
	if !isValid {
		logger.WithFields(logrus.Fields{"message": handshake.Protocol, "signature": handshake.Signature}).Warnf("Invalid signer %s", handshake.Signer)
		return false
	}
	return true
}
// func (handshake *NodeHandshake) hasValidStake(cfg *configs.MainConfiguration) bool {
// 	if handshake.NodeType == constants.ValidatorNodeType && cfg.Validator {
// 		return chain.HasValidStake(handshake.Signer, cfg)
// 	}
// 	return true
// }


func NodeHandshakeFromJSON(json string) (NodeHandshake, error) {
	data := NodeHandshake{}
	er := data.Init(json)
	return data, er
}

// func NodeHandshakeFromBytes(b []byte) NodeHandshake {
// 	var handshake NodeHandshake
// 	if err := json.Unmarshal(b, &handshake); err != nil {
// 		panic(err)
// 	}
// 	return handshake
// }

// func NodeHandshakeFromString(hs string) NodeHandshake {
// 	return NodeHandshakeFromBytes([]byte(hs))
// }

func NewNodeHandshake(config *configs.MainConfiguration, protocolId string, privateKey []byte, nodeType constants.NodeType) (*NodeHandshake, error) {
	pubKey := crypto.GetPublicKeySECP(privateKey)
	handshake := NodeHandshake{ config: config, Protocol: protocolId, Salt: utils.RandomString(6), ChainId: config.ChainId, NodeType: nodeType, Timestamp: uint64(time.Now().UnixMilli())}
	b, err := handshake.EncodeBytes();
	if(err != nil) {
		return nil, err
	}
	handshake.Signature, _ = crypto.SignSECP(b, privateKey)
    handshake.Signer = pubKey
	return &handshake, nil
}

