package models

import "gorm.io/gorm"

type Tarifas struct {
	gorm.Model
	ID    int     `gorm:"primaryKey"`
	Tipo  string  `gorm:"type:varchar(50);not null" json:"tipo"`
	Valor float64 `gorm:"type:float;not null" json:"valor"`
}
