package entities

import (
	// "errors"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/utils"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var logger = &utils.Logger

type Topic struct {
	Id   string    `json:"id"`
	Name string    `json:"n"`
	SubscriberCount  uint64    `json:"sC"`
	Owner  string    `json:"own"`
	Timestamp   int       `json:"ts"`
	Signature   string    `json:"sig"`
	Broadcast   bool      `json:"br"`
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

func (topic *Topic) Pack() []byte {
	b, _ := utils.MsgPackStruct(topic)
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
	err := utils.MsgPackUnpackStruct(b, topic)
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


func (topic *Topic) Hash() string {
	return hexutil.Encode(utils.Hash(topic.ToString()))
}

func (topic *Topic) ToString() string {
	values := []string{}
	values = append(values, fmt.Sprintf("%s", topic.Id))
	values = append(values, fmt.Sprintf("%s", topic.Name))
	values = append(values, fmt.Sprintf("%d", topic.Timestamp))
	values = append(values, fmt.Sprintf("%d", topic.SubscriberCount))
	values = append(values, fmt.Sprintf("%s", topic.Owner))
	// values = append(values, fmt.Sprintf("%s", topic.Signature))
	return strings.Join(values, ",")
}


