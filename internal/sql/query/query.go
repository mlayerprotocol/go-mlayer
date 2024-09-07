package query

import (
	"fmt"
	"os"

	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var logger = &log.Logger

func GetOne[T any, U any](filter T, data U) error {
	if err := utils.CheckEmpty(filter); err != nil {
		return err
	}
	// logger.Debugf("QUERY %v, value: %v, %v", filter, data)
	err := sql.SqlDb.Where(filter).Take(data).Error
	if err != nil {

		return err
	}
	return nil
}

func GetOneWithOr[T any, U any](where T, or T, data U) error {
	if err := utils.CheckEmpty(where); err != nil {
		return err
	}
	if err := utils.CheckEmpty(or); err != nil {
		return err
	}
	err := sql.SqlDb.Where(where).Or(or).Take(data).Error
	if err != nil {

		return err
	}
	return nil
}

func GetTx() *gorm.DB {
	return sql.SqlDb
}

func GetManyTx[T any](item T) *gorm.DB {
	return sql.SqlDb.Where(&item)
}

type Order string

const (
	OrderDec Order = "desc"
	OrderAsc Order = "asc"
)

func GetMany[T any, U any](item T, data *U, order *map[string]Order) error {
	tx := GetManyTx(item)
	if order != nil {
		logger.Debugf("ORDER BY")
		for k := range *order {
			logger.Debugf("%s %s", k, (*order)[k])
			tx = tx.Order(fmt.Sprintf("%s %s", k, (*order)[k]))
		}
	}
	err := tx.Find(data).Error
	if err != nil {
		return err
	}
	return nil
}
func GetManyWithLimit[T any, U any](item T, data *U, order *map[string]Order, limit int, offset int) error {
	tx := GetManyTx(item)
	if order != nil {
		logger.Debugf("ORDER BY")
		for k := range *order {
			logger.Debugf("%s %s", k, (*order)[k])
			tx = tx.Order(fmt.Sprintf("%s %s", k, (*order)[k]))
		}
	}
	if limit > 0 {
		tx.Limit(int(limit))
	}
	if offset > 0 {
		tx.Offset(offset)
	}
	err := tx.Find(data).Error
	if err != nil {
		return err
	}
	return nil
}


func GetWithIN[T any, U any, I any](item T, data *U, slice I) error {
	err := sql.SqlDb.Find(data, slice).Error

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

	rows, err := sql.SqlDb.Model(&models.SubscriptionState{}).Select(fmt.Sprintf("%s.*, %s.block_number", subStateTable, subEventTable)).Where(models.SubscriptionState{Subscription: subState}).
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
	logger.Debug("RSL ", list)
	return list, nil
}

func GetManyGroupBy[T any, U any](item T, data *U, gb string) error {
	err := sql.SqlDb.Where(&item).Select(fmt.Sprintf("topic_events.*, count(%s) as total", gb)).Group(gb).Find(data).Error
	if err != nil {
		return err
	}

	return nil
}

// func GetDependentEvents [T any] (event entities.Event) (*[]T, error) {

// 	var data []T
// 	err := sql.SqlDb.Where(
// 		&models.AuthorizationEvent{Event: entities.Event{PreviousEventHash: event.Hash}},
// 	).Or(&models.AuthorizationEvent{Event: entities.Event{AuthEventHash: event.Hash}},
// 	// ).Or("? LIKE ANY (associations)", fmt.Sprintf("%%%s%%", event.Hash)
// 	).Find(&data).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &data, nil
// }

func GetOneState[T any](filter any, data *T) error {
		var model any
			switch val := filter.(type) {
				case entities.Subnet:
					logger.Debug(val)
					model = models.SubnetState{Subnet: val}
				case entities.Authorization:
					model = models.AuthorizationState{Authorization: val}
				case entities.Topic:
					model = models.TopicState{Topic: val}
				case entities.Subscription:
					model = models.SubscriptionState{Subscription: val}
				case entities.Message:
					model = models.MessageState{Message: val}
			}
			// logger.Debugf("QUERY %v, value: %v, %v", filter, data)
	err := sql.SqlDb.Where(model).Take(data).Error
	if err != nil {

		return err
	}
	return nil
}



func SaveRecord[Model any](where Model,  createData *Model, updateData *Model, DB *gorm.DB) (model *Model, created bool, err error) {
	tx := DB
	if DB == nil {
		tx = sql.SqlDb
	}
	// dataByte, err := encoder.MsgPackStruct(event.Payload)
	// if err != nil {
	// 	return nil, false, err
	// }
	// authPayload := entities.ClientPayload{
	// 	Data: entities.Authorization{},
	// }
	// copier.Copy(&authPayload, &event.Payload)
	// data := models.AuthorizationEvent{
	// 	Event: *event,
	// }
	if err := utils.CheckEmpty(where); err != nil {
		return model, false, err
	}


	var result *gorm.DB
	// if utils.CheckEmpty(updateData) != nil {
		if err := utils.CheckEmpty(updateData); err != nil {
			result = tx.Where(where).FirstOrCreate(createData)
		} else {
			result = tx.Where(where).Assign(updateData).FirstOrCreate(createData)
			
		}
	// } else {
	// 	result = tx.Where(where).FirstOrCreate(createData)
	// }
	// if result.Error != nil {
	// 	if DB == nil {
	// 		tx.Rollback()
	// 	}
	// 	return model, false, result.Error
	// }
	// if DB == nil {
	// 	tx.Commit()
	// }
	return createData, result.RowsAffected > 0, nil
}

func SaveRecordWithMap[Model any](model *Model, where map[string]interface{}, data map[string]interface{}, update bool, DB *gorm.DB) (created bool, err error) {
	tx := DB
	if DB == nil {
		tx = sql.SqlDb
	}
	var result *gorm.DB
	if update {
		result = tx.Model(model).Where(where).Assign(data).FirstOrCreate(&data)
	} else {
		result = tx.Model(model).Where(where).Create(data)
	}
	if result.Error != nil {
		// tx.Rollback()

		return false, result.Error
	}
	// if DB == nil {
	// 	tx.Commit()
	// }
	return result.RowsAffected > 0, nil
}

func GetDependentEvents(event *entities.Event) (*[]entities.Event, error) {

	data := []entities.Event{}
	// err := db.SqlDb.Where(
	// 	&models.AuthorizationEvent{Event: entities.Event{PreviousEventHash: *entities.NewEventPath(entities.AuthModel, event.Hash)}},
	// ).Or(&models.AuthorizationEvent{Event: entities.Event{AuthEventHash: *entities.NewEventPath(entities.AuthModel, event.Hash)}},
	// // ).Or("? LIKE ANY (associations)", fmt.Sprintf("%%%s%%", event.Hash)
	// ).Find(&data).Error
	// if err != nil {
	// 	return nil, err
	// }
	prevEvent, _ := GetEventFromPath(&(event.PreviousEventHash))
	if prevEvent != nil {
		data = append(data, *prevEvent)
	}
	authEvent, _ := GetEventFromPath(&(event.AuthEventHash))
	if authEvent != nil {
		data = append(data, *authEvent)
	}
	return &data, nil
}

// func IncrementRecord[Model any](where Model, field string, DB *gorm.DB) (model *Model, created bool, err error) {
// 	tx := DB
// 	if DB == nil {
// 		tx = sql.SqlDb.Begin()
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
// 	logger.Debugf("111111::::: %s", where)
// 	result = tx.Where(where).Find(data)
// 	if result.Error != nil {
// 		tx.Rollback()

// 		return nil, false, result.Error
// 	}
// 	if result.RowsAffected == 0 {
// 		logger.Debugf("111111 TTTT ::::: %s", where)
// 		result = tx.Where(where).FirstOrCreate(&data)
// 	} else {
// 		logger.Debugf("111111 FFFF ::::: %s", where)
// 		result = tx.Where(where).Update(field, fmt.Sprintf("%s + 1", field))

// 	}
// 	// if update {
// 	// 	result = tx.Where(where).Assign(data).FirstOrCreate(&data)
// 	// } else {
// 	// 	result = tx.Where(where).FirstOrCreate(&data)
// 	// }
// 	logger.Debugf("22222 ::::: %s", where)
// 	if result.Error != nil {
// 		tx.Rollback()

// 		return nil, false, result.Error
// 	}
// 	logger.Debugf("33333 ::::: %s", where)
// 	if DB == nil {
// 		if tx.Commit().Error != nil {
// 			tx.Rollback()
// 			return nil, false, tx.Error
// 		}

// 	}
// 	logger.Debugf("44444 ::::: %s", where)
// 	return &data, result.RowsAffected > 0, nil
// }

func GetTableName(table any) string {
	stmt := &gorm.Statement{DB: sql.SqlDb}
	stmt.Parse(table)
	return stmt.Schema.Table
}

func GetAccountSubscriptions(account string) {

	topicTableName := GetTableName(models.TopicState{})
	subscriptionTableName := GetTableName(models.SubscriptionState{})

	rows, err := sql.SqlDb.Table(subscriptionTableName).Where(fmt.Sprintf("%s.subscriber = \"%s\"", subscriptionTableName, account)).Joins(fmt.Sprintf("right join %s on %s.topic = %s.id", topicTableName, subscriptionTableName, topicTableName)).Rows()
	defer rows.Close()

	// logger.Debug(rows.)

	if err != nil {
		logger.Debug(err)
	}

	var subscriptions []map[string]string
	for rows.Next() {
		var subscription map[string]string
		sql.SqlDb.ScanRows(rows, &subscription)
		subscriptions = append(subscriptions, subscription)
	}
	logger.Debugf("%s", subscriptions)

}

func GenerateImportScript[T any](db *gorm.DB, model T, where any, fileName string, cfg *configs.MainConfiguration ) (string, error) {
    var sqlScript string
	var rows []T
    result := db.Where(where).Order("created_at DESC").Find(&rows)
    if result.Error != nil {
        logger.Error(result.Error)
		return "", result.Error
    }
    for _, row := range rows {
        // Create a new DB session with DryRun mode
        stmt := db.Session(&gorm.Session{DryRun: true}).Clauses(clause.OnConflict{
            DoNothing: true, // To generate INSERT OR IGNORE, use DoNothing
        }).Create(&row).Statement

        sql := stmt.SQL.String()

        // Replace "INSERT" with "INSERT OR REPLACE" in the generated SQL query
        sql = "INSERT OR REPLACE" + sql[len("INSERT"):]

        // Append the generated SQL query to the script
        sqlScript += sql + ";\n"
    } 
	if fileName != "" {
		fileName = fmt.Sprintf("/tmp/%s.sql", fileName)
		return fileName, SaveToFile(fileName, sqlScript)
	}
    return sqlScript, nil
}

func SaveToFile(filename, data string) (error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

	_, err = file.WriteString(data)
    if err != nil {
        return err
    }

    return nil
}
