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

// Fetch Data Daily Film Sales Report
func (r *TransactionRepository) FetchDailySalesReportByFilm(date string) ([]domain.FilmSalesReport, error) {
	var reports []domain.FilmSalesReport
	query := `
        SELECT f.id AS film_id, f.title, COUNT(t.id) AS total_tickets_sold, SUM(s.price) AS total_revenue
		FROM transactions t
		JOIN tickets tk ON t.ticket_id = tk.id
		JOIN schedules s ON tk.schedule_id = s.id
		JOIN films f ON s.film_id = f.id
		WHERE DATE(t.created_at) = ?
		GROUP BY f.id, f.title
		ORDER BY total_tickets_sold DESC
	`
	if err := r.DB.Raw(query, date).Scan(&reports).Error; err != nil {
		return nil, err
	}
	return reports, nil
}

// Fetch Data Monthly Film Sales Report
func (r *TransactionRepository) FetchMonthlySalesReportByFilm(month string) ([]domain.FilmSalesReport, error) {
	var reports []domain.FilmSalesReport
	query := `
        SELECT f.id AS film_id, f.title, COUNT(t.id) AS total_tickets_sold, SUM(s.price) AS total_revenue
        FROM transactions t
        JOIN tickets tk ON t.ticket_id = tk.id
        JOIN schedules s ON tk.schedule_id = s.id
        JOIN films f ON s.film_id = f.id
        WHERE TO_CHAR(t.created_at, 'YYYY-MM') = $1
        GROUP BY f.id, f.title
        ORDER BY total_tickets_sold DESC
    `
	if err := r.DB.Raw(query, month).Scan(&reports).Error; err != nil {
		return nil, err
	}
	return reports, nil
}

// Fetch Data Daily Studio Sales Report
func (r *TransactionRepository) FetchDailySalesReportByStudio(date string) ([]domain.StudioSalesReport, error) {
	var reports []domain.StudioSalesReport
	query := `
        SELECT st.id AS studio_id, st.name, COUNT(t.id) AS total_tickets_sold, SUM(s.price) AS total_revenue
        FROM transactions t
        JOIN tickets tk ON t.ticket_id = tk.id
        JOIN schedules s ON tk.schedule_id = s.id
        JOIN studios st ON s.studio_id = st.id
        WHERE DATE(t.created_at) = ?
        GROUP BY st.id, st.name
        ORDER BY total_tickets_sold DESC
    `
	if err := r.DB.Raw(query, date).Scan(&reports).Error; err != nil {
		return nil, err
	}
	return reports, nil
}

// Fetch Data Monthly Studio Sales Report
func (r *TransactionRepository) FetchMonthlySalesReportByStudio(month string) ([]domain.StudioSalesReport, error) {
	var reports []domain.StudioSalesReport
	query := `
        SELECT st.id AS studio_id, st.name, COUNT(t.id) AS total_tickets_sold, SUM(s.price) AS total_revenue
        FROM transactions t
        JOIN tickets tk ON t.ticket_id = tk.id
        JOIN schedules s ON tk.schedule_id = s.id
        JOIN studios st ON s.studio_id = st.id
        WHERE TO_CHAR(t.created_at, 'YYYY-MM') = $1
        GROUP BY st.id, st.name
        ORDER BY total_tickets_sold DESC
    `
	if err := r.DB.Raw(query, month).Scan(&reports).Error; err != nil {
		return nil, err
	}
	return reports, nil
}
