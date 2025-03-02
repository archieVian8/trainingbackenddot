package usecase

import (
	"trainingbackenddot/domain"
	"trainingbackenddot/infrastructure/db"
)

type FilmUsecase struct {
	FilmRepo *db.FilmRepository
}

func NewFilmUsecase(filmRepo *db.FilmRepository) *FilmUsecase {
	return &FilmUsecase{FilmRepo: filmRepo}
}

// Function for add new film
func (u *FilmUsecase) AddFilm(film *domain.Film) error {
	return u.FilmRepo.CreateFilm(film)
}

// Function for get all films
func (u *FilmUsecase) GetAllFilms() ([]domain.Film, error) {
	return u.FilmRepo.GetAllFilms()
}

// Function for update film
func (u *FilmUsecase) UpdateFilm(id uint, updatedFilm *domain.Film) error {
	return u.FilmRepo.UpdateFilm(id, updatedFilm)
}

// Function for remove film
func (u *FilmUsecase) DeleteFilm(id uint) error {
	return u.FilmRepo.DeleteFilm(id)
}
