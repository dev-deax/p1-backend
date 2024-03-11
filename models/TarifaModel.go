package models

import "gorm.io/gorm"

type Tarifa struct {
	gorm.Model
	Tipo  string
	Valor float64
}
