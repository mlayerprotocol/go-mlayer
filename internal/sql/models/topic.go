package models

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"gorm.io/gorm"
)

type TopicState struct {
	entities.Topic
	BaseModel
}

func (d *TopicState) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == ""  {
		hash, err := entities.GetId(*d)
		if err != nil {
			panic(err)
		}
		d.ID = hash
	}
	return nil
  }

  
type TopicEvent struct {
	entities.Event
	BaseModel
	// TopicID     uint64
	// Topic		TopicState
}


