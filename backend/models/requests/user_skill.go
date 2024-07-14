package requests

import validation "github.com/go-ozzo/ozzo-validation"

type GetUserSkillRequest struct {
	UserID int `json:"user_id"`
}

type PostUserSkillRequest struct {
	UserID int             `json:"user_id"`
	Skills []PostUserSkill `json:"skills"`
}

type PostUserSkill struct {
	SkillID    int  `json:"skill_id"`
	Level      int  `json:"level"`
	Interested bool `json:"interested"`
}

func (r PostUserSkillRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.UserID,
			validation.Required.Error("ユーザーIDは必須です"),
		),
	)
}

func (r *PostUserSkill) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(
			&r.SkillID,
			validation.Required.Error("スキルIDは必須です"),
		),
		validation.Field(
			&r.Level,
			validation.Required.Error("スキルレベルは必須です"),
		),
		validation.Field(
			&r.Interested,
			validation.Required.Error("興味があるかどうかは必須です"),
		),
	)
}
