package migration

import (
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	"gorm.io/gorm"
)



 func DropTopicIdColumnFromMessageState(db *gorm.DB) (err error) {
	if db.Migrator().HasColumn(&models.MessageState{}, "TopicId") {
		err = db.Migrator().DropColumn(&models.MessageState{}, "TopicId")
	}
	return err
 }


 func DropAttachmentsColumnFromMessageState(db *gorm.DB) (err error) {
	if db.Migrator().HasColumn(&models.MessageState{}, "Attachments") {
		err = db.Migrator().DropColumn(&models.MessageState{}, "Attachments")
	}
	return err
 }