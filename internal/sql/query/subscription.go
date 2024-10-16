package query

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"gorm.io/gorm"
)

// Save Subnet state only when it doesnt exist
func GetSubscriptionStateBySubscriber(subnet string, topic string, subscribers []entities.DIDString, DB *gorm.DB) (*[]models.SubscriptionState, error) {
	data := []models.SubscriptionState{}
	tx := DB
	if DB == nil {
		tx = sql.SqlDb
	}
	subsc := []entities.DIDString{}
	for _, sub := range subscribers {
		if sub == "" {
			continue
		}
		subsc = append(subsc, sub)
	}
	err := tx.Where(models.SubscriptionState{
		Subscription: entities.Subscription{ Subnet: subnet, Topic: topic },
	}).Where("subscriber IN ?", subsc).Find(&data).Error
	if err != nil {
		logger.Debugf("ERROR::: %v", err)
		return nil, err
	}
	// if DB == nil {
	// 	tx.Commit()
	// }
	return &data, nil
}
