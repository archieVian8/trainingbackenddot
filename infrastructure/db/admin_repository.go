package db

import (
	"trainingbackenddot/domain"

	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{DB: db}
}

// Create new Admin
func (r *AdminRepository) CreateAdmin(admin *domain.Admin) error {
	return r.DB.Create(admin).Error
}

// Get Admin by email
func (r *AdminRepository) GetAdminByEmail(email string) (*domain.Admin, error) {
	var admin domain.Admin
	err := r.DB.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

// Get View All Admin
func (r *AdminRepository) GetAllAdmins() ([]domain.Admin, error) {
	var admins []domain.Admin
	err := r.DB.Find(&admins).Error
	return admins, err
}
