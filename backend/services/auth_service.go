package services

import (
	"backend/apperrors"
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
		err = apperrors.CreateTokenFailed.Wrap(err, "failed to create token")
		return "", err
	}

	return accessToken, nil
}

func (as *AuthService) VerifyToken(accessToken string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		err = apperrors.VerifyTokenFailed.Wrap(err, "failed to verify token")
		return nil, err
	}

	return claims, nil
}
