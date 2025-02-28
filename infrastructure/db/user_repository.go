package db

import (
	"trainingbackenddot/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Create new Admin
func (r *UserRepository) CreateUser(user *domain.User) error {
	return r.DB.Create(user).Error
}

// Get Admin by email
func (r *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Get View All User
func (r *UserRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	err := r.DB.Find(&users).Error
	return users, err
}
