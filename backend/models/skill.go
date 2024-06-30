package models

import (
	models "backend/models/shared"
)

type Skill struct {
	ID              int            `json:"id"`
	SkillCategoryID int            `json:"skill_category_id"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	CreatedAt       models.JstTime `json:"created_at"`
	UpdatedAt       models.JstTime `json:"updated_at"`
}
