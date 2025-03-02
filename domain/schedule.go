package domain

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	StudioID   uint
	Studio     Studio
	FilmID     uint
	Film       Film
	ShowTime   string `gorm:"not null"`
	Price      float64
	Promo      int
	PromoPrice float64
	PromoTime  time.Time
	PromoEnds  time.Time
}
