package services

import (
	"backend/models"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type AuthService struct{}

func NewAuthService() AuthService {
	return AuthService{}
}

func (as *AuthService) CreateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
		"iat": time.Now().Unix(),
		"iss": "skill-tracker",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
