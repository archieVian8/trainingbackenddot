package domain

import "gorm.io/gorm"

type Studio struct {
	gorm.Model
	Name       string `gorm:"unique;not null"`
	Capacity   int    `gorm:"not null"`
	Facilities string
}
