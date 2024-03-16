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
	Page uint `json:"page,omitempty"`
	PerPage uint `json:"perPage,omitempty"`
	Count uint `json:"count,omitempty"`
	Version string		`json:"version,omitempty"`
}

type ClientResponse struct {
	Data any  `json:"data"`
	Error string		`json:"error,omitempty"`
	
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

