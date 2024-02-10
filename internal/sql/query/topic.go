package query

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"gorm.io/gorm"
)

// Save topic state only when it doesnt exist
func UpdateTopicState(topic *entities.Topic, DB *gorm.DB) (*models.TopicState, error) {
	data := models.TopicState{
		// Privilege 	: auth.Priviledge,
		Topic: *topic,
}
	tx := DB
	if DB == nil {
		tx = db.Db.Begin()
	}
	err := tx.Where(models.TopicState{
		Topic: entities.Topic{Hash: topic.Hash,
			Account: topic.Account,},
			}).Assign(data).FirstOrCreate(&data).Error
	if err != nil {
		return nil, err
	}
	if DB == nil {
		tx.Commit()
	}
	return &data, nil
}