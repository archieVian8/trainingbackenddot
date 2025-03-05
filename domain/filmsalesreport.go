package domain

type FilmSalesReport struct {
	FilmID           uint    `json:"film_id"`
	Title            string  `json:"title"`
	TotalTicketsSold int     `json:"total_tickets_sold"`
	TotalRevenue     float64 `json:"total_revenue"`
}
