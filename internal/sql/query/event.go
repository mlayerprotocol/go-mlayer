package query

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
)


func GetManyWithEvent[T any, U any](filter T, event entities.Event, data *U) error {
	err := db.Db.Where(&filter).Joins().Find(data).Error
	if err != nil {
		return err
	}

	return nil
}

// func GetEvent(grantor string, agent string) (*models.Config, error) {

// 	data := models.Config{}
// 	err := db.Db.Where(&models.AuthorizationState{Grantor: grantor, Agent: agent}).First(&data).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &data, nil
// }

// func SaveEvent(event entities.Event, data models.Event) (*models.Event, error) {

// 	tx := db.Db.Begin()
// 	err := tx.Where(models.Event{
// 			Hash: event.Hash,
// 			}).Assign(models.Event{
// 				Parents   : event.Parents,
// 				Synced : sql.NullBool{Valid: true, Bool: event.Synced},
// 				Hash: event.Hash,
// 				StateHash:  event.StateHash,
// 				EventType: int16(event.EventType),
// 				Payload: event.MsgPack(),
// 				Validator: event.Validator,
// 				Timestamp : event.Timestamp,
// 				Signature : event.Signature,
// 				// IsValid :  event.IsValid
// 				 }).FirstOrCreate(&data).Error
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}
// 	tx.Commit()
// 	return &data, nil
// }
