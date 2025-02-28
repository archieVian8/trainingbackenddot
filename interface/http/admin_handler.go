package http

import (
	"net/http"
	"trainingbackenddot/domain"
	"trainingbackenddot/usecase"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	AdminUC *usecase.AdminUseCase
}

func NewAdminHandler(adminUC *usecase.AdminUseCase) *AdminHandler {
	return &AdminHandler{AdminUC: adminUC}
}

// Endpoint Signup Admin
func (h *AdminHandler) SignupAdmin(c *gin.Context) {
	var admin domain.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.AdminUC.SignupAdmin(&admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Admin successfully created"})
}

// Endpoint Signin Admin
func (h *AdminHandler) SigninAdmin(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin, err := h.AdminUC.SigninAdmin(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sign in successful", "admin": admin})
}

// Endpoint View All Admin
func (h *AdminHandler) ViewAllAdmins(c *gin.Context) {
	admins, err := h.AdminUC.GetAllAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"admins": admins})
}
