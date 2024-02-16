package query

import (
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
	"gorm.io/gorm"
)

var logger = &log.Logger

func GetOne[T any, U any](filter T, data *U) error {
	err := db.Db.Where(&filter).First(data).Error
	if err != nil {

		return err
	}
	return nil
}

func GetMany[T any, U any](item T, data *U) error {
	err := db.Db.Where(&item).Find(data).Error
	if err != nil {
		return err
	}

	return nil
}

// func GetDependentEvents [T any] (event entities.Event) (*[]T, error) {

// 	var data []T
// 	err := db.Db.Where(
// 		&models.AuthorizationEvent{Event: entities.Event{PreviousEventHash: event.Hash}},
// 	).Or(&models.AuthorizationEvent{Event: entities.Event{AuthEventHash: event.Hash}},
// 	// ).Or("? LIKE ANY (associations)", fmt.Sprintf("%%%s%%", event.Hash)
// 	).Find(&data).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &data, nil
// }

func SaveRecord[Model any](where Model, data Model, update bool, DB *gorm.DB) (model *Model, created bool, err error) {
	tx := DB
	if DB == nil {
		tx = db.Db.Begin()
	}
	// dataByte, err := encoder.MsgPackStruct(event.Payload)
	if err != nil {
		return nil, false, err
	}
	// authPayload := entities.ClientPayload{
	// 	Data: entities.Authorization{},
	// }
	// copier.Copy(&authPayload, &event.Payload)
	// data := models.AuthorizationEvent{
	// 	Event: *event,
	// }
	var result *gorm.DB
	if update {
		result = tx.Where(where).Assign(data).FirstOrCreate(&data)
	} else {
		result = tx.Where(where).FirstOrCreate(&data)
	}
	if result.Error != nil {
		tx.Rollback()

		return nil, false, result.Error
	}
	if DB == nil {
		tx.Commit()
	}
	return &data, result.RowsAffected > 0, nil
}
