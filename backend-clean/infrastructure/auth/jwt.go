package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/ddd-todo/project-backend/application/service"
	"github.com/golang-jwt/jwt/v5"
)

type JwtAuthService struct{}

func NewAuthService() service.AuthService {
	return &JwtAuthService{}
}

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func (a *JwtAuthService) GenerateToken(userID, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(getJWTSecret()))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// Development default - should be set in production
		secret = "your-secret-key-change-this-in-production"
	}
	return secret
}
