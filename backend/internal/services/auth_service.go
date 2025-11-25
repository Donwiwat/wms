package services

import (
	"database/sql"
	"errors"
	"time"

	"wms-backend/internal/models"
	"wms-backend/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthService interface defines authentication service methods
type AuthService interface {
	Login(req *models.LoginRequest) (*models.AuthResponse, error)
	Register(req *models.RegisterRequest) (*models.AuthResponse, error)
	GenerateToken(user *models.User, secretKey string, expiresIn time.Duration) (string, error)
	ValidateToken(tokenString, secretKey string) (*jwt.Token, error)
}

// authService implements AuthService
type authService struct {
	userRepo repositories.UserRepository
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

// Login authenticates a user and returns a token
func (s *authService) Login(req *models.LoginRequest) (*models.AuthResponse, error) {
	// Get user by username
	user, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("account is disabled")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate token (we'll need to pass secret and expiry from config in handlers)
	// For now, return user without token - token generation will be handled in handlers
	return &models.AuthResponse{
		User: *user,
	}, nil
}

// Register creates a new user account
func (s *authService) Register(req *models.RegisterRequest) (*models.AuthResponse, error) {
	// Check if username already exists
	if _, err := s.userRepo.GetByUsername(req.Username); err == nil {
		return nil, errors.New("username already exists")
	}

	// Check if email already exists
	if _, err := s.userRepo.GetByEmail(req.Email); err == nil {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Set default role if not provided
	role := req.Role
	if role == "" {
		role = "user"
	}

	// Create user
	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     role,
		IsActive: true,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Return user without password
	user.Password = ""
	return &models.AuthResponse{
		User: *user,
	}, nil
}

// GenerateToken generates a JWT token for a user
func (s *authService) GenerateToken(user *models.User, secretKey string, expiresIn time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(expiresIn).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ValidateToken validates a JWT token
func (s *authService) ValidateToken(tokenString, secretKey string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretKey), nil
	})
}
