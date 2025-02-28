package domain

import "gorm.io/gorm"

type Film struct {
	gorm.Model
	Title       string `gorm:"unique;not null"`
	Genre       string `gorm:"not null"`
	Duration    int    `gorm:"not null"`
	Description string
}
