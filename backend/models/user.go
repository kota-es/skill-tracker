package models

import (
	"backend/models/shared"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	ID            int            `json:"id"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	LastName      string         `json:"lastname"`
	FirstName     string         `json:"firstname"`
	LastNameKana  string         `json:"lastname_kana"`
	FirstNameKana string         `json:"firstname_kana"`
	Role          string         `json:"role"`
	CreatedAt     shared.JstTime `json:"created_at"`
	UpdatedAt     shared.JstTime `json:"updated_at"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(
			&u.Email,
			validation.Required.Error("メールアドレスは必須です"),
			is.Email.Error("メールアドレスの形式が正しくありません"),
		),
		validation.Field(
			&u.Password,
			validation.Required.Error("パスワードは必須です"),
		),
		validation.Field(
			&u.LastName,
			validation.Required.Error("姓は必須です"),
		),
		validation.Field(
			&u.FirstName,
			validation.Required.Error("名は必須です"),
		),
		validation.Field(
			&u.LastNameKana,
			validation.Required.Error("姓かなは必須です"),
		),
		validation.Field(
			&u.FirstNameKana,
			validation.Required.Error("名かなは必須です"),
		),
		validation.Field(
			&u.Role,
			validation.In("admin", "user").Error("ロールはadminかuserのいずれかである必要があります"),
		),
	)
}
