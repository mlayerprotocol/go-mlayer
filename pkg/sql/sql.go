package sql

import (
	"fmt"
	"strings"
	"time"

	db "github.com/cosmos/cosmos-db"
	"github.com/mlayerprotocol/go-mlayer/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var SqlDB *gorm.DB
var SqlDBErr error


func InitializeDb(driver string, dsn string, migrations []string) (*gorm.DB, error) {
	var dialect gorm.Dialector
	switch driver {
	case "postgress":
		dialect = postgres.Open(dsn)
	case "mysql":
		dialect = mysql.Open(dsn)
	default:
		dialect = sqlite.Open(dsn)
	}
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logLevel()),
	})

	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&ConfigModel{})
	db.AutoMigrate(&MessageModel{})
	

	return db, err
}

func init() {
	cfg := utils.Config

	SqlDB, SqlDBErr = InitializeDb(utils.Config.DbDialect, getDSN(cfg.DbDialect), Migrations)
	if SqlDBErr != nil {
		panic(SqlDBErr)
	}
	db, err := db.DB()
	if (err != nil) {
		panic(err)
	}
	db.SetMaxIdleConns(cfg.DbMaxIdleConns)
	db.SetMaxOpenConns(cfg.DbMaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.DbMaxConnLifetime) * time.Second)
}

func logLevel() logger.LogLevel {
	if utils.Config.LogLevel == "info" {
		return logger.Info
	}
	return logger.Silent
}

func getDSN(dialect string) string {
	cfg := utils.Config
	dsn := ""
	switch strings.ToLower(utils.Config.DbDialect) {
	case "sqlite":
		dsn = cfg.DbStoragePath/Users/Projects/go/src/go-ssrc/pkg/core/sql/query.go
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",  utils.Config.DbUser, utils.Config.DbPassword,  utils.Config.DbHost,  utils.Config.DbPort, utils.Config.DbDatabase)
	case "postgres":
	default:
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", utils.Config.DbHost, utils.Config.DbUser, utils.Config.DbPassword, utils.Config.DbDatabase, utils.Config.DbPort, utils.Config.DbSSLMode, utils.Config.DbTimezone )
	}
	return dsn;
}
