package entities

import (
	// "errors"

	"encoding/binary"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/mlayerprotocol/go-mlayer/internal/crypto"

	"github.com/mlayerprotocol/go-mlayer/utils/encoder"
)



type Topic struct {
	Id   string    `json:"id"`
	Ref string `json:"ref,omitempty"`
	Name string    `json:"n,omitempty" binding:"required"`
	Handle string    `json:"h,omitempty" binding:"required"`
	Description string    `json:"desc,omitempty"`
	SubscriberCount  uint64    `json:"sC,omitempty"`
	Owner  AddressString    `json:"own,omitempty" binding:"required"`
	Timestamp   int       `json:"ts,omitempty" binding:"required"`
	Public bool `json:"pub,omitempty"`
	Signature   string    `json:"sig,omitempty" binding:"required"`
	Broadcast   bool      `json:"br,omitempty"`
	Hash string `json:"hash,omitempty"`
}

func (topic *Topic) Key() string {
	return fmt.Sprintf("/%s/%s", topic.Owner, topic.Id)
}

func (topic *Topic) ToJSON() []byte {
	m, e := json.Marshal(topic)
	if e != nil {
		logger.Errorf("Unable to parse subscription to []byte")
	}
	return m
}

func (topic *Topic) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(topic)
	return b
}


func TopicToByte(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)

	fmt.Println(b)
	return b
}

func TopicFromBytes(b []byte) (Topic, error) {
	var topic Topic
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &topic)
	return topic, err
}
func UnpackTopic(b []byte) (Topic, error) {
	var topic Topic
	err := encoder.MsgPackUnpackStruct(b, topic)
	return topic, err
} 


func (p *Topic) CanSend(channel string, sender AddressString) bool {
	// check if user can send
	return true
}

func (p *Topic) IsMember(channel string, sender AddressString) bool {
	// check if user can send
	return true
}


func (topic Topic) GetHash() string {
	b, err := topic.EncodeBytes()
	if err !=nil {

	}
	return hexutil.Encode(crypto.Keccak256Hash(b))
}

func (topic Topic) ToString() string {
	values := []string{}
	values = append(values, fmt.Sprintf("%s", topic.Id))
	values = append(values, fmt.Sprintf("%s", topic.Name))
	values = append(values, fmt.Sprintf("%d", topic.Timestamp))
	values = append(values, fmt.Sprintf("%d", topic.SubscriberCount))
	values = append(values, fmt.Sprintf("%s", topic.Owner))
	values = append(values, fmt.Sprintf("%t", topic.Public))
	// values = append(values, fmt.Sprintf("%s", topic.Signature))
	return strings.Join(values, ",")
}

func (topic Topic) EncodeBytes()  ([]byte, error) {
	return encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: topic.Name},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: topic.SubscriberCount},
		encoder.EncoderParam{Type: encoder.AddressEncoderDataType, Value: topic.Owner},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: topic.Timestamp},
	)
}

type TopicClientPayload  struct {
	ClientPayload
	Data Topic `json:"d"`
}


