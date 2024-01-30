package models

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
)


type DeleteEvent struct {
	entities.Event	 `msgpack:",noinline"`
	//IsValid   bool `gorm:"default:false" json:"isVal"`
	EventType int16 `json:"t"`
	// Payload datatypes.JSON  `json:"pld"`
}

