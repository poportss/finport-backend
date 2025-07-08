package models

import "gorm.io/gorm"

type Type string

const (
	CryptoWalletType Type = "crypto"
	StocksWalletType Type = "stocks"
	ReitsWalletType  Type = "reits"
	BondsWalletType  Type = "bonds"
	EtfWalletType    Type = "etfs"
)

type WalletType struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;not null"`
}
