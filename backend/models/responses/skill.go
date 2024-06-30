package responses

import (
	models "backend/models"
)

type SkillResponse struct {
	models.Skill
	LevelExplanation []LevelExplanation `json:"explanation"`
}

type LevelExplanation struct {
	Level       int    `json:"level"`
	Explanation string `json:"explanation"`
}
