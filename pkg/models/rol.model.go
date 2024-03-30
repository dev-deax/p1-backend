package models

type Rol struct {
	ID        int    `json:"id_rol"`
	NombreRol string `gorm:"type:varchar(50);not null;" json:"nombre_rol"`
}
