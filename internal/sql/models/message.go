package models

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"gorm.io/gorm"
)


type MessageState struct {
	gorm.Model
	ID          	string   `gorm:"type:string;primaryKey"`
	Message       string
	Subject     string
	To string
	Signature string
	Timestamp   uint64 
}

type MessageEvent struct {
	BaseModel
	entities.Event
	MessageID        	uint64
	Message 			MessageState
}