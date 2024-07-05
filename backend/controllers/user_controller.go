package controllers

import (
	"backend/apperrors"
	"backend/containers"
	"backend/models"
	"encoding/json"
	"net/http"
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
