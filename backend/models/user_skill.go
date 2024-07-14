package models

import "backend/models/shared"

type UserSkill struct {
	ID         int            `json:"id"`
	UserID     int            `json:"user_id"`
	SkillID    int            `json:"skill_id"`
	Level      int            `json:"level"`
	Interested bool           `json:"interested"`
	CreatedAt  shared.JstTime `json:"created_at"`
	UpdatedAt  shared.JstTime `json:"updated_at"`
}
