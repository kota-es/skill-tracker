package models

import (
	models "backend/models/shared"
)

type Skill struct {
	ID              int                   `json:"id"`
	SkillCategoryID int                   `json:"skill_category_id"`
	Name            string                `json:"name"`
	Description     string                `json:"description"`
	Levels          []PhilteredSkillLevel `json:"levels"`
	CreatedAt       models.JstTime        `json:"created_at"`
	UpdatedAt       models.JstTime        `json:"updated_at"`
}

type PhilteredSkillLevel struct {
	Level       int    `json:"level"`
	Explanation string `json:"explanation"`
}
