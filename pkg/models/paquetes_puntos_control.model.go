package models

import (
	"time"

	"gorm.io/gorm"
)

type PaquetesPuntosControl struct {
	gorm.Model
	//no id on PaquetesPuntosControl model gorm
	TiempoPuntoControl float64    `gorm:"type:float" json:"tiempo_en_punto_control"`
	FechaLlegada       *time.Time `gorm:"type:datetime" json:"fecha_llegada"`
	FechaSalida        *time.Time `gorm:"type:datetime" json:"fecha_salida"`
	PaqueteID          int        `gorm:"primaryKey;autoIncrement:false;not null;" json:"paquete_id"`
	PuntoControlID     int        `gorm:"primaryKey;autoIncrement:false;not null;" json:"punto_control_id"`
	Paquete            Paquete
	PuntoControl       PuntoControl
}
