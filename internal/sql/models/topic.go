package models

import (
	"github.com/google/uuid"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"gorm.io/gorm"
)

type TopicState struct {
	ID   string    `json:"id" gorm:"type:uuid;primaryKey"`
	entities.Topic
}

func (d *TopicState) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	hash, err := (*d).Topic.GetHash()
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
type TopicEvent struct {
	BaseModel
	entities.Event
	// TopicID     uint64
	// Topic		TopicState
}
