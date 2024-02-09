package models

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
)



type AuthorizationEvent struct {
	BaseModel `msgpack:",noinline"`
	// Event `msgpack:",noinline"`
	entities.Event	 `msgpack:",noinline"`
	//IsValid   bool `gorm:"default:false" json:"isVal"`
	// EventType int16 `json:"t"`
	// Payload entities.ClientPayload  `json:"pld" gorm:"serializer:json" msgpack:",noinline"`
}

type AuthorizationState struct {
	BaseModel `msgpack:",noinline"`
	entities.Authorization `msgpack:",noinline"`
	// Privilege 	constants.AuthorizationPrivilege  `json:"priv" gorm:"type:int"`
}
