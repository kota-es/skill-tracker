package models

import (
	"backend/models/shared"
)

type SkillLevel struct {
	ID          int            `json:"id"`
	SkillID     int            `json:"skill_id"`
	Level       int            `json:"level"`
	Explanation string         `json:"explanation"`
	CreatedAt   shared.JstTime `json:"created_at"`
	UpdatedAt   shared.JstTime `json:"updated_at"`
}
