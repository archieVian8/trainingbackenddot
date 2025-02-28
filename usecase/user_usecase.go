package usecase

import (
	"errors"
	"trainingbackenddot/domain"
	"trainingbackenddot/infrastructure/db"
)

type UserUseCase struct {
	UserRepo *db.UserRepository
}

func NewUserUseCase(userRepo *db.UserRepository) *UserUseCase {
	return &UserUseCase{UserRepo: userRepo}
}

// Signup User
func (uc *UserUseCase) SignupUser(user *domain.User) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return uc.UserRepo.CreateUser(user)
}

// Signin User
func (uc *UserUseCase) SigninUser(email, password string) (*domain.User, error) {
	user, err := uc.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("wrong email or password")
	}

	if !checkPasswordHash(password, user.Password) {
		return nil, errors.New("wrong email or password")
	}

	return user, nil
}

// Function for view all user
func (u *UserUseCase) GetAllUsers() ([]domain.User, error) {
	return u.UserRepo.GetAllUsers()
}
