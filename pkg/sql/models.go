package sql

import (
	"gorm.io/gorm"
)

type MessageModel struct {
	gorm.Model
	ID          int64   `gorm:"primaryKey;autoIncrement:true"`
	Message        string `json:"message"`
	Subject     string  `json:"subject"`
	To string `json:"to"`
	Signature string `json:"signature"`

	AvailableBalance    float64 `json:"available_balance"`
}

type ConfigModel struct {
	gorm.Model
	Key   string `gorm:"key;unique"`
	Value string `gorm:"value"`
}


var Migrations = []string{CreateConfigTable, UpdatedInscriptionsTable, CreateAccountTable, CreateBrc20AccountBalanceTable, CreateBrc20TokenTable, CreateInscriptionTable}
