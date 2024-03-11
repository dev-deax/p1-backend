package models

import "gorm.io/gorm"

type Paquete struct {
	gorm.Model
	Peso      float64
	Destino   string
	Precio    float64
	Estado    string
	ClienteID int
}
