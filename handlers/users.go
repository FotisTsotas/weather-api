package handlers

import (
	"errors"
	"net/http"
	"weather-api/models"
	"weather-api/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	authService *services.AuthService
}

func NewUserHandler(authService *services.AuthService) *UserHandler {
	return &UserHandler{authService: authService}
}

func (h *UserHandler) Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.authService.Signup(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save user"})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, token, err := h.authService.Login(credentials.Email, credentials.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}
