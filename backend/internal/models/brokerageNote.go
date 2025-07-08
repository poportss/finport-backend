package models

import (
	"gorm.io/gorm"
	"time"
)

type BrokerageNote struct {
	gorm.Model
	UserID      uint   `gorm:"not null;index"`
	BrokerageID uint   `gorm:"not null"`
	FilePath    string `gorm:"not null"`
	ProcessedAt *time.Time
}
