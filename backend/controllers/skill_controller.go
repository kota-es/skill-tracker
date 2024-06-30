package controllers

import (
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	skill, skillLevels, err := c.services.Skill.Create(reqSkill)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
