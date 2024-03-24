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
	Cliente   Cliente       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Subtotal  float64       `gorm:"autoCreateTime" json:"subtotal"`
	Impuestos float64       `gorm:"autoCreateTime" json:"impuestos"`
	Total     float64       `gorm:"autoCreateTime" json:"total"`
	Items     []ItemFactura `gorm:"foreignKey:FacturaID"`
}
