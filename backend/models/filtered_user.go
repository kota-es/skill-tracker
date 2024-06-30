package models

import "backend/models/shared"

type FilteredUser struct {
	ID        int            `json:"id"`
	Email     string         `json:"email"`
	LastName  string         `json:"lastname"`
	FirstName string         `json:"firstname"`
	Role      string         `json:"role"`
	CreatedAt shared.JstTime `json:"created_at"`
	UpdatedAt shared.JstTime `json:"updated_at"`
}
