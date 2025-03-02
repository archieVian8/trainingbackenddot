package db

import (
	"trainingbackenddot/domain"

	"gorm.io/gorm"
)

type FilmRepository struct {
	DB *gorm.DB
}

func NewFilmRepository(db *gorm.DB) *FilmRepository {
	return &FilmRepository{DB: db}
}

// Create new Film
func (r *FilmRepository) CreateFilm(film *domain.Film) error {
	return r.DB.Create(film).Error
}

// Get All Films
func (r *FilmRepository) GetAllFilms() ([]domain.Film, error) {
	var films []domain.Film
	err := r.DB.Find(&films).Error
	return films, err
}

// Update Film
func (r *FilmRepository) UpdateFilm(id uint, updatedFilm *domain.Film) error {
	return r.DB.Model(&domain.Film{}).Where("id = ?", id).Updates(updatedFilm).Error
}

// Delete Film
func (r *FilmRepository) DeleteFilm(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&domain.Film{}).Error
}
