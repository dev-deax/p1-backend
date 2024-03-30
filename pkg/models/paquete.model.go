package models

import (
	"gorm.io/gorm"
)

type Paquete struct {
	gorm.Model
	ID              int            `gorm:"primaryKey" json:"id"`
	Peso            float64        `gorm:"type:float"`
	CuotaDestino    float64        `gorm:"type:float"`
	TarifaOperacion float64        `gorm:"type:float"`
	PecioLb         float64        `gorm:"type:float"`
	Factura         Factura        `gorm:"many2many:detalle_factura;"`
	PuntosControl   []PuntoControl `gorm:"many2many:paquetes_puntos_controls;"`
	Rutas           []Ruta         `gorm:"many2many:paquetes_rutas;"`
	DestinoID       int            `json:"destino"`
	EstadoID        int            `json:"estado_id"`
	Destino         Destino
	Estado          Estado
	// PuntoControlID int            `json:"punto_control"`
	// FacturaID      int            `json:"factura"`
	// FechaEntrega   *time.Time     `gorm:"type:datetime"`
	// PrecioTotal  float64    `gorm:"type:float"`
	// FechaIngreso *time.Time `gorm:"type:datetime"`
	// ClienteID int `json:"cliente"`
	// Cliente   Cliente
}
