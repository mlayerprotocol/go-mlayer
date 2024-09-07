package query

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"gorm.io/gorm"
)

// Save Subnet state only when it doesnt exist
func UpdateSubnetState(Subnet *entities.Subnet, DB *gorm.DB) (*models.SubnetState, error) {
	data := models.SubnetState{
		// Privilege 	: auth.Priviledge,
		Subnet: *Subnet,
	}
	tx := DB
	if DB == nil {
		tx = db.SqlDb
	}
	err := tx.Where(models.SubnetState{
		Subnet: entities.Subnet{Hash: Subnet.Hash,
			Account: Subnet.Account},
	}).Assign(data).FirstOrCreate(&data).Error
	if err != nil {
		return nil, err
	}
	// if DB == nil {
	// 	tx.Commit()
	// }
	return &data, nil
}
