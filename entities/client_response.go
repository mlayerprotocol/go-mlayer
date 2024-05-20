package entities

import (
	"encoding/json"
)

type ResponsePayload interface {
	GetHash() ([]byte, error)
	ToString() string
	EncodeBytes() ([]byte, error)
}

type ResponseMeta struct {
	Page    uint   `json:"page,omitempty"`
	PerPage uint   `json:"perPage,omitempty"`
	Count   uint   `json:"count,omitempty"`
	Version string `json:"version,omitempty"`
}

type ClientResponse struct {
	Data  any    `json:"data"`
	Error string `json:"error,omitempty"`

	Meta ResponseMeta `json:"_meta"`
}

func (cr ClientResponse) ToMap() map[string]any {
	v, _ := json.Marshal(cr)
	var data map[string]any
	json.Unmarshal(v, &data)
	return data
}

func NewClientResponse(cr ClientResponse) ClientResponse {
	cr.Meta.Version = "1.0.1"
	return cr
}

 type ResponseInterval struct {
	FromTime uint64 `json:"fromT,omitempty"`
	ToTime uint64 `json:"toT,omitempty"`
	FromBlock uint64 `json:"fromB,omitempty"`
	ToBlock uint64 `json:"toB,omitempty"`
}
type TopicResponse struct {
	Updates []Event `json:"updates,omitempty"`
	Joins []Event `json:"joins,omitempty"`
	Leaves []Event `json:"leaves,omitempty"`
	Messages []Event `json:"msg,omitempty"`
}

type ConnectionState  uint8

const (
	OfflineState ConnectionState = 0
	OnlineState ConnectionState = 1
)
type Presence struct {
	Account DIDString `json:"acct"`
	MetaData json.RawMessage `json:"metaD"`
	ConnectionState ConnectionState `json:"connS"`
	ActiveAgo uint64 `json:"actA"`
}

 type SyncResponse struct {
	TimeFrame ResponseInterval `json:"time,omitempty"`
	Authorization []Authorization `json:"auths,omitempty"`
	Presence []Presence `json:"presence,omitempty"`
	Topics map[string]TopicResponse `json:"topics,omitempty"`
 }

	
