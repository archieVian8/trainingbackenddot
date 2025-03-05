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
	if err := r.DB.Create(ticket).Error; err != nil {
		return err
	}

	return r.DB.Preload("Schedule").Preload("Schedule.Studio").Preload("Schedule.Film").
		First(ticket, ticket.ID).Error
}

// Get tickets by ID
func (r *TicketRepository) GetByID(ticketID uint) (*domain.Ticket, error) {
	var ticket domain.Ticket
	err := r.DB.
		Preload("Schedule").
		Preload("Schedule.Studio").
		Preload("Schedule.Film").
		Where("id = ?", ticketID).
		First(&ticket).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

// Get tickets based on schedule and seat number
func (r *TicketRepository) GetByScheduleAndSeat(scheduleID uint, seatNumber string) (*domain.Ticket, error) {
	var ticket domain.Ticket
	err := r.DB.
		Preload("Schedule").
		Preload("Schedule.Studio").
		Preload("Schedule.Film").
		Where("schedule_id = ? AND seat_number = ?", scheduleID, seatNumber).
		First(&ticket).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

// Updating tickets
func (r *TicketRepository) Update(ticket *domain.Ticket) error {
	return r.DB.Save(ticket).Error
}
