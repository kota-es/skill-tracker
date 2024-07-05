package controllers

import (
	"backend/apperrors"
	"backend/containers"
	"backend/models/requests"
	"backend/models/responses"
	"encoding/json"
	"net/http"
)

type SkillController struct {
	services *containers.ServiceContainer
}

func NewSkillController(container *containers.ServiceContainer) *SkillController {
	return &SkillController{container}
}

func (c *SkillController) PostSkill(w http.ResponseWriter, r *http.Request) {
	reqSkill := requests.NewSkillRequest{}
	json.NewDecoder(r.Body).Decode(&reqSkill)

	err := reqSkill.Validate()
	if err != nil {
		err = apperrors.BadParam.Wrap(err, err.Error())
		apperrors.ErrorHandler(w, r, err)
		return
	}

	skill, skillLevels, err := c.services.Skill.Create(reqSkill)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}

	var response responses.SkillResponse
	response.Skill = skill

	response.LevelExplanation = make([]responses.LevelExplanation, 0)
	for _, level := range skillLevels {
		var levelResponse responses.LevelExplanation
		levelResponse.Level = level.Level
		levelResponse.Explanation = level.Explanation

		response.LevelExplanation = append(response.LevelExplanation, levelResponse)
	}

	json.NewEncoder(w).Encode(response)

}
