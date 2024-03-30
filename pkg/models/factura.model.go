package models

import (
	"time"

	"gorm.io/gorm"
)

type Factura struct {
	gorm.Model
	ID        int       `gorm:"primaryKey" json:"id"`
	Fecha     time.Time `gorm:"autoCreateTime" json:"fecha"`
	ClienteID int
	Cliente   Cliente   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Total     float64   `gorm:"autoCreateTime" json:"total"`
	PaqueteID int       `json:"paquete"` // FacturaID int `json:"factura"`
	Paquetes  []Paquete `gorm:"many2many:detalle_factura;"`
}
