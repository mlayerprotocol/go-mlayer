package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID string `gorm:"primaryKey" json:"id,omitempty"`
	
  }
  
  // Note: Gorm will fail if the function signature
  //  does not include `*gorm.DB` and `error`
  
  func (bm *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	bm.ID = uuid.NewString()
	return
  }

var Models = []interface{}{
	Config{},
	TopicState{},
	TopicEvent{},
	MessageState{},
	MessageEvent{},
	AuthorizationState{},
	AuthorizationEvent{},
}
