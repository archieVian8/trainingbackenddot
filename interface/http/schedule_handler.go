package http

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"trainingbackenddot/domain"
	"trainingbackenddot/usecase"

	"github.com/gin-gonic/gin"
)

type ScheduleHandler struct {
	ScheduleUC     *usecase.ScheduleUsecase
	NotificationUC *usecase.NotificationUsecase
}

func NewScheduleHandler(scheduleoUC *usecase.ScheduleUsecase, notificationUC *usecase.NotificationUsecase) *ScheduleHandler {
	return &ScheduleHandler{
		ScheduleUC:     scheduleoUC,
		NotificationUC: notificationUC,
	}
}

// Endpoint Create New Schedule
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

	scheduleWithFilm, err := h.ScheduleUC.GetScheduleByID(int(schedule.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve schedule data"})
		return
	}

	message := fmt.Sprintf("A new schedule has been made for the film %s", scheduleWithFilm.Film.Title)
	h.NotificationUC.SendNotification(message)

	c.JSON(http.StatusCreated, gin.H{"message": "Schedule created successfully"})
}

// Endpoint View All Schedules
func (h *ScheduleHandler) ViewAllSchedules(c *gin.Context) {
	schedules, err := h.ScheduleUC.ViewAllSchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedules"})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// Endpoint Update Schedule
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

// Endpoint Delete Schedule
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

// Endpoint Create New Promo
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

	schedule, err := h.ScheduleUC.GetScheduleByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data jadwal"})
		return
	}

	message := fmt.Sprintf(
		"New discount available!\nFilm: %s\nStudio: %s\nSchedule: %s\nPrice: Rp%.2f\nDiscount: %d%%\nPromo Price: Rp%.2f",
		schedule.Film.Title,
		schedule.Studio.Name,
		schedule.ShowTime,
		schedule.Price,
		schedule.Promo,
		schedule.PromoPrice,
	)

	// Send notification to user
	h.NotificationUC.SendNotification(message)

	c.JSON(http.StatusOK, gin.H{"message": "Promo applied successfully"})

}
