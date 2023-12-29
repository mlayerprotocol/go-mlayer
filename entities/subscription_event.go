package entities

import (
	// "errors"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strings"

	utils "github.com/mlayerprotocol/go-mlayer/utils"

	"github.com/ethereum/go-ethereum/common/hexutil"
)


type SubscriptionEvent struct {
	TopicId   string    `json:"chId"`
	Subscriber  string    `json:"s"`
	Timestamp   int       `json:"ts"`
	Signature   string    `json:"sig"`
	Action      utils.SubscriptionAction `json:"a"`
	Broadcast   bool      `json:"br"`
}

func (sub *SubscriptionEvent) Key() string {
	return fmt.Sprintf("/%s/%s", sub.Subscriber, sub.TopicId)
}

func (sub *SubscriptionEvent) ToJSON() []byte {
	m, e := json.Marshal(sub)
	if e != nil {
		logger.Errorf("Unable to parse subscription to []byte")
	}
	return m
}

func (sub *SubscriptionEvent) Pack() []byte {
	b, _ := utils.MsgPackStruct(sub)
	return b
}

func ToByte(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)

	fmt.Println(b)
	return b
}

func SubscriptionFromBytes(b []byte) (SubscriptionEvent, error) {
	var sub SubscriptionEvent
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &sub)
	return sub, err
}
func UnpackSubscription(b []byte) (SubscriptionEvent, error) {
	var sub SubscriptionEvent
	err := utils.MsgPackUnpackStruct(b, sub)
	return sub, err
}



func (sub *SubscriptionEvent) Hash() string {
	return hexutil.Encode(utils.Hash(sub.ToString()))
}

func (sub *SubscriptionEvent) ToString() string {
	values := []string{}
	values = append(values, fmt.Sprintf("%s", sub.TopicId))
	// values = append(values, fmt.Sprintf("%s", sub.ChannelName))
	values = append(values, fmt.Sprintf("%d", sub.Timestamp))
	values = append(values, fmt.Sprintf("%s", sub.Action))
	return strings.Join(values, ",")
}

// [130 163 70 111 111 163 102 111 111 163 66 111 111 34]
// [130 163 102 111 111 163 102 111 111 163 98 111 111 34]
// [130 163 102 111 111 163 102 111 111 163 98 111 111 34]
