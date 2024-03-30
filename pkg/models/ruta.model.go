package models

import "gorm.io/gorm"

type Ruta struct {
	gorm.Model
	ID        int       `json:"id"`
	Nombre    string    `gorm:"type:varchar(255)" json:"nombre"`
	Capacidad int       `gorm:"type:int" json:"capacidad"`
	Activo    bool      `gorm:"default:true" json:"activo"`
	DestinoID int       `json:"destinoID"`
	Paquetes  []Paquete `gorm:"many2many:paquetes_rutas;"`
	Destino   Destino
	// PuntosControl []PuntoControl
	// TarifaOperacion float64 `gorm:"type:float" json:"tarifaOperacion"`
	// Activo          bool    `gorm:"type:boolean" json:"activo"`
}
