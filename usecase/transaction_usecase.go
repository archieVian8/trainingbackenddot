package usecase

import (
	"errors"

	"trainingbackenddot/domain"
	"trainingbackenddot/infrastructure/db"
)

type TransactionUsecase interface {
	PayTicket(ticketID uint, paymentMethod string) error
	GetAllTransactions() ([]domain.Transaction, error)
}

type transactionUsecase struct {
	TransactionRepo *db.TransactionRepository
	TicketRepo      *db.TicketRepository
}

func NewTransactionUsecase(transactionRepo *db.TransactionRepository, ticketRepo *db.TicketRepository) TransactionUsecase {
	return &transactionUsecase{
		TransactionRepo: transactionRepo,
		TicketRepo:      ticketRepo,
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

	// Update ticket status
	ticket.Status = "confirmed"
	err = u.TicketRepo.Update(ticket)
	if err != nil {
		return err
	}

	return nil
}

// Function for Admin view all transactions
func (u *transactionUsecase) GetAllTransactions() ([]domain.Transaction, error) {
	return u.TransactionRepo.GetAll()
}
