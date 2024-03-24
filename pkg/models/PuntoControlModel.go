package models

import "gorm.io/gorm"

type PuntoControl struct {
	gorm.Model
	ID              int     `gorm:"primaryKey" json:"id"`
	Nombre          string  `gorm:"type:varchar(255)" json:"nombre"`
	TarifaOperacion float64 `gorm:"type:float" json:"tarifa_operacion"`
	CapacidadCola   int     `gorm:"type:int" json:"capacidad_cola"`
	RutaID          int     `json:"ruta_id"`
	Ruta            Ruta

	Usuarios []Usuario `gorm:"many2many:usuarios_punto_controls"`
	Paquetes []Paquete `gorm:"many2many:paquetes_puntos_controls;"`
}
