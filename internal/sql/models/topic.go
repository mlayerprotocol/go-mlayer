package models

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"gorm.io/gorm"
)

type TopicState struct {
	gorm.Model
	entities.Topic
}

type TopicEvent struct {
	BaseModel
	entities.Event
	// TopicID     uint64
	// Topic		TopicState
}
