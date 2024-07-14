package controllers

import (
	"backend/apperrors"
	"backend/containers"
	"backend/models/requests"
	"backend/models/responses"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func (c *SkillController) GetSkillCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := c.services.Skill.GetCategories()
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

func (c *SkillController) GetAllSkills(w http.ResponseWriter, r *http.Request) {
	skills, err := c.services.Skill.GetAllSkills()
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(skills)
}

func (c *SkillController) GetUserSkills(w http.ResponseWriter, r *http.Request) {
	strUserId := chi.URLParam(r, "user_id")

	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "user_id must be number")
		apperrors.ErrorHandler(w, r, err)
		return
	}

	skills, err := c.services.Skill.GetUserSkills(userId)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(skills)
}

func (c *SkillController) PostUserSkill(w http.ResponseWriter, r *http.Request) {
	request := requests.PostUserSkillRequest{}
	json.NewDecoder(r.Body).Decode(&request)

	err := request.Validate()
	if err != nil {
		err = apperrors.BadParam.Wrap(err, err.Error())
		apperrors.ErrorHandler(w, r, err)
		return
	}

	err = c.services.Skill.UpdateUserSkill(request)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
