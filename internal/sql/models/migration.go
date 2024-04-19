package models

type MigrationState struct {
	Key string `gorm:"unique"`
	BaseModel
}


