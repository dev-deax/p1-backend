package models

import "gorm.io/gorm"

type Ruta struct {
	gorm.Model
	ID            int    `json:"id"`
	Nombre        string `gorm:"type:varchar(255)" json:"nombre"`
	Capacidad     int    `gorm:"type:int" json:"capacidad"`
	DestinoID     int    `json:"destinoID"`
	Destino       Destino
	PuntosControl []PuntoControl `gorm:"many2many:ruta_punto_control"`
	// TarifaOperacion float64 `gorm:"type:float" json:"tarifaOperacion"`
	// Activo          bool    `gorm:"type:boolean" json:"activo"`
}
