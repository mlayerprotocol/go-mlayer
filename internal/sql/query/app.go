package query

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"gorm.io/gorm"
)

// Save Application state only when it doesnt exist
func UpdateApplicationState(Application *entities.Application, DB *gorm.DB) (*models.ApplicationState, error) {
	data := models.ApplicationState{
		// Privilege 	: auth.Priviledge,
		Application: *Application,
	}
	tx := DB
	if DB == nil {
		tx = db.SqlDb
	}
	err := tx.Where(models.ApplicationState{
		Application: entities.Application{Hash: Application.Hash,
			Account: Application.Account},
	}).Assign(data).FirstOrCreate(&data).Error
	if err != nil {
		return nil, err
	}
	// if DB == nil {
	// 	tx.Commit()
	// }
	return &data, nil
}

func GetApplicationById(id string) (*models.ApplicationState, error) {
	state := models.ApplicationState{}

	err := GetOne(models.ApplicationState{
		Application: entities.Application{ID: id},
	}, &state)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &state, nil

}