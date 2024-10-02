package migration

import (
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	"gorm.io/gorm"
)


 func AddClaimedFieldToEventCount(db *gorm.DB) (err error) {
	if !db.Migrator().HasTable(&models.EventCounter{}) {
		return nil
	}
	if !db.Migrator().HasColumn(&models.EventCounter{}, "Claimed") {
		err = db.Migrator().AddColumn(&models.EventCounter{}, "Claimed")
	}
	return err
 }
