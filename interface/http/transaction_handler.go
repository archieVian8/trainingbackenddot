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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var request struct {
		PaymentMethod string `json:"payment_method"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.TransactionUC.PayTicket(uint(id), request.PaymentMethod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process payment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment successful"})
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
