package migration

import (
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	"gorm.io/gorm"
)


 func MigrateAuthIndex(db *gorm.DB) (err error) {
	// if db.Migrator().HasIndex(&models.AuthorizationState{}, "signature_data") {
		err = db.Migrator().DropIndex(&models.AuthorizationState{}, "signature_data")
	// }
	return err
 }
