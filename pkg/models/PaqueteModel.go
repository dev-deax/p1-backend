package models

import (
	"time"

	"gorm.io/gorm"
)

type Paquete struct {
	gorm.Model
	ID           int        `gorm:"primaryKey" json:"id"`
	Peso         float64    `gorm:"type:float"`
	PrecioTotal  float64    `gorm:"type:float"`
	Estado       string     `gorm:"type:enum('en_bodega', 'en_ruta', 'entregado')"`
	FechaIngreso *time.Time `gorm:"type:datetime"`
	FechaEntrega *time.Time `gorm:"type:datetime"`

	ClienteID      int `json:"cliente"`
	Cliente        Cliente
	DestinoID      int `json:"destino"`
	Destino        Destino
	RutaID         int `json:"ruta"`
	Ruta           Ruta
	PuntoControlID int `json:"punto_control"`

	PuntosControl []PuntoControl `gorm:"many2many:paquetes_puntos_controls;"`
}
