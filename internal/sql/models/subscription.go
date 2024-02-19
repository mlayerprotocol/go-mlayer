package models

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"gorm.io/gorm"
)

type SubscriptionEvent struct {
	BaseModel `msgpack:",noinline"`
	// Event `msgpack:",noinline"`
	entities.Event `msgpack:",noinline"`
	Status         string `gorm:"index" json:"st"`
	//IsValid   bool `gorm:"default:false" json:"isVal"`
	// EventType int16 `json:"t"`
	// Payload entities.ClientPayload  `json:"pld" gorm:"serializer:json" msgpack:",noinline"`
}

type SubscriptionState struct {
	BaseModel             `msgpack:",noinline"`
	entities.Subscription `msgpack:",noinline"`
	// Privilege 	constants.AuthorizationPrivilege  `json:"priv" gorm:"type:int"`
}


func (d *SubscriptionState) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == ""  {
		hash, err := entities.GetId(*d)
		if err != nil {
			panic(err)
		}
		d.ID = hash
	}
	return nil
}
