package query

import (
	"fmt"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/internal/chain"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
)

// EventCount         uint64 `json:"ec"`
// MessageCount       uint64 `json:"mc"`
// TopicCount         uint64 `json:"tc"`
// AuthorizationCount uint64 `json:"ac"`
// Cycle uint64 `json:"cy"`
// Volume uint64 `json:"vol"`
func IncrementBlockStat(blockNumber uint64, eventType *constants.EventType ) (model *models.BlockStat, created bool, err error) {
	tx := db.Db
	fieldName := "event_count"
	switch *eventType {
	case constants.CreateTopicEvent:
		fieldName = "topic_count"
	case constants.SendMessageEvent:
		fieldName = "message_count"
}

	// var data models.BlockStat
	// fieldName := "event_count"
	increment := 1;
	if eventType != nil {
		switch *eventType {
		case constants.CreateTopicEvent:
			fieldName = "topic_count"
		case constants.SendMessageEvent:
			fieldName = "message_count"
		default:
			fieldName = "message_count"
			increment = 0
			return nil, false, nil
		}
	 }
	// eErr := tx.Where(where).
	// 	First(&data).Error
	// if eErr != nil {
	// 	if eErr != gorm.ErrRecordNotFound {
	// 		return nil, false, eErr
	// 	}
	// 	eErr = tx.Create(&where).Error
	// 	if eErr != nil {
	// 		return nil, false, eErr
	// 	}
	// }
	epoch:= chain.GetEpoch(blockNumber)
	cycle:= chain.GetCycle(blockNumber)

	tx.Exec(fmt.Sprintf("UPDATE %s SET event_count = event_count + 1, %s = %s + %d, cycle = %d, epoch = %d WHERE block_number = ?", GetTableName(models.BlockStat{}), fieldName, fieldName, increment, cycle, epoch),  blockNumber)
	
	// data.EventCount = data.EventCount + 1
	// tx.Save(&data)
	

	return nil, tx.RowsAffected > 0, nil
}
