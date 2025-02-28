package domain

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	TicketID      uint
	Ticket        Ticket
	PaymentMethod string `gorm:"not null"`
	PaymentStatus string `gorm:"default:'pending'"`
}
