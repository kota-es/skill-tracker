package models

import (
	"backend/models/shared"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UserProfile struct {
	ID        int            `json:"id"`
	UserID    int            `json:"user_id"`
	Notes     string         `json:"notes"`
	Desires   string         `json:"desires"`
	Dislikes  string         `json:"dislikes"`
	CreatedAt shared.JstTime `json:"created_at"`
	UpdatedAt shared.JstTime `json:"updated_at"`
}

func (up UserProfile) Validate() error {
	return validation.ValidateStruct(&up,
		validation.Field(
			&up.UserID,
			validation.Required.Error("ユーザーIDは必須です"),
		),
	)
}
