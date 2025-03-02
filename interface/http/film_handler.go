package http

import (
	"net/http"
	"strconv"
	"trainingbackenddot/domain"
	"trainingbackenddot/usecase"

	"github.com/gin-gonic/gin"
)

type FilmHandler struct {
	FilmUsecase *usecase.FilmUsecase
}

func NewFilmHandler(filmUsecase *usecase.FilmUsecase) *FilmHandler {
	return &FilmHandler{FilmUsecase: filmUsecase}
}

// Endpoint create new film
func (h *FilmHandler) AddFilm(c *gin.Context) {
	var film domain.Film
	if err := c.ShouldBindJSON(&film); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.FilmUsecase.AddFilm(&film); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add film"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Film added successfully"})
}

// Endpoint view all films
func (h *FilmHandler) GetAllFilms(c *gin.Context) {
	films, err := h.FilmUsecase.GetAllFilms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve films"})
		return
	}

	c.JSON(http.StatusOK, films)
}

// Endpoint update film
func (h *FilmHandler) UpdateFilm(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var film domain.Film
	if err := c.ShouldBindJSON(&film); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.FilmUsecase.UpdateFilm(uint(id), &film); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update film"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Film updated successfully"})
}

// Endpoint Delete Film
func (h *FilmHandler) DeleteFilm(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.FilmUsecase.DeleteFilm(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete film"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Film deleted successfully"})
}
