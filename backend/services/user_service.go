package services

import (
	"backend/models"
	"backend/repositories"
	"database/sql"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) UserService {
	return UserService{db}
}

func (us *UserService) Create(user models.User) (models.User, error) {
	user, err := repositories.InsertUser(us.db, user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
