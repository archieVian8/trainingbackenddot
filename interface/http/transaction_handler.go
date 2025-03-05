package http

import (
	"net/http"
	"strconv"

	"trainingbackenddot/usecase"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	TransactionUC usecase.TransactionUsecase
}

func NewTransactionHandler(transactionUC usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{TransactionUC: transactionUC}
}

// Endpoint Ticket Payment
func (h *TransactionHandler) PayTicket(c *gin.Context) {
	var request struct {
		UserID        uint    `json:"user_id"`
		PaymentMethod string  `json:"payment_method"`
		Amount        float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	ticketID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ticket ID"})
		return
	}

	transaction, message, err := h.TransactionUC.ProcessPayment(uint(ticketID), request.UserID, request.PaymentMethod, request.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     message,
		"transaction": transaction,
	})
}

// Endpoint Admin View All Transactions
func (h *TransactionHandler) ViewAllTransactions(c *gin.Context) {
	transactions, err := h.TransactionUC.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// Endpoint Admin View Daily Film Sales Report
func (h *TransactionHandler) ViewDailySalesByFilm(c *gin.Context) {
	date := c.Query("date")
	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Date parameter is required (YYYY-MM-DD)"})
		return
	}
	reports, err := h.TransactionUC.GetDailySalesReportByFilm(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reports)
}

// Endpoint Admin View Monthly Film Sales Report
func (h *TransactionHandler) ViewMonthlySalesByFilm(c *gin.Context) {
	month := c.Query("date")
	if month == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Month parameter is required (YYYY-MM)"})
		return
	}
	reports, err := h.TransactionUC.GetMonthlySalesReportByFilm(month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reports)
}

// Endpoint Admin View Daily Studio Sales Report
func (h *TransactionHandler) ViewDailySalesByStudio(c *gin.Context) {
	date := c.Query("date")
	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Date parameter is required (YYYY-MM-DD)"})
		return
	}
	reports, err := h.TransactionUC.GetDailySalesReportByStudio(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reports)
}

// Endpoint Admin View Monthly Studio Sales Report
func (h *TransactionHandler) ViewMonthlySalesByStudio(c *gin.Context) {
	month := c.Query("date")
	if month == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Month parameter is required (YYYY-MM)"})
		return
	}
	reports, err := h.TransactionUC.GetMonthlySalesReportByStudio(month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reports)
}
