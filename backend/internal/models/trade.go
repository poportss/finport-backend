package models

import (
	"gorm.io/gorm"
	"time"
)

type Trade struct {
	gorm.Model
	UserID      uint      `gorm:"not null;index"`
	BrokerageID uint      `gorm:"not null"`
	AssetID     uint      `gorm:"not null"`
	Type        string    `gorm:"not null"` // Buy or Sell
	Quantity    int       `gorm:"not null"`
	Price       float64   `gorm:"not null"`
	Date        time.Time `gorm:"not null"`
	NoteNumber  string
}
