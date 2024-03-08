package models

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"gorm.io/gorm"
)




type MessageState struct {
	entities.Message
	BaseModel
}

func (d *MessageState) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == ""  {
		hash, err := entities.GetId(*d)
		if err != nil {
			panic(err)
		}
		d.ID = hash
	}
	return nil
  }

  
type MessageEvent struct {
	entities.Event
	BaseModel
}