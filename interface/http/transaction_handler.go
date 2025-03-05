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

	ticketID, err := strconv.ParseUint(c.Param("ticket_id"), 10, 64)
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
