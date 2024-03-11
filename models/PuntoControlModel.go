package models

import "gorm.io/gorm"

type PuntoControl struct {
	gorm.Model
	Nombre string
	RutaID int
}
