package models

import "gorm.io/gorm"

type PuntoControl struct {
	gorm.Model
	ID              int     `gorm:"primaryKey"`
	Nombre          string  `gorm:"type:varchar(255)" json:"nombre"`
	TarifaOperacion float64 `gorm:"type:float" json:"tarifa_operacion"`
	CapacidadCola   int     `gorm:"type:int" json:"capacidad_cola"`
	Activo          bool    `gorm:"default:true" json:"activo"`
	UsuarioID       int
	RutaID          int
	Ruta            Ruta
	Paquetes        []Paquete `gorm:"many2many:paquetes_puntos_controls;"`
	Usuario         Usuario
}
