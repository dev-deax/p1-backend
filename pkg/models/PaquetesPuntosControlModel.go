package models

import (
	"time"

	"gorm.io/gorm"
)

type PaquetesPuntosControl struct {
	gorm.Model
	PaqueteID            int        `gorm:"primaryKey" json:"paquete_id"`
	PuntoControlID       int        `gorm:"primaryKey" json:"punto_control_id"`
	TiempoEnPuntoControl float64    `gorm:"type:float" json:"tiempo_en_punto_control"`
	FechaLlegada         *time.Time `gorm:"type:datetime" json:"fecha_llegada"`
	FechaSalida          *time.Time `gorm:"type:datetime" json:"fecha_salida"`
}
