package usecase

import (
	"errors"

	"trainingbackenddot/domain"
	"trainingbackenddot/infrastructure/db"
)

type TicketUsecase interface {
	BookTicket(scheduleID, userID uint, seatNumber []string) ([]*domain.Ticket, error)
	GetTicket(id uint) (*domain.Ticket, error)
}

type ticketUsecase struct {
	ticketRepo   *db.TicketRepository
	scheduleRepo *db.ScheduleRepository
}

func NewTicketUsecase(ticketRepo *db.TicketRepository, scheduleRepo *db.ScheduleRepository) TicketUsecase {
	return &ticketUsecase{ticketRepo: ticketRepo, scheduleRepo: scheduleRepo}
}

// Function for Booking Ticket
func (u *ticketUsecase) BookTicket(scheduleID, userID uint, seatNumbers []string) ([]*domain.Ticket, error) {
	schedule, err := u.scheduleRepo.GetByID(scheduleID)
	if err != nil {
		return nil, errors.New("schedule not found")
	}

	if len(seatNumbers) > schedule.Studio.Capacity {
		return nil, errors.New("not enough available seats")
	}

	var tickets []*domain.Ticket
	for _, seat := range seatNumbers {
		existingTicket, _ := u.ticketRepo.GetByScheduleAndSeat(scheduleID, seat)
		if existingTicket != nil {
			return nil, errors.New("seat " + seat + " already booked")
		}

		ticket := &domain.Ticket{
			ScheduleID: scheduleID,
			UserID:     userID,
			SeatNumber: seat,
			Status:     "pending",
		}

		err := u.ticketRepo.Create(ticket)
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

// Function for Viewing Ticket Details
func (u *ticketUsecase) GetTicket(id uint) (*domain.Ticket, error) {
	return u.ticketRepo.GetByID(id)
}
