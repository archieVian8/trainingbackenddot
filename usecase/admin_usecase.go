package usecase

import (
	"errors"
	"trainingbackenddot/domain"
	"trainingbackenddot/infrastructure/db"

	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase struct {
	AdminRepo *db.AdminRepository
}

func NewAdminUseCase(adminRepo *db.AdminRepository) *AdminUseCase {
	return &AdminUseCase{AdminRepo: adminRepo}
}

// Function for Signup Admin
func (uc *AdminUseCase) SignupAdmin(admin *domain.Admin) error {
	hashedPassword, err := hashPassword(admin.Password)
	if err != nil {
		return err
	}
	admin.Password = hashedPassword

	return uc.AdminRepo.CreateAdmin(admin)
}

// Function for Sigin Admin
func (uc *AdminUseCase) SigninAdmin(email, password string) (*domain.Admin, error) {
	admin, err := uc.AdminRepo.GetAdminByEmail(email)
	if err != nil {
		return nil, errors.New("wrong email or password")
	}

	if !checkPasswordHash(password, admin.Password) {
		return nil, errors.New("wrong email or password")
	}

	return admin, nil
}

// Function for password hash
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Function for password verification
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Function for view all admin
func (u *AdminUseCase) GetAllAdmins() ([]domain.Admin, error) {
	return u.AdminRepo.GetAllAdmins()
}
