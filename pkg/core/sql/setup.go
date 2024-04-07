package sql

import (
	"fmt"
	"os"
	"strings"
	"time"

	config "github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/internal/sql/models"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	dbLogger "gorm.io/gorm/logger"
)

var Db *gorm.DB
var SqlDBErr error

var logger = &log.Logger

func InitializeDb(driver string, dsn string) (*gorm.DB, error) {
	logger.Infof("Initializing %s db...", driver)
	var dialect gorm.Dialector
	switch driver {
	case "postgres":
		dialect = postgres.Open(dsn)
	case "mysql":
		dialect = mysql.Open(dsn)
	default:
		dialect = sqlite.Open(dsn)
	}
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: dbLogger.Default.LogMode(logLevel()),
	})

	if err != nil {
		return nil, err
	}
	for _, model := range models.Models {
		err := db.AutoMigrate(&model)
		if err != nil {
			logger.Errorf("UnmarshalError %v", err)
		}

	}

	return db, err
}

func Init() {
	cfg := config.Config
	logger.Infof("DB Dialect %v", config.Config.SQLDB)
	Db, SqlDBErr = InitializeDb(config.Config.SQLDB.DbDialect, getDSN(cfg.SQLDB.DbDialect))
	if SqlDBErr != nil {
		panic(SqlDBErr)
	}
	db, err := Db.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(cfg.SQLDB.DbMaxConnLifetime)
	db.SetMaxOpenConns(cfg.SQLDB.DbMaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.SQLDB.DbMaxConnLifetime) * time.Second)
}

func logLevel() dbLogger.LogLevel {
	if config.Config.LogLevel == "info" {
		return dbLogger.Info
	}
	return dbLogger.Error
}

func getDSN(dialect string) string {
	cfg := config.Config
	dsn := ""
	switch strings.ToLower(config.Config.SQLDB.DbDialect) {
	case "sqlite":
		err := os.MkdirAll(cfg.SQLDB.DbStoragePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
		if strings.HasSuffix(cfg.SQLDB.DbStoragePath, "/") {
			dsn = fmt.Sprintf("%sdb.sqlite", cfg.SQLDB.DbStoragePath)
		} else {
			dsn = fmt.Sprintf("%s/db.sqlite", cfg.SQLDB.DbStoragePath)
		}
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config.SQLDB.DbUser, config.Config.SQLDB.DbPassword, config.Config.SQLDB.DbHost, config.Config.SQLDB.DbPort, config.Config.SQLDB.DbDatabase)
	case "postgres":
	default:
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", config.Config.SQLDB.DbHost, config.Config.SQLDB.DbUser, config.Config.SQLDB.DbPassword, config.Config.SQLDB.DbDatabase, config.Config.SQLDB.DbPort, config.Config.SQLDB.DbSSLMode, config.Config.SQLDB.DbTimezone)
	}
	return dsn
}
