package db

import (
	"trainingbackenddot/domain"

	"gorm.io/gorm"
)

type StudioRepository struct {
	DB *gorm.DB
}

func NewStudioRepository(db *gorm.DB) *StudioRepository {
	return &StudioRepository{DB: db}
}

// Create Studio
func (r *StudioRepository) CreateStudio(studio *domain.Studio) error {
	return r.DB.Create(studio).Error
}

// View All Studio
func (r *StudioRepository) GetAllStudios() ([]domain.Studio, error) {
	var studios []domain.Studio
	err := r.DB.Find(&studios).Error
	return studios, err
}

// Update Studio
func (r *StudioRepository) UpdateStudio(id uint, updatedStudio *domain.Studio) error {
	return r.DB.Model(&domain.Studio{}).Where("id = ?", id).Updates(updatedStudio).Error
}

// Delete Studio
func (r *StudioRepository) DeleteStudio(id uint) error {
	return r.DB.Delete(&domain.Studio{}, id).Error
}
