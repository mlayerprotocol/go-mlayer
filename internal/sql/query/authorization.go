package query

import (
	"encoding/hex"

	"github.com/jinzhu/copier"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"gorm.io/gorm"
)



func GetOneAuthorizationState(auth entities.Authorization) (*models.AuthorizationState, error) {

	data := models.AuthorizationState{}
	err := db.Db.Where(&models.AuthorizationState{
		Authorization: auth,
		}).First(&data).Error
	if err != nil {
		
		return nil, err
	}
	return &data, nil
}

func GetOneAuthorizationEvent(event entities.Event) (*models.AuthorizationEvent, error) {

	data := models.AuthorizationEvent{}
	err := db.Db.Where(&models.AuthorizationEvent{Event: event}).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetManyAuthorizationEvents(event entities.Event) (*models.AuthorizationEvent, error) {

	data := models.AuthorizationEvent{}
	err := db.Db.Where(&models.AuthorizationEvent{Event: event}).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func GetDependentEvents(event entities.Event) (*[]models.AuthorizationEvent, error) {

	data := []models.AuthorizationEvent{}
	err := db.Db.Where(
		&models.AuthorizationEvent{Event: entities.Event{PreviousEventHash: event.Hash}},
	).Or(&models.AuthorizationEvent{Event: entities.Event{AuthEventHash: event.Hash}},
	// ).Or("? LIKE ANY (associations)", fmt.Sprintf("%%%s%%", event.Hash)
	).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Save authorization only when it doesnt exist
func SaveAuthorizationState(auth *entities.Authorization, DB *gorm.DB) (*models.AuthorizationState, error) {
	data := models.AuthorizationState{
		// Privilege 	: auth.Priviledge,
		Authorization: *auth,
}
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

func SaveAuthorizationEvent(event *entities.Event, update bool, DB *gorm.DB) (model *models.AuthorizationEvent, created bool, err error) {
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
	}
	var result *gorm.DB
	if update {
		result = tx.Where(models.AuthorizationEvent{
			Event: entities.Event{Hash: event.Hash},
				}).Assign(data).FirstOrCreate(&data)
	} else {
		result = tx.Where(models.AuthorizationEvent{
			Event: entities.Event{PayloadHash: event.PayloadHash},
				}).FirstOrCreate(&data)
	}
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

func UpdateAuthorizationEvent(where entities.Event, updateFields entities.Event, DB *gorm.DB) (model *models.AuthorizationEvent,  err error) {
	tx := DB
	if DB == nil {
		tx = db.Db.Begin()
	}
	// dataByte, err := encoder.MsgPackStruct(event.Payload)
	if err != nil {
		return nil, err
	}
	
	result := tx.Where(models.AuthorizationEvent{
			Event: where,
				}).Updates(models.AuthorizationEvent{
					Event: updateFields,
						}).First(&model)
	
	if result.Error != nil {  
		tx.Rollback()
		logger.Errorf("SQL: %v", result.Error)
		return nil, result.Error
	}
	if DB == nil {
		tx.Commit()
	}
	return model,  nil
}


func SaveAuthorizationStateAndEvent(authEvent *entities.Event, tx *gorm.DB) (*models.AuthorizationState, *models.AuthorizationEvent,  error) {
	if tx == nil {
		tx = db.Db.Begin()
	}
	
	auth := (*authEvent).Payload.Data.(entities.Authorization)
	
	hash, err := auth.GetHash()
	auth.Hash = hex.EncodeToString(hash)
	auth.EventHash = authEvent.Hash
	event , created, err := SaveAuthorizationEvent(authEvent, false, tx)
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
		state, err := GetOneAuthorizationState(entities.Authorization{Account: auth.Account, Agent: auth.Agent})
		if err != nil {
			logger.Errorf("SQL: %v", err)
			tx.Rollback()
			return nil, nil, err
		}
		tx.Commit()
		return state, event, nil
	}
	
	
}
