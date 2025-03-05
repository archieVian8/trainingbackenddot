package db

import (
	"trainingbackenddot/domain"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

// Create a new transaction
func (r *TransactionRepository) Create(transaction *domain.Transaction) error {
	return r.DB.Create(transaction).Error
}

// Get transactions based on ticket ID
func (r *TransactionRepository) GetByTicketID(ticketID uint) (*domain.Transaction, error) {
	var transaction domain.Transaction
	if err := r.DB.
		Preload("Ticket").
		Preload("Ticket.Schedule").
		Preload("Ticket.Schedule.Studio").
		Preload("Ticket.Schedule.Film").
		Where("ticket_id = ?", ticketID).
		First(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

// Get all transactions (admin only)
func (r *TransactionRepository) GetAll() ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.DB.
		Preload("Ticket").
		Preload("Ticket.Schedule").
		Preload("Ticket.Schedule.Studio").
		Preload("Ticket.Schedule.Film").
		Find(&transactions).Error
	return transactions, err
}

// Update an existing transaction
func (r *TransactionRepository) Update(transaction *domain.Transaction) error {
	return r.DB.Save(transaction).Error
}
