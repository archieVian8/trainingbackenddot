package domain

import "gorm.io/gorm"

type Studio struct {
	gorm.Model
	Name       string `gorm:"not null"`
	Capacity   int    `gorm:"not null"`
	Facilities string
}

func (Studio) TableName() string {
	return "studios"
}
