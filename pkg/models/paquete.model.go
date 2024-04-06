package models

import (
	"gorm.io/gorm"
)

type Paquete struct {
	gorm.Model
	ID              int     `gorm:"primaryKey"`
	Descripcion     string  `gorm:"type:varchar(256)"`
	Peso            float64 `gorm:"type:float"`
	CuotaDestino    float64 `gorm:"type:float"`
	TarifaOperacion float64 `gorm:"type:float"`
	PecioLb         float64 `gorm:"type:float"`
	FacturaID       int
	PuntosControl   []PuntoControl `gorm:"many2many:paquetes_puntos_controls;"`
	Rutas           []Ruta         `gorm:"many2many:paquetes_rutas;"`
	DestinoID       int
	EstadoID        int
	Destino         Destino
	Estado          Estado
	Factura         Factura
}
