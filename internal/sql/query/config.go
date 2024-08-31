package query

import (
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
)

func GetConfig(key string) (*models.Config, error) {

	data := models.Config{}
	err := db.SqlDb.Where(&models.Config{Key: key}).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func SetConfig(key string, value string) (*models.Config, error) {

	data := models.Config{}
	err := db.SqlDb.Where(models.Config{Key: key}).Assign(models.Config{Value: value}).FirstOrCreate(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
