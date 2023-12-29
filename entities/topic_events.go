package entities

import (
	// "errors"

	"encoding/json"
	"fmt"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/utils"

	"github.com/ethereum/go-ethereum/common/hexutil"
)



type TopicEvent struct {
	Id   string    `json:"id"`
	Topic   Topic    `json:"t"`
	Validator  AddressString   `json:"v"`
	Timestamp   int       `json:"ts"`
	Signature   string    `json:"sig"`
	Action      utils.SubscriptionAction `json:"a"`
	Broadcast   bool      `json:"br"`
}

func (topic *TopicEvent) Key() string {
	return fmt.Sprintf("/%s/%s", topic.Validator, topic.Id)
}

func (topic *TopicEvent) TopicEventToJSON() []byte {
	m, e := json.Marshal(topic)
	if e != nil {
		logger.Errorf("Unable to parse subscription to []byte")
	}
	return m
}

func (topic *TopicEvent) Pack() []byte {
	b, _ := utils.MsgPackStruct(topic)
	return b
}


func TopicEventFromBytes(b []byte) (TopicEvent, error) {
	var sub TopicEvent
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &sub)
	return sub, err
}
func UnpackTopicEvent(b []byte) (SubscriptionEvent, error) {
	var sub SubscriptionEvent
	err := utils.MsgPackUnpackStruct(b, sub)
	return sub, err
}



func (topic *TopicEvent) Hash() string {
	return hexutil.Encode(utils.Hash(topic.ToString()))
}

func (topic *TopicEvent) ToString() string {
	values := []string{}
	values = append(values, fmt.Sprintf("%s", topic.Id))
	// values = append(values, fmt.Sprintf("%s", sub.ChannelName))
	values = append(values, fmt.Sprintf("%d", topic.Timestamp))
	values = append(values, fmt.Sprintf("%s", topic.Action))
	return strings.Join(values, ",")
}

