package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	ID        int     `json:"id"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	LastName  string  `json:"lastname"`
	FirstName string  `json:"firstname"`
	Role      string  `json:"role"`
	CreatedAt JstTime `json:"created_at"`
	UpdatedAt JstTime `json:"updated_at"`
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
			&u.Role,
			validation.In("admin", "user").Error("ロールはadminかuserのいずれかである必要があります"),
		),
	)
}
