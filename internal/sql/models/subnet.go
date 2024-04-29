package models

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"gorm.io/gorm"
)

type SubnetState struct {
	entities.Subnet
	BaseModel
}

func (d *SubnetState) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == "" {
		hash, err := entities.GetId(*d)
		if err != nil {
			panic(err)
		}
		d.ID = hash
	}
	return nil
}

type SubnetEvent struct {
	entities.Event
	BaseModel
	// SubnetID     uint64
	// Subnet		SubnetState
}
