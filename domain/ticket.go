package domain

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	ScheduleID uint
	Schedule   Schedule
	UserID     uint
	SeatNumber string `gorm:"not null"`
	Status     string `gorm:"default:'pending'"`
}
