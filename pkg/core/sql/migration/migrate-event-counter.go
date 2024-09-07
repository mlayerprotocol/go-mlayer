package migration

import (
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	"gorm.io/gorm"
)


 func AddClaimedFieldToEventCount(db *gorm.DB) (err error) {
	err = db.Migrator().AddColumn(&models.EventCounter{}, "Claimed")
	return err
 }
