package services

import (
	"backend/apperrors"
	"backend/models"
	"backend/repositories"
	"database/sql"
	"errors"
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
		err = apperrors.InsertDataFailed.Wrap(err, "failed to insert user")
		return models.User{}, err
	}

	return user, nil
}

func (us *UserService) FindByEmail(email string) (models.User, error) {
	user, err := repositories.FindUserByEmail(us.db, email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoData.Wrap(err, "user not found")
			return models.User{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "failed to get user")
		return models.User{}, err
	}

	return user, nil
}

func (us *UserService) FindByID(id int) (models.User, error) {
	user, err := repositories.FindUserByID(us.db, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoData.Wrap(err, "user not found")
			return models.User{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "failed to get user")
		return models.User{}, err
	}

	return user, nil
}
