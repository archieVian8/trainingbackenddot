package http

import (
	"net/http"
	"strconv"
	"time"
	"trainingbackenddot/domain"
	"trainingbackenddot/usecase"

	"github.com/gin-gonic/gin"
)

type ScheduleHandler struct {
	ScheduleUC *usecase.ScheduleUsecase
}

func NewScheduleHandler(scheduleoUC *usecase.ScheduleUsecase) *ScheduleHandler {
	return &ScheduleHandler{ScheduleUC: scheduleoUC}
}

// Create New Schedule
func (h *ScheduleHandler) CreateSchedule(c *gin.Context) {
	var schedule domain.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.ScheduleUC.CreateSchedule(&schedule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Schedule created successfully"})
}

// View All Schedules
func (h *ScheduleHandler) ViewAllSchedules(c *gin.Context) {
	schedules, err := h.ScheduleUC.ViewAllSchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedules"})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// Update Schedule
func (h *ScheduleHandler) UpdateSchedule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
		return
	}

	var updatedSchedule domain.Schedule
	if err := c.ShouldBindJSON(&updatedSchedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.ScheduleUC.UpdateSchedule(uint(id), &updatedSchedule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule updated successfully"})
}

// Delete Schedule
func (h *ScheduleHandler) DeleteSchedule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
		return
	}

	err = h.ScheduleUC.DeleteSchedule(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete schedule"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule deleted successfully"})
}

// Create New Promo
func (h *ScheduleHandler) ApplyPromo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
		return
	}

	var promoData struct {
		Promo     int    `json:"promo"`
		PromoTime string `json:"promo_time"`
		PromoEnds string `json:"promo_ends"`
	}

	if err := c.ShouldBindJSON(&promoData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	promoStart, err := time.Parse("2006-01-02 15:04:05", promoData.PromoTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid promo_time format, use YYYY-MM-DD HH:MM:SS"})
		return
	}

	promoEnd, err := time.Parse("2006-01-02 15:04:05", promoData.PromoEnds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid promo_ends format, use YYYY-MM-DD HH:MM:SS"})
		return
	}

	err = h.ScheduleUC.ApplyPromo(uint(id), promoData.Promo, promoStart, promoEnd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to apply promo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Promo applied successfully"})
}
