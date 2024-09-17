package migration

import (
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	"gorm.io/gorm"
)



 func DropTopicIdColumnFromMessageState(db *gorm.DB) (err error) {
	return  db.Migrator().DropColumn(&models.MessageState{}, "TopicId")
 }


 func DropAttachmentsColumnFromMessageState(db *gorm.DB) (err error) {
	return  db.Migrator().DropColumn(&models.MessageState{}, "Attachments")
 }