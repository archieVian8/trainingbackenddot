package usecase

import (
	"trainingbackenddot/domain"
	"trainingbackenddot/infrastructure/db"
)

type StudioUsecase struct {
	StudioRepo *db.StudioRepository
}

func NewStudioUsecase(studioRepo *db.StudioRepository) *StudioUsecase {
	return &StudioUsecase{StudioRepo: studioRepo}
}

// Function for Create Studio
func (u *StudioUsecase) CreateStudio(studio *domain.Studio) error {
	return u.StudioRepo.CreateStudio(studio)
}

// Function for View All Studio
func (u *StudioUsecase) GetAllStudios() ([]domain.Studio, error) {
	return u.StudioRepo.GetAllStudios()
}

// Function for Update Studio
func (u *StudioUsecase) UpdateStudio(id uint, updatedStudio *domain.Studio) error {
	return u.StudioRepo.UpdateStudio(id, updatedStudio)
}

// Function for Delete Studio
func (u *StudioUsecase) DeleteStudio(id uint) error {
	return u.StudioRepo.DeleteStudio(id)
}
