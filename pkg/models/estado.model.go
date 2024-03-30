package models

import "gorm.io/gorm"

type Estado struct {
	gorm.Model
	ID   int    `json:"id"`
	Nome string `gorm:"type:varchar(50);not null" json:"nome"`
}
