package domain

import "gorm.io/gorm"

type Schedule struct {
	gorm.Model
	StudioID uint
	Studio   Studio
	FilmID   uint
	Film     Film
	ShowTime string `gorm:"not null"`
	Price    float64
}
