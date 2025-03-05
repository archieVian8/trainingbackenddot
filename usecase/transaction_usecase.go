package usecase

import (
	"errors"
	"strings"

	"trainingbackenddot/domain"
	"trainingbackenddot/infrastructure/db"
)

type TransactionUsecase interface {
	PayTicket(ticketID uint, paymentMethod string) error
	GetAllTransactions() ([]domain.Transaction, error)
	ProcessPayment(ticketID uint, userID uint, paymentMethod string, amount float64) (*domain.Transaction, string, error)
}

type transactionUsecase struct {
	TransactionRepo *db.TransactionRepository
	TicketRepo      *db.TicketRepository
	ScheduleRepo    *db.ScheduleRepository
	StudioRepo      *db.StudioRepository
}

func NewTransactionUsecase(
	transactionRepo *db.TransactionRepository,
	ticketRepo *db.TicketRepository,
	scheduleRepo *db.ScheduleRepository,
	studioRepo *db.StudioRepository) TransactionUsecase {

	return &transactionUsecase{
		TransactionRepo: transactionRepo,
		TicketRepo:      ticketRepo,
		ScheduleRepo:    scheduleRepo,
		StudioRepo:      studioRepo,
	}
}

// Function for Ticket payment
func (u *transactionUsecase) PayTicket(ticketID uint, paymentMethod string) error {
	// Get a ticket
	ticket, err := u.TicketRepo.GetByID(ticketID)
	if err != nil {
		return errors.New("ticket not found")
	}

	// Check if it has been paid
	existingTransaction, err := u.TransactionRepo.GetByTicketID(ticketID)
	if err == nil && existingTransaction != nil {
		return errors.New("ticket already paid")
	}

	// Create a new transaction
	transaction := &domain.Transaction{
		TicketID:      ticket.ID,
		PaymentMethod: paymentMethod,
		PaymentStatus: "success",
	}

	err = u.TransactionRepo.Create(transaction)
	if err != nil {
		return err
	}

	// Update ticket status to "confirmed"
	ticket.Status = "confirmed"
	err = u.TicketRepo.Update(ticket)
	if err != nil {
		return err
	}

	// Get schedule and studio
	schedule, err := u.ScheduleRepo.GetByID(ticket.ScheduleID)
	if err != nil {
		return errors.New("schedule not found")
	}

	studio, err := u.StudioRepo.GetByID(schedule.StudioID)
	if err != nil {
		return errors.New("studio not found")
	}

	// Reduce studio capacity based on number of seats booked
	seatCount := len(strings.Split(ticket.SeatNumber, ","))
	if studio.Capacity < seatCount {
		return errors.New("not enough seats available")
	}
	studio.Capacity -= seatCount

	// Update studio capacity
	err = u.StudioRepo.Update(studio)
	if err != nil {
		return errors.New("failed to update studio capacity")
	}

	return nil
}

// Function for Admin view all transactions
func (u *transactionUsecase) GetAllTransactions() ([]domain.Transaction, error) {
	return u.TransactionRepo.GetAll()
}

// Function for Process Payment
func (u *transactionUsecase) ProcessPayment(ticketID uint, userID uint, paymentMethod string, amount float64) (*domain.Transaction, string, error) {
	ticket, err := u.TicketRepo.GetByID(ticketID)
	if err != nil {
		return nil, "Payment failed, ticket not found", errors.New("ticket not found")
	}

	if ticket.UserID != userID {
		return nil, "Payment failed, unauthorized transaction", errors.New("unauthorized transaction")
	}

	schedule, err := u.ScheduleRepo.GetByID(ticket.ScheduleID)
	if err != nil {
		return nil, "Payment failed, schedule not found", errors.New("schedule not found")
	}

	promoPrice := schedule.PromoPrice
	if promoPrice == 0 {
		promoPrice = schedule.Price
	}

	if amount > promoPrice {
		return nil, "Payment failed, money more than price", errors.New("amount exceeds price")
	}
	if amount < promoPrice {
		return nil, "Payment failed, money less than price", errors.New("amount less than price")
	}

	existingTransaction, err := u.TransactionRepo.GetByTicketID(ticketID)
	if err == nil && existingTransaction != nil && existingTransaction.PaymentStatus == "success" {
		return nil, "Payment failed, ticket already paid", errors.New("ticket already paid")
	}

	transaction := &domain.Transaction{
		TicketID:      ticket.ID,
		PaymentMethod: paymentMethod,
		PaymentStatus: "success",
	}

	err = u.TransactionRepo.Create(transaction)
	if err != nil {
		return nil, "Payment failed, transaction error", err
	}

	ticket.Status = "confirmed"
	err = u.TicketRepo.Update(ticket)
	if err != nil {
		return nil, "Payment failed, ticket update error", err
	}

	seatCount := len(strings.Split(ticket.SeatNumber, ","))
	schedule.Studio.Capacity -= seatCount
	err = u.StudioRepo.Update(&schedule.Studio)
	if err != nil {
		return nil, "Payment failed, studio update error", err
	}

	transaction, err = u.TransactionRepo.GetByTicketID(ticket.ID)
	if err != nil {
		return nil, "Payment failed, failed to retrieve transaction", err
	}

	return transaction, "Payment successful", nil
}
