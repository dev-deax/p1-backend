package models

import "gorm.io/gorm"

type ItemFactura struct {
	gorm.Model
	FacturaID      int
	Factura        Factura
	Descripcion    string
	Cantidad       int
	PrecioUnitario float64
	Total          float64
}
