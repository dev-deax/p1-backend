package models

import "gorm.io/gorm"

type Tarifa struct {
	gorm.Model
	ID    int     `gorm:"primaryKey" json:"id"`
	Tipo  string  `gorm:"type:varchar(50);not null" json:"tipo"`
	Valor float64 `gorm:"type:float;not null" json:"valor"`
}
