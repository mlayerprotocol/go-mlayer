package query

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"gorm.io/gorm"
)

// Save subNetwork state only when it doesnt exist
func UpdateSubNetworkState(subNetwork *entities.SubNetwork, DB *gorm.DB) (*models.SubNetworkState, error) {
	data := models.SubNetworkState{
		// Privilege 	: auth.Priviledge,
		SubNetwork: *subNetwork,
	}
	tx := DB
	if DB == nil {
		tx = db.Db.Begin()
	}
	err := tx.Where(models.SubNetworkState{
		SubNetwork: entities.SubNetwork{Hash: subNetwork.Hash,
			Account: subNetwork.Account},
	}).Assign(data).FirstOrCreate(&data).Error
	if err != nil {
		return nil, err
	}
	if DB == nil {
		tx.Commit()
	}
	return &data, nil
}
