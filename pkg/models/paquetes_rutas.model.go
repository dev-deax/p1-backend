package models

import "gorm.io/gorm"

type PaqueteRuta struct {
	gorm.Model
	PaqueteID int `gorm:"foreignKey:PaqueteID"`
	RutaID    int `gorm:"foreignKey:RutaID"`
}
