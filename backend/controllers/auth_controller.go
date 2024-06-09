package controllers

import (
	"backend/containers"
	"backend/models"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	services *containers.ServiceContainer
}

func NewAuthController(container *containers.ServiceContainer) *AuthController {
	return &AuthController{container}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	reqUser := models.User{}
	json.NewDecoder(r.Body).Decode(&reqUser)

	user, err := c.services.User.FindByEmail(reqUser.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password))

	if err != nil {
		http.Error(w, "パスワードが違います", http.StatusBadRequest)
		return
	}

	token, err := c.services.Auth.CreateToken(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(48 * time.Hour),
		HttpOnly: true,
		Secure:   true, // 開発環境ではfalse、運用環境ではtrue
		Path:     "/",
	}

	http.SetCookie(w, &cookie)

	w.Write([]byte("Login successful"))
}
