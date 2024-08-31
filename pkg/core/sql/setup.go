package sql

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mlayerprotocol/go-mlayer/configs"
	config "github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/sql/migration"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	dbLogger "gorm.io/gorm/logger"
)

var SqlDb *gorm.DB
var SqlDBErr error

var logger = &log.Logger

func InitializeDb(driver string, dsn string) (*gorm.DB, error) {
	logger.Infof("Initializing %s db", driver)
	var dialect gorm.Dialector
	switch driver {
	case "postgres":
		dialect = postgres.Open(dsn)
	case "mysql":
		dialect = mysql.Open(dsn)
	default:
		dialect = sqlite.Open(dsn)
	}
	SqlDb, err := gorm.Open(dialect, &gorm.Config{
		Logger: dbLogger.Default.LogMode(logLevel()),
	})

	if err != nil {
		return nil, err
	}
	
	if driver == "sqlite" {
		//d, _ := SqlDb.DB()
		SqlDb.Exec("PRAGMA busy_timeout = 1000")
	}
	for _, model := range models.Models {
		err := SqlDb.AutoMigrate(&model)
		if err != nil {
			logger.Errorf("SQLD_BERROR: %v", err)
		}
	}
	
	
	return SqlDb, err
}

func GetTableName(table any, db *gorm.DB) string {
	stmt := &gorm.Statement{DB: db}
	stmt.Parse(table)
	return stmt.Schema.Table
}

func Init(cfg *configs.MainConfiguration) {
	
	SqlDb, SqlDBErr = InitializeDb(config.Config.SQLDB.DbDialect, getDSN(cfg))
	if SqlDBErr != nil {
		panic(SqlDBErr)
	}
	for _, migration := range migration.Migrations {
		var m models.MigrationState;
		key := strings.ToLower(fmt.Sprintf("%s:%s", migration.DateTime,  migration.Id))
		err := SqlDb.Where(models.MigrationState{Key: key }).First(&m).Error
		if err == gorm.ErrRecordNotFound {
			err := migration.Migrate(SqlDb)
			
			if err == nil {
				SqlDb.Create(&models.MigrationState{Key: key })
			} else {
				log.Logger.Error("Migration Error", err)
				panic(err)
			}
		}
	}
	db, err := SqlDb.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(cfg.SQLDB.DbMaxConnLifetime)
	db.SetMaxOpenConns(cfg.SQLDB.DbMaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.SQLDB.DbMaxConnLifetime) * time.Second)
	// SqlDb.Exec("DROP TRIGGER IF EXISTS subnet_events_sync_trigger;")
	counterTable := GetTableName(models.EventCounter{}, SqlDb)
	subnetSyncTrigger, subnetSyncFunc := EventSyncedTrigger(config.Config.SQLDB.DbDialect, GetTableName(models.SubnetEvent{}, SqlDb), counterTable)
	SqlDb.Exec(string(subnetSyncFunc))
	SqlDb.Exec(string(subnetSyncTrigger))

	authSyncTrigger, authSyncFunc := EventSyncedTrigger(config.Config.SQLDB.DbDialect, GetTableName(models.AuthorizationEvent{}, SqlDb), counterTable)
	SqlDb.Exec(string(authSyncFunc))
	SqlDb.Exec(string(authSyncTrigger))
	
	
}

func logLevel() dbLogger.LogLevel {
	if config.Config.LogLevel == "info" {
		return dbLogger.Info
	}
	return dbLogger.Error
}

func getDSN(cfg *configs.MainConfiguration) string {
	dsn := ""
	switch strings.ToLower(config.Config.SQLDB.DbDialect) {
	case "sqlite":
		err := os.MkdirAll(cfg.SQLDB.DbStoragePath, os.ModePerm)
		if err != nil {
			logger.Errorf("Error creating sqlite storage directory at %s", config.Config.SQLDB.DbStoragePath)
			panic(err)
		}
		if strings.HasSuffix(cfg.SQLDB.DbStoragePath, "/") {
			dsn = fmt.Sprintf("%sdb.sqlite", cfg.SQLDB.DbStoragePath)
		} else {
			dsn = fmt.Sprintf("%s/db.sqlite", cfg.SQLDB.DbStoragePath)
		}
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config.SQLDB.DbUser, config.Config.SQLDB.DbPassword, config.Config.SQLDB.DbHost, config.Config.SQLDB.DbPort, config.Config.SQLDB.DbDatabase)
	// case "postgres":
	default:
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", config.Config.SQLDB.DbHost, config.Config.SQLDB.DbUser, config.Config.SQLDB.DbPassword, config.Config.SQLDB.DbDatabase, config.Config.SQLDB.DbPort, config.Config.SQLDB.DbSSLMode, config.Config.SQLDB.DbTimezone)
	}
	return dsn
}
