package http

import (
	"net/http"

	"trainingbackenddot/usecase"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	NotificationUC *usecase.NotificationUsecase
}

func NewNotificationHandler(notificationUC *usecase.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{NotificationUC: notificationUC}
}

// GetNotifications handles retrieving notifications for a user
func (h *NotificationHandler) GetNotifications(c *gin.Context) {
	notifications, err := h.NotificationUC.GetNotifications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}
