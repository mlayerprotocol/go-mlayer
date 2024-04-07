package query

import (
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"gorm.io/gorm"
)

func IncrementBlockStat(where models.BlockStat) (model *models.BlockStat, created bool, err error) {
	tx := db.Db

	var data models.BlockStat
	logger.Infof("+++++000000::::: %d", where.BlockNumber)
	logger.Infof("111111::::: %s", where)
	eErr := tx.Where(where).
		First(&data).Error
	if eErr != nil {
		if eErr != gorm.ErrRecordNotFound {
			return nil, false, eErr
		}
		eErr = tx.Create(&where).Error
		if eErr != nil {
			return nil, false, eErr
		}
	}
	logger.Infof("222222::::: %s", where)
	data.Count = data.Count + 1
	tx.Save(&data)
	logger.Infof("333333::::: %s", where)

	return &data, tx.RowsAffected > 0, nil
}
