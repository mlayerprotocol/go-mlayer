package entities

import (
	// "errors"

	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
	"github.com/mlayerprotocol/go-mlayer/utils/constants"
	"github.com/mlayerprotocol/go-mlayer/utils/encoder"
)

var logger = &log.Logger

type Subscription struct {
	TopicId   string    `json:"chId"`
	Subscriber  string    `json:"s"`
	Timestamp   uint64       `json:"ts"`
	Action 		constants.EventType `json:"act"`
	Signature   string    `json:"sig"`
	Broadcasted   bool      `json:"br"`
}

func (sub *Subscription) Key() string {
	return fmt.Sprintf("/%s/%s", sub.Subscriber, sub.TopicId)
}

func (sub *Subscription) ToJSON() []byte {
	m, e := json.Marshal(sub)
	if e != nil {
		logger.Errorf("Unable to parse subscription to []byte")
	}
	return m
}

func (sub *Subscription) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(sub)
	return b
}



func SubscriptionFromBytes(b []byte) (Subscription, error) {
	var sub Subscription
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &sub)
	return sub, err
}
func UnpackSubscription(b []byte) (Subscription, error) {
	var sub Subscription
	err := encoder.MsgPackUnpackStruct(b, sub)
	return sub, err
}



func (sub *Subscription) Hash() string {
	b, err := sub.EncodeBytes()
	if(err != nil) {
		log.Logger.Errorf("Subscription Hashing error, %v", err)
	}
	return hexutil.Encode(crypto.Keccak256Hash(b))
}

// func (sub *Subscription) ToString() string {
// 	values := []string{}
// 	values = append(values, fmt.Sprintf("%s", sub.TopicId))
// 	// values = append(values, fmt.Sprintf("%s", sub.ChannelName))
// 	values = append(values, fmt.Sprintf("%d", sub.Timestamp))
// 	values = append(values, fmt.Sprintf("%d", sub.Action))
// 	return strings.Join(values, "")
// }

func (sub *Subscription) EncodeBytes()  ([]byte, error) {
	// var buffer bytes.Buffer
	// buffer.Write([]byte(sub.TopicId))
	// buffer.Write(encoder.NumberToByte(sub.Timestamp))
	// buffer.Write(encoder.NumberToByte(uint64(sub.Action)))
	
	return encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: sub.TopicId},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: uint64(sub.Action)},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: sub.Timestamp},
	)
}
