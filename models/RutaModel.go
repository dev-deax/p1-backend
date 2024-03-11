package models

import "gorm.io/gorm"

type Ruta struct {
	gorm.Model
	Origen        string
	Destino       string
	PuntosControl []PuntoControl
}
