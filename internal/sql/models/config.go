package models

import "gorm.io/gorm"

type Config struct {
	gorm.Model
	Key   string `gorm:"key;unique"`
	Value string `gorm:"value"`
}
