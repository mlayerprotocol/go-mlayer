package models

import (
	"github.com/google/uuid"
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
	gorm.Model
	ID                    string `json:"id" gorm:"type:uuid;primaryKey"`
	BaseModel             `msgpack:",noinline"`
	entities.Subscription `msgpack:",noinline"`
	// Privilege 	constants.AuthorizationPrivilege  `json:"priv" gorm:"type:int"`
}

func (bm *SubscriptionState) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	if bm.ID == "" {
		bm.ID = uuid.NewString()
	}
	return
}
