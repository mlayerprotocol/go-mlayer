package query

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
)

// func GetManyWithEvent[T any, U any](filter T, event entities.Event, data *U) error {
// 	err := db.Db.Where(&filter).Joins().Find(data).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

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

func EventExist(ePath *entities.EventPath) bool {
	if ePath.Hash == "" {
		return true
	}
	event, _ := GetEventFromPath(ePath)
	return event != nil
}

func GetEventFromPath(ePath *entities.EventPath) (*entities.Event, error) {
	if ePath.Model == entities.SubscriptionEventModel {
		var data *models.SubscriptionEvent
		err := GetOne(models.SubscriptionEvent{
			Event: entities.Event{Hash: ePath.Hash},
		}, &data)
		if err != nil {
			return nil, err
		}
		return &data.Event, nil
	}

	if ePath.Model == entities.TopicEventModel {
		var data *models.TopicEvent
		err := GetOne(models.TopicEvent{
			Event: entities.Event{Hash: ePath.Hash},
		}, &data)
		if err != nil {
			return nil, err
		}
		return &data.Event, nil
	}

	if ePath.Model == entities.AuthEventModel {
		var data *models.AuthorizationEvent
		err := GetOne(models.AuthorizationEvent{
			Event: entities.Event{Hash: ePath.Hash},
		}, &data)
		if err != nil {
			return nil, err
		}
		return &data.Event, nil
	}

	if ePath.Model == entities.MessageEventModel {
		var data *models.MessageEvent
		err := GetOne(models.MessageEvent{
			Event: entities.Event{Hash: ePath.Hash},
		}, &data)
		if err != nil {
			return nil, err
		}
		return &data.Event, nil
	}

	return nil, nil
}
