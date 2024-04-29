package models

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"gorm.io/gorm"
)

type SubNetworkState struct {
	entities.SubNetwork
	BaseModel
}

func (d *SubNetworkState) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == "" {
		hash, err := entities.GetId(*d)
		if err != nil {
			panic(err)
		}
		d.ID = hash
	}
	return nil
}

type SubNetworkEvent struct {
	entities.Event
	BaseModel
	// SubNetworkID     uint64
	// SubNetwork		SubNetworkState
}
