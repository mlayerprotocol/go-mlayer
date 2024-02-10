package models

import (
	"github.com/google/uuid"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"gorm.io/gorm"
)

type TopicState struct {
	gorm.Model
	ID   string    `json:"id" gorm:"type:uuid;primaryKey"`
	entities.Topic
}

func (bm *TopicState) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	if bm.ID == ""  {
		bm.ID = uuid.NewString()
	}
	return
  }

type TopicEvent struct {
	BaseModel
	entities.Event
	// TopicID     uint64
	// Topic		TopicState
}
