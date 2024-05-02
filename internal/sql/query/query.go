package query

import (
	"fmt"

	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	db "github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
	"gorm.io/gorm"
)

var logger = &log.Logger

func GetOne[T any, U any](filter T, data *U) error {
	err := db.Db.Where(&filter).Take(data).Error
	if err != nil {

		return err
	}
	return nil
}
func GetTx() *gorm.DB {
	return db.Db
}

func GetManyTx[T any](item T) *gorm.DB {
	return db.Db.Where(&item)
}

type Order string
const (
	OrderDec Order = "desc"
	OrderAsc Order = "asc"
)

func GetMany[T any, U any](item T, data *U, order *map[string]Order) error {
	tx := GetManyTx(item)
	if order != nil {
		logger.Infof("ORDER BY")
		for k := range(*order) {
			logger.Infof("%s %s",  k, (*order)[k])
			tx = tx.Order(fmt.Sprintf("%s %s", k, (*order)[k]))
		}
	}	
	err := tx.Find(data).Error
	if err != nil {
		return err
	}
	return nil
}

func GetWithIN[T any, U any, I any](item T, data *U, slice I) error {
	err := db.Db.Find(data, slice).Error

	if err != nil {
		return err
	}

	return nil
}

type Result struct {
	models.SubscriptionState
	Block uint
}

func GetSubscriptionsByBlock(subState entities.Subscription, from uint, to uint, block bool) ([]Result, error) {
	subStateTable := GetTableName(models.SubscriptionState{})
	subEventTable := GetTableName(models.SubscriptionEvent{})

	rows, err := db.Db.Model(&models.SubscriptionState{}).Select(fmt.Sprintf("%s.*, %s.block_number", subStateTable, subEventTable)).Where(models.SubscriptionState{Subscription: subState}).
		Joins(fmt.Sprintf("left join %s on  %s.event_hash = %s.hash", subEventTable, subStateTable, subEventTable)).
		Where(fmt.Sprintf("%s.block_number >= %d and %s.block_number < %d", subEventTable, from, subEventTable, to)).Rows()

	list := []Result{}
	for rows.Next() {
		rsl := Result{}
		rows.Scan(rsl)
		list = append(list, rsl)
	}
	if err != nil {
		return list, err
	}
	logger.Info("RSL ", list)
	return list, nil
}

func GetManyGroupBy[T any, U any](item T, data *U, gb string) error {
	err := db.Db.Where(&item).Select(fmt.Sprintf("topic_events.*, count(%s) as total", gb)).Group(gb).Find(data).Error
	if err != nil {
		return err
	}

	return nil
}

// func GetDependentEvents [T any] (event entities.Event) (*[]T, error) {

// 	var data []T
// 	err := db.Db.Where(
// 		&models.AuthorizationEvent{Event: entities.Event{PreviousEventHash: event.Hash}},
// 	).Or(&models.AuthorizationEvent{Event: entities.Event{AuthEventHash: event.Hash}},
// 	// ).Or("? LIKE ANY (associations)", fmt.Sprintf("%%%s%%", event.Hash)
// 	).Find(&data).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &data, nil
// }

func SaveRecord[Model any](where Model, data Model, update bool, DB *gorm.DB) (model *Model, created bool, err error) {
	tx := DB
	if DB == nil {
		tx = db.Db.Begin()
	}
	// dataByte, err := encoder.MsgPackStruct(event.Payload)
	if err != nil {
		return nil, false, err
	}
	// authPayload := entities.ClientPayload{
	// 	Data: entities.Authorization{},
	// }
	// copier.Copy(&authPayload, &event.Payload)
	// data := models.AuthorizationEvent{
	// 	Event: *event,
	// }
	var result *gorm.DB
	if update {
		result = tx.Where(where).Assign(data).FirstOrCreate(&data)
	} else {
		result = tx.Create(&data)
	}
	if result.Error != nil {
		tx.Rollback()

		return nil, false, result.Error
	}
	if DB == nil {
		tx.Commit()
	}
	return &data, result.RowsAffected > 0, nil
}

func SaveRecordWithMap[Model any](model *Model, where map[string]interface{}, data map[string]interface{},  update bool, DB *gorm.DB ) (created bool, err error) {
	tx := DB
	if DB == nil {
		tx = db.Db.Begin()
	}
	var result *gorm.DB
	if update {
		result = tx.Model(model).Where(where).Assign(data).FirstOrCreate(&data)
	} else {
	 result = tx.Model(model).Where(where).Create(data)
	}
	if result.Error != nil {
		tx.Rollback()

		return  false, result.Error
	}
	if DB == nil {
		tx.Commit()
	}
	return  result.RowsAffected > 0, nil
}

// func IncrementRecord[Model any](where Model, field string, DB *gorm.DB) (model *Model, created bool, err error) {
// 	tx := DB
// 	if DB == nil {
// 		tx = db.Db.Begin()
// 	}
// 	// dataByte, err := encoder.MsgPackStruct(event.Payload)
// 	if err != nil {
// 		return nil, false, err
// 	}
// 	// authPayload := entities.ClientPayload{
// 	// 	Data: entities.Authorization{},
// 	// }
// 	// copier.Copy(&authPayload, &event.Payload)
// 	// data := models.AuthorizationEvent{
// 	// 	Event: *event,
// 	// }
// 	var result *gorm.DB
// 	var data Model
// 	logger.Infof("111111::::: %s", where)
// 	result = tx.Where(where).Find(data)
// 	if result.Error != nil {
// 		tx.Rollback()

// 		return nil, false, result.Error
// 	}
// 	if result.RowsAffected == 0 {
// 		logger.Infof("111111 TTTT ::::: %s", where)
// 		result = tx.Where(where).FirstOrCreate(&data)
// 	} else {
// 		logger.Infof("111111 FFFF ::::: %s", where)
// 		result = tx.Where(where).Update(field, fmt.Sprintf("%s + 1", field))

// 	}
// 	// if update {
// 	// 	result = tx.Where(where).Assign(data).FirstOrCreate(&data)
// 	// } else {
// 	// 	result = tx.Where(where).FirstOrCreate(&data)
// 	// }
// 	logger.Infof("22222 ::::: %s", where)
// 	if result.Error != nil {
// 		tx.Rollback()

// 		return nil, false, result.Error
// 	}
// 	logger.Infof("33333 ::::: %s", where)
// 	if DB == nil {
// 		if tx.Commit().Error != nil {
// 			tx.Rollback()
// 			return nil, false, tx.Error
// 		}

// 	}
// 	logger.Infof("44444 ::::: %s", where)
// 	return &data, result.RowsAffected > 0, nil
// }

func GetTableName(table any) string {
	stmt := &gorm.Statement{DB: db.Db}
	stmt.Parse(table)
	return stmt.Schema.Table
}

func GetAccountSubscriptions(account string) {

	topicTableName := GetTableName(models.TopicState{})
	subscriptionTableName := GetTableName(models.SubscriptionState{})

	rows, err := db.Db.Table(subscriptionTableName).Where(fmt.Sprintf("%s.subscriber = \"%s\"", subscriptionTableName, account)).Joins(fmt.Sprintf("right join %s on %s.topic = %s.id", topicTableName, subscriptionTableName, topicTableName)).Rows()
	defer rows.Close()

	// logger.Info(rows.)

	if err != nil {
		logger.Info(err)
	}

	var subscriptions []map[string]string
	for rows.Next() {
		var subscription map[string]string
		db.Db.ScanRows(rows, &subscription)
		subscriptions = append(subscriptions, subscription)
	}
	logger.Infof("%s", subscriptions)

}
