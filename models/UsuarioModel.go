package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nombre     string
	Rol        string
	Email      string
	Contrasena string
}
