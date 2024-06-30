package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type NewSkillRequest struct {
	Name              string             `json:"name"`
	Description       string             `json:"description"`
	LevelExplanation  []LevelExplanation `json:"level_explanation"`
	IsNewCategory     bool               `json:"is_new_category"`
	SkillCategoryID   int                `json:"skill_category_id"`
	SkillCategoryName string             `json:"skill_category_name"`
}

type LevelExplanation struct {
	Level       int    `json:"level"`
	Explanation string `json:"explanation"`
}

func (r NewSkillRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Name,
			validation.Required.Error("スキル名は必須です"),
		),
		validation.Field(
			&r.Description,
			validation.Required.Error("スキル概要は必須です"),
		),
	)
}

func (r LevelExplanation) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.Level,
			validation.Required.Error("スキルレベルは必須です"),
		),
		validation.Field(
			&r.Explanation,
			validation.Required.Error("レベルの説明は必須です"),
		),
	)
}
