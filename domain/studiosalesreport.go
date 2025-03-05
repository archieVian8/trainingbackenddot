package domain

type StudioSalesReport struct {
	StudioID         uint    `json:"studio_id"`
	Name             string  `json:"name"`
	TotalTicketsSold int     `json:"total_tickets_sold"`
	TotalRevenue     float64 `json:"total_revenue"`
}
