package migration

import (
	"gorm.io/gorm"
)

type Migration struct {
	Id       string
	DateTime string
	Migrate  func(db *gorm.DB) error
}

var Migrations = []Migration{}

func init() {
	// Migrations = append(Migrations, Migration{
	// 	Id: "migrate-auth-index",
	// 	DateTime: "2024-04-15 6:00AM",
	// 	Migrate: MigrateAuthIndex,
	// })
	Migrations = append(Migrations, Migration{
		Id: "migrate-add-claimed-to-event-counter",
		DateTime: "2024-08-12 6:00AM",
		Migrate: AddClaimedFieldToEventCount,
	})

}
