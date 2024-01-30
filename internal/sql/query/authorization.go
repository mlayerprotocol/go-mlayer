package query

import (
	"encoding/hex"

	"github.com/jinzhu/copier"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
	"gorm.io/gorm"
)

var logger = &log.Logger;

func GetAuthorizationState(grantor string, agent string) (*models.AuthorizationState, error) {

	data := models.AuthorizationState{}
	err := db.Db.Where(&models.AuthorizationState{
		Authorization: entities.Authorization{Grantor: grantor, Agent: agent},
		}).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Save authorization only when it doesnt exist
func SaveAuthorizationState(auth *entities.Authorization, DB *gorm.DB) (*models.AuthorizationState, error) {
	data := models.AuthorizationState{
		Privilege 	: uint8(auth.Priviledge),
		Authorization: *auth,
}
logger.Infof("Auth Event Id========> %s", auth.AuthorizationEventID)
	tx := DB
	if DB == nil {
		tx = db.Db.Begin()
	}
	err := tx.Where(models.AuthorizationState{
		Authorization: entities.Authorization{Grantor: auth.Grantor,
			Agent: auth.Agent,},
			
			}).Assign(data).FirstOrCreate(&data).Error
	if err != nil {
		return nil, err
	}
	if DB == nil {
		tx.Commit()
	}
	return &data, nil
}

func SaveAuthorizationEvent(event *entities.Event, DB *gorm.DB) (model *models.AuthorizationEvent, created bool, err error) {
	tx := DB
	if DB == nil {
		tx = db.Db.Begin()
	}
	// dataByte, err := encoder.MsgPackStruct(event.Payload)
	if err != nil {
		return nil, false, err
	}
	authPayload := entities.ClientPayload{
		Data: entities.Authorization{},
	}
	copier.Copy(&authPayload, &event.Payload)
	data := models.AuthorizationEvent{
		Event: *event,
		Payload: authPayload,
	}
	result := tx.Where(models.AuthorizationEvent{
			Event: entities.Event{Hash: event.Hash},
				}).FirstOrCreate(&data)
	
	if result.Error != nil {  
		tx.Rollback()
		logger.Errorf("SQL: %v", result.Error)
		return nil, false, result.Error
	}
	if DB == nil {
		tx.Commit()
	}
	return &data, result.RowsAffected > 0,  nil
}


func SaveAuthorizationStateAndEvent(authEvent *entities.Event, tx *gorm.DB) (*models.AuthorizationState, *models.AuthorizationEvent,  error) {
	if tx == nil {
		tx = db.Db.Begin()
	}
	
	auth := (*authEvent).Payload.Data.(entities.Authorization)
	
	hash, err := auth.GetHash()
	auth.Hash = hex.EncodeToString(hash)
	auth.EventHash = authEvent.Hash
	event , created, err := SaveAuthorizationEvent(authEvent, tx)
	if err != nil {
		logger.Errorf("SQL: %v", err)
		return nil, nil, err
	}
	auth.AuthorizationEventID = event.Event.ID
	if created {
		state, err := SaveAuthorizationState(&auth, tx)
		if err != nil {
			logger.Errorf("SQL: %v", err)
			tx.Rollback()
			return nil, nil, err
		}
		tx.Commit()
		return state, event, nil
	} else {
		state, err := GetAuthorizationState(auth.Grantor, auth.Agent)
		if err != nil {
			logger.Errorf("SQL: %v", err)
			tx.Rollback()
			return nil, nil, err
		}
		tx.Commit()
		return state, event, nil
	}
	
	
}
