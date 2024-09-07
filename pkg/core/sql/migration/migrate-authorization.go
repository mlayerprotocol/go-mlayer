package migration

import (
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	"gorm.io/gorm"
)


 func MigrateAuthIndex(db *gorm.DB) (err error) {
	err = db.Migrator().DropIndex(&models.AuthorizationState{}, "signature_data")
	return err
 }
