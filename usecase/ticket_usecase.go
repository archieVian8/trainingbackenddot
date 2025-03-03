package usecase

import (
	"errors"

	"trainingbackenddot/domain"
	"trainingbackenddot/infrastructure/db"
)

type TicketUsecase interface {
	BookTicket(scheduleID, userID uint, seatNumber string) (*domain.Ticket, error)
	GetTicket(id uint) (*domain.Ticket, error)
}

type ticketUsecase struct {
	ticketRepo *db.TicketRepository
}

func NewTicketUsecase(ticketRepo *db.TicketRepository) TicketUsecase {
	return &ticketUsecase{ticketRepo: ticketRepo}
}

// Function for Booking Ticket
func (u *ticketUsecase) BookTicket(scheduleID, userID uint, seatNumber string) (*domain.Ticket, error) {
	// Check if the seat is already booked
	existingTicket, _ := u.ticketRepo.GetByScheduleAndSeat(scheduleID, seatNumber)
	if existingTicket != nil {
		return nil, errors.New("seat already booked")
	}

	// Create a new ticket
	ticket := &domain.Ticket{
		ScheduleID: scheduleID,
		UserID:     userID,
		SeatNumber: seatNumber,
		Status:     "pending",
	}

	err := u.ticketRepo.Create(ticket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

// Function for Viewing Ticket Details
func (u *ticketUsecase) GetTicket(id uint) (*domain.Ticket, error) {
	return u.ticketRepo.GetByID(id)
}
