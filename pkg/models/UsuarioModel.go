package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	ID       int    `json:"id"`
	Nombre   string `gorm:"type:varchar(255)" json:"nombre"`
	Apellido string `gorm:"type:varchar(255)" json:"apellido"`
	Email    string `gorm:"type:varchar(255);unique" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	Rol      string `gorm:"type:enum('administrador', 'operador', 'recepcionista')"`
	Activo   bool   `gorm:"default:true" json:"activo"`
}
