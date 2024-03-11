package models

import (
	"time"

	"gorm.io/gorm"
)

type Factura struct {
	gorm.Model
	Fecha       time.Time
	ClienteID   int
	Paquetes    []Paquete
	PrecioTotal float64
}
