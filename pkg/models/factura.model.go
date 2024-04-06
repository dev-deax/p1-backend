package models

import (
	"gorm.io/gorm"
)

type Factura struct {
	gorm.Model
	ID        int       `gorm:"primaryKey"`
	Total     float64   `gorm:"autoCreateTime"`
	Cliente   Cliente   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Paquetes  []Paquete `gorm:"foreignKey:FacturaID"`
	ClienteID int
}
