package controllers

import (
	"backend/apperrors"
	"backend/containers"
	"backend/models"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
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
		apperrors.ErrorHandler(w, r, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password))

	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := c.services.Auth.CreateToken(user)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookieSecure, err := strconv.ParseBool(os.Getenv("COOKIE_SECURE"))
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(48 * time.Hour),
		HttpOnly: true,
		Secure:   cookieSecure, // 開発環境ではfalse、運用環境ではtrue
		Path:     "/",
	}

	http.SetCookie(w, &cookie)

	w.Write([]byte("Login successful"))
}

func (c *AuthController) Me(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	tokenString := cookie.Value
	claims, err := c.services.Auth.VerifyToken(tokenString)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	id, ok := claims["id"].(float64)
	if !ok {
		apperrors.ErrorHandler(w, r, err)
		http.Error(w, "Invalid token claims", http.StatusBadRequest)
		return
	}

	userID := int(id)

	user, err := c.services.User.FindByID(userID)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filteredUser := models.FilteredUser{
		ID:        user.ID,
		Email:     user.Email,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	json.NewEncoder(w).Encode(filteredUser)
}
