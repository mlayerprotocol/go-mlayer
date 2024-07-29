package entities

import (
	"encoding/json"
	"math/big"
	"time"

	// "math"

	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto/schnorr"
)

type SubnetCount struct {
	Subnet     string `json:"sNet"`
	EventCount uint64 `json:"eC"`
}

func (msg *SubnetCount) EncodeBytes() ([]byte, error) {
	d, _ := utils.UuidToBytes(msg.Subnet)
	return encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: d[len(d)-6:]},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: msg.EventCount},
	)
}

type Proof struct {
	Signature json.RawMessage   `json:"sig"`
	Signers   []json.RawMessage `json:"signers"`
	Commitment   json.RawMessage `json:"comm"`
}

// RewardBatch
type RewardBatch struct {
	Id           string          `json:"id"`
	Index        uint64           `json:"idx"`
	Data         []SubnetCount   `json:"d"`
	DataBoundary [2]SubnetCount  `json:"dBound"`
	DataHash     json.RawMessage `json:"dH"`
	ChainId      configs.ChainId `json:"pre"`
	Closed       bool            `json:"cl"`
	Cycle        uint64          `json:"nh"`
	MessageCost  json.RawMessage `json:"cost"`
	TotalValue   json.RawMessage `json:"tv"`
	Proofs       Proof           `json:"proof"`
	Timestamp    uint64          `json:"ts"`
	// Signature    json.RawMessage `json:"sig"`
	// Signer       string          `json:"sign"`
	config       *configs.MainConfiguration  `msgpack:"_"`
	// sync.Mutex
}

// func (msg *Block) ToJSON() []byte {
// 	var buf bytes.Buffer
// 	enc := msgpack.NewEncoder(&buf)
// 	enc.SetCustomStructTag("json")
// 	enc.Encode(msg)
// 	return buf.Bytes()
// }

func (msg *RewardBatch) Append(subnetCount SubnetCount) {
	msg.Data = append(msg.Data, subnetCount)
	encoded, err := subnetCount.EncodeBytes()
	if err != nil {
		panic("Unable to encode SubnetCount data")
	}
	msg.DataHash = crypto.Keccak256Hash(append(msg.DataHash, encoded...))
	subnetTotal := big.NewInt(0).SetBytes(msg.MessageCost).Mul(big.NewInt(int64(subnetCount.EventCount)), big.NewInt(1))
	msg.TotalValue = big.NewInt(0).Add(big.NewInt(0).SetBytes(msg.TotalValue), subnetTotal).Bytes()
	length := len(msg.Data)
	if length == 1 {
		msg.DataBoundary[0] = msg.Data[0]
	}
	if len(msg.Data) >= 100 {
		msg.Closed = true
		msg.DataBoundary[1] = msg.Data[length-1]
	}
}
func (msg *RewardBatch) EncodeBytes() ([]byte, error) {
	return encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: msg.ChainId.Bytes()},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: msg.Cycle},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: msg.DataHash},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: msg.Id},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: msg.Index},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: msg.Timestamp},
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: msg.TotalValue},
	)
}
func (rb *RewardBatch) GetHash(chainId configs.ChainId) ([32]byte, error) {
	rb.ChainId = chainId
	b, err := rb.EncodeBytes()
	if err != nil {
		return [32]byte{}, err
	}
	return [32]byte(crypto.Keccak256Hash(b)), nil
}
// func (rb *RewardBatch) Sign(privateKey string) (error) {
// 	pubKey := crypto.GetPublicKeySECP(privateKey)
// 	rb.Timestamp = uint64(time.Now().UnixMilli())
// 	b, err := rb.EncodeBytes();
// 	if(err != nil) {
// 		return err
// 	}
// 	rb.Signature, _ = crypto.SignSECP(b, privateKey)
//     rb.Signer = pubKey
// 	return nil
// }

// func (rb *RewardBatch) IsValid() (bool) {
// 	// Important security update. Do not remove.
// 	// Prevents cross chain replay attack
// 	rb.Prefix = rb.config.AddressPrefix  // Important security update. Do not remove
// 	//
// 	// if math.Abs(float64(uint64(time.Now().UnixMilli()) - mp.Timestamp)) > constants.VALID_HANDSHAKE_SECONDS {
// 	// 	logger.WithFields(logrus.Fields{"data": mp}).Warnf("Hanshake Expired: %d", uint64(time.Now().UnixMilli()) - mp.Timestamp)
// 	// 	return false
// 	// }
// 	signer, err := hex.DecodeString(string(rb.Signer));
// 	if err != nil {
// 		logger.Error("Unable to decode signer")
// 		return false
// 	}
// 	data, err := rb.EncodeBytes()
// 	if err != nil {
// 		logger.Error("Unable to decode signer")
// 		return false
// 	}
// 	// signature, err := hex.DecodeString(string(mp.Signature));
// 	// if err != nil {
// 	// 	logger.Error(err)
// 	// 	return false
// 	// }
// 	isValid, err := crypto.VerifySignatureSECP(signer, data, rb.Signature)
// 	if err != nil {
// 		logger.Error(err)
// 		return false
// 	}
// 	if !isValid {
// 	//	logger.WithFields(logrus.Fields{"message": mp.Protocol, "signature": mp.Signature}).Warnf("Invalid signer %s", mp.Signer)
// 		return false
// 	}

// 	return true
// }

// func (msg *Block) ToString() string {
// 	values := []string{}
// 	values = append(values, fmt.Sprintf("%s", string(msg.BlockId)))
// 	values = append(values, fmt.Sprintf("%d", msg.Size))
// 	values = append(values, fmt.Sprintf("%d", msg.NodeHeight))
// 	values = append(values, fmt.Sprintf("%s", strconv.Itoa(msg.Timestamp)))
// 	values = append(values, fmt.Sprintf("%s", msg.Hash))
// 	return strings.Join(values, "")
// }

// func (msg *Block) Sign(privateKey string) Block {

// 	msg.Timestamp = int(time.Now().Unix())
// 	_, sig := Sign(msg.ToString(), privateKey)
// 	msg.Signature = sig
// 	return *msg
// }

func NewRewardBatch(config *configs.MainConfiguration, cycle uint64, index uint64, cycleMessageCost *big.Int) *RewardBatch {
	id := utils.RandomString(32)
	return &RewardBatch{Id: id,
		config:      config,
		Timestamp:   uint64(time.Now().UnixMilli()),
		Cycle:       uint64(cycle),
		MessageCost: cycleMessageCost.Bytes(),
		Index: index,
	}
}

func RewardBatchFromBytes(b []byte) (*RewardBatch, error) {
	var message RewardBatch
	err := json.Unmarshal(b, &message)
	return &message, err
}

func UnpackRewardBatch(b []byte) (*RewardBatch, error) {
	var message RewardBatch
	err := encoder.MsgPackUnpackStruct(b, &message)
	return &message, err
}

func (msg *RewardBatch) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(msg)
	return b
}


type SignatureRequestData struct {
	Commitment schnorr.EthAddress `json:"comm"`
	Challenge json.RawMessage `json:"chal"`
	AggPubKey json.RawMessage `json:"aggPub"`
	// Message [32]byte `json:"msg"`
}
func (sr *SignatureRequestData) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(sr)
	return b
}
func UnpackSignatureRequestData(b []byte) (*SignatureRequestData, error) {
	var message SignatureRequestData
	err := encoder.MsgPackUnpackStruct(b, &message)
	return &message, err
}