package controllers

import (
	"backend/apperrors"
	"backend/containers"
	"backend/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	services *containers.ServiceContainer
}

func NewUserController(container *containers.ServiceContainer) *UserController {
	return &UserController{container}
}

func (c *UserController) PostUser(w http.ResponseWriter, r *http.Request) {
	reqUser := models.User{}
	json.NewDecoder(r.Body).Decode(&reqUser)

	err := reqUser.Validate()
	if err != nil {
		err = apperrors.BadParam.Wrap(err, err.Error())
		apperrors.ErrorHandler(w, r, err)
		return
	}

	user, err := c.services.User.Create(reqUser)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	strUserId := chi.URLParam(r, "user_id")
	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "Invalid user_id")
		apperrors.ErrorHandler(w, r, err)
		return
	}

	user, err := c.services.User.FindByID(userId)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (c *UserController) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	strUserId := chi.URLParam(r, "user_id")
	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "Invalid user_id")
		apperrors.ErrorHandler(w, r, err)
		return
	}

	user, err := c.services.User.FindProfileByID(userId)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (c *UserController) PostUserProfile(w http.ResponseWriter, r *http.Request) {
	request := models.UserProfile{}
	json.NewDecoder(r.Body).Decode(&request)

	err := c.services.User.UpdateProfile(request)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
