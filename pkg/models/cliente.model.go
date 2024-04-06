package models

import "gorm.io/gorm"

type Cliente struct {
	gorm.Model
	ID        int    `gorm:"primaryKey"`
	NIT       int    `gorm:"type:int;unique" json:"nit"`
	Nombre    string `gorm:"type:varchar(255)" json:"nombre"`
	Apellido  string `gorm:"type:varchar(255)" json:"apellido"`
	Dirección string `gorm:"type:varchar(255)" json:"direccion"`
	Teléfono  int    `gorm:"type:int" json:"telefono"`
}
