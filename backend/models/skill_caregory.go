package models

import models "backend/models/shared"

type SkillCategory struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt models.JstTime `json:"created_at"`
	UpdatedAt models.JstTime `json:"updated_at"`
}
