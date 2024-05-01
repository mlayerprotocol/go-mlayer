package models

import (
	"github.com/mlayerprotocol/go-mlayer/entities"
	"gorm.io/gorm"
)

type WalletState struct {
	entities.Wallet
	BaseModel
}

func (d *WalletState) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == "" {
		hash, err := entities.GetId(*d)
		if err != nil {
			panic(err)
		}
		d.ID = hash
	}
	return nil
}

type WalletEvent struct {
	entities.Event
	BaseModel
	// WalletID     uint64
	// Wallet		WalletState
}
