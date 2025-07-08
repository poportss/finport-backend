package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	UserID       uint       `gorm:"not null;index"`
	Name         string     `gorm:"not null"`
	WalletTypeID uint       `gorm:"not null;index"`
	WalletType   WalletType `gorm:"foreignKey:WalletTypeID"`
}
