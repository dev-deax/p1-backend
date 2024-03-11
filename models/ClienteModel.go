package models

import "gorm.io/gorm"

type Cliente struct {
	gorm.Model
	Nit       string
	Nombre    string
	Direccion string
	Telefono  string
}
