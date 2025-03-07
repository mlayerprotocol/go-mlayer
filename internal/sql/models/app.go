package models

import (
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"gorm.io/gorm"
)

type ApplicationState struct {
	entities.Application
	BaseModel
}

func (d *ApplicationState) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == "" {
		hash, err := entities.GetId(*d, d.ID)
		if err != nil {
			panic(err)
		}
		d.ID = hash
	}
	return nil
}

func (d ApplicationState) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(&d.Application)
	return b
}

type ApplicationEvent struct {
	entities.Event
	BaseModel
	// ApplicationID     uint64
	// Application		ApplicationState
}
