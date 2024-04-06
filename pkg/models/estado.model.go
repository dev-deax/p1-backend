package models

import "gorm.io/gorm"

type Estado struct {
	gorm.Model
	ID     int    `gorm:"primaryKey"`
	Nombre string `gorm:"type:varchar(50);not null" json:"nombre"`
}
