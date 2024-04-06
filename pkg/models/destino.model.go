package models

import "gorm.io/gorm"

type Destino struct {
	gorm.Model
	ID           int     `gorm:"primaryKey"`
	Nombre       string  `gorm:"type:varchar(255); not null;" json:"nombre"`
	CuotaDestino float64 `json:"cuota_destino"`
	// RutaID       int     `json:"ruta_id"`
	// Ruta         []Ruta  `gorm:"foreignKey:DestinoID"`
}
