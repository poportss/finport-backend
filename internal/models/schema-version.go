package models

import "gorm.io/gorm"

type SchemaVersion struct {
	gorm.Model
	Service string `gorm:"service"`
	Version int    `gorm:"version"`
}
