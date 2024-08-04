package models

type SearchedUser struct {
	UserID        int         `json:"user_id"`
	Firstname     string      `json:"firstname"`
	Lastname      string      `json:"lastname"`
	FirstnameKana string      `json:"firstname_kana"`
	LastnameKana  string      `json:"lastname_kana"`
	Profile       UserProfile `json:"profile"`
	Skills        []UserSkill `json:"skills"`
}
