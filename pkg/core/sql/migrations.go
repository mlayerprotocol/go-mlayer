package sql

import (
	"gorm.io/gorm"
)

type Migration struct {
	id  string;
	migrate func(db gorm.DB)
}
 func Migrations(gorm.DB) []Migration {
	migrations := []Migration{}

	return migrations
 }
