package migration

import (
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	"gorm.io/gorm"
)


 func DropOwnerColumnFromSubnetState(db *gorm.DB) (err error) {
	if db.Migrator().HasColumn(&models.SubnetState{}, "Owner") {
		err = db.Migrator().DropColumn(&models.SubnetState{}, "Owner")
	}
	return err
 }

 func DropAgentColumnFromSubnetState(db *gorm.DB) (err error) {
	if db.Migrator().HasColumn(&models.SubnetState{}, "Agent") {
		err = db.Migrator().DropColumn(&models.SubnetState{}, "Agent")
	}
	return err
 }
