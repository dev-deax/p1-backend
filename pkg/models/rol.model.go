package models

type Rol struct {
	ID     int    `gorm:"primaryKey"`
	Nombre string `gorm:"type:varchar(50);not null;" json:"nombre"`
}
