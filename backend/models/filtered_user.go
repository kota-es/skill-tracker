package models

type FilteredUser struct {
	ID        int     `json:"id"`
	Email     string  `json:"email"`
	LastName  string  `json:"lastname"`
	FirstName string  `json:"firstname"`
	Role      string  `json:"role"`
	CreatedAt JstTime `json:"created_at"`
	UpdatedAt JstTime `json:"updated_at"`
}
