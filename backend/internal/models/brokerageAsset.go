package models

import "gorm.io/gorm"

type BrokerageAsset struct {
	gorm.Model
	BrokerageID uint   `gorm:"not null;index"`
	Symbol      string // ex: PETR4
	Quantity    float64
	BuyPrice    float64
	NowPrice    float64
	AvgPrice    float64
}
