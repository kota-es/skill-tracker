package controllers

import (
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
	user, err := c.services.User.Create(reqUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
