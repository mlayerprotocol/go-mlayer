package models

import (
	"database/sql"

	"github.com/mlayerprotocol/go-mlayer/entities"
	"gorm.io/gorm"
)

type TopicState struct {
	gorm.Model
	ID          	string   `gorm:"type:uuid;primaryKey"`
	Name        string
	Handle        string `gorm:"unique"`
	Description     string
	SubscriberCount int64
	Owner string
	Timestamp   uint64 
	IsPublic sql.NullBool `gorm:"default:false"`
	Hash string
	Signature string
	AvailableBalance    float64 `gorm:"default:0"`
}

type TopicEvent struct {
	BaseModel
	entities.Event
	TopicID     uint64
	Topic		TopicState
}
