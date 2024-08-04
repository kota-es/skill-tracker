package models

type SearchedUser struct {
	UserID        int         `json:"user_id"`
	Firstname     string      `json:"first_name"`
	Lastname      string      `json:"last_name"`
	FirstnameKana string      `json:"firstname_kana"`
	LastnameKana  string      `json:"lastname_kana"`
	Profile       UserProfile `json:"profile"`
	Skills        []UserSkill `json:"skills"`
}
