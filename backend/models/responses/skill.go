package responses

import (
	models "backend/models"
)

type SkillResponse struct {
	models.Skill
	Levels []SkillLevel `json:"levels"`
}

type SkillLevel struct {
	Level       int    `json:"level"`
	Explanation string `json:"explanation"`
}
