package migration

import (
	"reflect"
	"runtime"

	"gorm.io/gorm"
)

type Migration struct {
	Id       string
	DateTime string
	Migrate  func(db *gorm.DB) error
}

var Migrations = []Migration{}

func AddMigration(migration func(db *gorm.DB) error, dateTime string) {
	migrationName := runtime.FuncForPC(reflect.ValueOf(migration).Pointer()).Name()
	Migrations = append(Migrations, Migration{
		Id: migrationName,
		DateTime: dateTime,
		Migrate: migration,
	})
}

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

	AddMigration(DropOwnerColumnFromSubnetState, "2024-09-16 5:23PM") 
	AddMigration(DropTopicIdColumnFromMessageState, "2024-09-16 5:00PM") 
	AddMigration(DropAttachmentsColumnFromMessageState, "2024-09-16 5:31PM") 

}


