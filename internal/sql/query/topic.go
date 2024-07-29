package query

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"gorm.io/gorm"
)

func GetTopicEvents() (*[]models.TopicEvent, error) {
	var topicEvents []models.TopicEvent
	order := &map[string]Order{"timestamp": OrderDec}
	err := GetMany(models.TopicEvent{
		Event: entities.Event{
			BlockNumber: 1,
		},
	}, &topicEvents, order)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &topicEvents, nil
}

func GetTopic(where models.TopicState) (*models.TopicState, error) {
	topicState := models.TopicState{}

	err := GetOne(&where, &topicState)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &topicState, nil

}

func GetTopicById(id string) (*models.TopicState, error) {
	topicState := models.TopicState{}

	err := GetOne(models.TopicState{
		Topic: entities.Topic{ID: id},
	}, &topicState)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &topicState, nil

}
func GetTopicByHash(hash string) (*models.TopicState, error) {
	topicState := models.TopicState{}

	err := GetOne(models.TopicState{
		Topic: entities.Topic{Hash: hash},
	}, &topicState)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &topicState, nil

}

func GetTopics(subTopic entities.Topic) (*[]models.TopicState, error) {
	var topicStates []models.TopicState
	order := &map[string]Order{"timestamp": OrderDec}
	err := GetMany(models.TopicState{
		Topic: subTopic,
	}, &topicStates, order)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &topicStates, nil
}

// Save topic state only when it doesnt exist
func UpdateTopicState(topic *entities.Topic, DB *gorm.DB) (*models.TopicState, error) {
	data := models.TopicState{
		// Privilege 	: auth.Priviledge,
		Topic: *topic,
}
	tx := DB
	if DB == nil {
		tx = db.SqlDb
	}
	err := tx.Where(models.TopicState{
		Topic: entities.Topic{Hash: topic.Hash,
			Account: topic.Account,},
			}).Assign(data).FirstOrCreate(&data).Error
	if err != nil {
		return nil, err
	}
	
	return &data, nil
}