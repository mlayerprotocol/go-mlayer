package entities

import (
	// "errors"

	"encoding/json"
	"fmt"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/utils"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type EventData interface {
	GetHash() string
}

type Event struct {
	Id string    `json:"id"`
	Hash   string    `json:"h"`
	Data  EventData    `json:"d"`
	Timestamp   int       `json:"ts"`
	Signature   string    `json:"sig"`
	Type      utils.EventType `json:"t"`
	Broadcast   bool      `json:"br"`
	IsValid   bool      `json:"val"`
	Parents   []string      `json:"p"`
}

func (e *Event) Key() string {
	return fmt.Sprintf("/%s/%s", e.Hash)
}

func (e *Event) ToJSON() []byte {
	m, err := json.Marshal(e)
	if err != nil {
		logger.Errorf("Unable to parse event to []byte")
	}
	return m
}

func (e *Event) Pack() []byte {
	b, _ := utils.MsgPackStruct(e)
	return b
}


func EventFromBytes(b []byte) (Event, error) {
	var e Event
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &e)
	return e, err
}
func UnpackEvent(b []byte) (Event, error) {
	var e Event
	err := utils.MsgPackUnpackStruct(b, e)
	return e, err
}



func (e *Event) GetHash() string {
	return hexutil.Encode(utils.Hash(e.ToString()))
}

func (e *Event) ToString() string {
	values := []string{}
	values = append(values, fmt.Sprintf("%s", e.Id))
	values = append(values, fmt.Sprintf("%s", e.Data.GetHash()))
	values = append(values, fmt.Sprintf("%s", e.Type))
	values = append(values, fmt.Sprintf("%d", e.Timestamp))
	values = append(values, fmt.Sprintf("%d", strings.Join(e.Parents, ",")))
	return strings.Join(values, ",")
}


