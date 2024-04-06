package models

import "gorm.io/gorm"

type PaquetesRutas struct {
	gorm.Model
	PaqueteID int `gorm:"foreignKey:PaqueteID"`
	RutaID    int `gorm:"foreignKey:RutaID"`
	Paquete   Paquete
	Ruta      Ruta
}
