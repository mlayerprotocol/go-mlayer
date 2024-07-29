package query

// EventCount         uint64 `json:"ec"`
// MessageCount       uint64 `json:"mc"`
// TopicCount         uint64 `json:"tc"`
// AuthorizationCount uint64 `json:"ac"`
// Cycle uint64 `json:"cy"`
// Volume uint64 `json:"vol"`
// func IncrementBlockStat(blockNumber uint64, eventType *constants.EventType ) (model *models.BlockStat, created bool, err error) {
// 	tx := db.SqlDb
// 	fieldName := "event_count"
// 	switch *eventType {
// 	case constants.CreateTopicEvent:
// 		fieldName = "topic_count"
// 	case constants.SendMessageEvent:
// 		fieldName = "message_count"
// }

// 	// var data models.BlockStat
// 	// fieldName := "event_count"
// 	increment := 1;
// 	if eventType != nil {
// 		switch *eventType {
// 		case constants.CreateTopicEvent:
// 			fieldName = "topic_count"
// 		case constants.SendMessageEvent:
// 			fieldName = "message_count"
// 		default:
// 			fieldName = "message_count"
// 			increment = 0
// 			return nil, false, nil
// 		}
// 	 }
// 	// eErr := tx.Where(where).
// 	// 	First(&data).Error
// 	// if eErr != nil {
// 	// 	if eErr != gorm.ErrRecordNotFound {
// 	// 		return nil, false, eErr
// 	// 	}
// 	// 	eErr = tx.Create(&where).Error
// 	// 	if eErr != nil {
// 	// 		return nil, false, eErr
// 	// 	}
// 	// }
// 	epoch := chain.GetEpoch(blockNumber)
// 	cycle := chain.GetCycle(blockNumber)

// 	 v, err :=  chain.API.GetCurrentMessageCost()
// 	 if err != nil {
// 		logger.Error(err)
// 		return
// 	 }
// 	tx.Exec(fmt.Sprintf("UPDATE %s SET event_count = event_count + 1, %s = %s + %d, message_cost = %s, cycle = %d, epoch = %d WHERE block_number = ?", GetTableName(models.BlockStat{}), fieldName, fieldName, increment,v, cycle, epoch),  blockNumber)

// 	// data.EventCount = data.EventCount + 1
// 	// tx.Save(&data)

// 	return nil, tx.RowsAffected > 0, nil
// }
