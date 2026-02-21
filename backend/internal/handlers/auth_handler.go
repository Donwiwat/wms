package handlers

import (
	"net/http"
	"os"
	"time"

	"wms-backend/internal/models"
	"wms-backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	authService services.AuthService
	validator   *validator.Validate
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   validator.New(),
	}
}

// Login handles user login
// @Summary User login
// @Description Authenticate user credentials and return JWT access token for API access
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "User login credentials"
// @Success 200 {object} models.AuthResponse "Successful authentication"
// @Failure 400 {object} models.ErrorResponse "Invalid request format"
// @Failure 401 {object} models.ErrorResponse "Invalid credentials"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authenticate user
	authResponse, err := h.authService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT token
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "your-super-secret-jwt-key-here"
	}

	expiresIn := 24 * time.Hour
	token, err := h.authService.GenerateToken(&authResponse.User, secretKey, expiresIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	authResponse.Token = token
	authResponse.User.Password = "" // Don't return password

	c.JSON(http.StatusOK, authResponse)
}

// Register handles user registration
// @Summary User registration
// @Description Create a new user account in the system with automatic JWT token generation
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.RegisterRequest true "User registration data"
// @Success 201 {object} models.AuthResponse "User created successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid request data"
// @Failure 409 {object} models.ErrorResponse "User already exists"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Register user
	authResponse, err := h.authService.Register(&req)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "username already exists" || err.Error() == "email already exists" {
			status = http.StatusConflict
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT token
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "your-super-secret-jwt-key-here"
	}

	expiresIn := 24 * time.Hour
	token, err := h.authService.GenerateToken(&authResponse.User, secretKey, expiresIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	authResponse.Token = token

	c.JSON(http.StatusCreated, authResponse)
}
