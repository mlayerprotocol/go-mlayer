package query

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
)

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
	if ePath.Model == entities.SubscriptionEventModel {
		var subData *models.SubscriptionEvent
		err := GetOne(models.SubscriptionEvent{
			Event: entities.Event{Hash: ePath.Hash},
		}, &subData)
		return err == nil
	}

	if ePath.Model == entities.TopicEventModel {
		var topData *models.TopicEvent
		err := GetOne(models.TopicEvent{
			Event: entities.Event{Hash: ePath.Hash},
		}, &topData)
		return err == nil
	}

	if ePath.Model == entities.AuthEventModel {
		var authData *models.AuthorizationEvent
		err := GetOne(models.AuthorizationEvent{
			Event: entities.Event{Hash: ePath.Hash},
		}, &authData)
		return err == nil
	}

	if ePath.Model == entities.MessageEventModel {
		var msgData *models.MessageEvent
		err := GetOne(models.MessageEvent{
			Event: entities.Event{Hash: ePath.Hash},
		}, &msgData)
		return err == nil
	}

	return false
}
