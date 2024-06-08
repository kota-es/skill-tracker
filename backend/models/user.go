package models

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
