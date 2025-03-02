package http

import (
	"net/http"
	"strconv"
	"trainingbackenddot/domain"
	"trainingbackenddot/usecase"

	"github.com/gin-gonic/gin"
)

type StudioHandler struct {
	StudioUC *usecase.StudioUsecase
}

func NewStudioHandler(studioUC *usecase.StudioUsecase) *StudioHandler {
	return &StudioHandler{StudioUC: studioUC}
}

// Endpoint Create Studio
func (h *StudioHandler) CreateStudio(c *gin.Context) {
	var studio domain.Studio
	if err := c.ShouldBindJSON(&studio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.StudioUC.CreateStudio(&studio); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Studio successfully added"})
}

// Endpoint View All Studio
func (h *StudioHandler) GetAllStudios(c *gin.Context) {
	studios, err := h.StudioUC.GetAllStudios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, studios)
}

// Endpoint Update Studio
func (h *StudioHandler) UpdateStudio(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedStudio domain.Studio
	if err := c.ShouldBindJSON(&updatedStudio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.StudioUC.UpdateStudio(uint(id), &updatedStudio); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Studio successfully updated"})
}

// Endpoint Delete Studio
func (h *StudioHandler) DeleteStudio(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.StudioUC.DeleteStudio(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Studio successfully deleted"})
}
