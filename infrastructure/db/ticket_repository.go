package db

import (
	"trainingbackenddot/domain"

	"gorm.io/gorm"
)

type TicketRepository struct {
	DB *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *TicketRepository {
	return &TicketRepository{DB: db}
}

// Create a new ticket
func (r *TicketRepository) Create(ticket *domain.Ticket) error {
	return r.DB.Create(ticket).Error
}

// Get tickets by ID
func (r *TicketRepository) GetByID(id uint) (*domain.Ticket, error) {
	var ticket domain.Ticket
	if err := r.DB.First(&ticket, id).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

// Get tickets based on schedule and seat number
func (r *TicketRepository) GetByScheduleAndSeat(scheduleID uint, seatNumber string) (*domain.Ticket, error) {
	var ticket domain.Ticket
	if err := r.DB.Where("schedule_id = ? AND seat_number = ?", scheduleID, seatNumber).First(&ticket).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

// Updating tickets
func (r *TicketRepository) Update(ticket *domain.Ticket) error {
	return r.DB.Save(ticket).Error
}
