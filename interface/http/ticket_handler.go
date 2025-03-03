package http

import (
	"net/http"
	"strconv"

	"trainingbackenddot/usecase"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	TicketUC usecase.TicketUsecase
}

func NewTicketHandler(ticketUC usecase.TicketUsecase) *TicketHandler {
	return &TicketHandler{TicketUC: ticketUC}
}

// Endpoint Booking tickets (selecting seats)
func (h *TicketHandler) BookTicket(c *gin.Context) {
	var request struct {
		ScheduleID uint   `json:"schedule_id"`
		UserID     uint   `json:"user_id"`
		SeatNumber string `json:"seat_number"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ticket booking process
	ticket, err := h.TicketUC.BookTicket(request.ScheduleID, request.UserID, request.SeatNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to book ticket"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

// Endpoint View ticket details
func (h *TicketHandler) GetTicket(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	ticket, err := h.TicketUC.GetTicket(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	c.JSON(http.StatusOK, ticket)
}
