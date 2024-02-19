package models

import (
	"github.com/google/uuid"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"gorm.io/gorm"
)



type AuthorizationEvent struct {
	// Event `msgpack:",noinline"`
	entities.Event	 `msgpack:",noinline"`
	//IsValid   bool `gorm:"default:false" json:"isVal"`
	// EventType int16 `json:"t"`
	// Payload entities.ClientPayload  `json:"pld" gorm:"serializer:json" msgpack:",noinline"`
	BaseModel `msgpack:",noinline"`
}

type AuthorizationState struct {
	entities.Authorization `msgpack:",noinline"`
	BaseModel
}

func (d *AuthorizationState) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	hash, err := (*d).Authorization.GetHash()
	if err != nil {
       return err
    }
	u, err := uuid.FromBytes(hash)
	if err != nil {
      return err
    }
	if d.ID == ""  {
		d.ID = u.String()
	}
	return nil
  }

