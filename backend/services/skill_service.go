package services

import (
	"backend/apperrors"
	"backend/models"
	"backend/models/requests"
	"backend/repositories"
	"database/sql"
)

type SkillService struct {
	db *sql.DB
}

func NewSkillService(db *sql.DB) SkillService {
	return SkillService{db}
}

func (ss *SkillService) Create(reqSkill requests.NewSkillRequest) (models.Skill, []models.SkillLevel, error) {

	var skillLevels = make([]models.SkillLevel, 0)

	tx, err := ss.db.Begin()
	if err != nil {
		err = apperrors.TransactinoFailed.Wrap(err, "failed to start transaction")
		return models.Skill{}, skillLevels, err
	}
	defer tx.Rollback()

	if reqSkill.IsNewCategory {
		var skillCategory models.SkillCategory
		skillCategory.Name = reqSkill.SkillCategoryName
		newSkillCategory, err := repositories.InsertSkillCategory(tx, ss.db, skillCategory)
		if err != nil {
			return models.Skill{}, skillLevels, err
		}

		reqSkill.SkillCategoryID = newSkillCategory.ID
	}

	modelSkill := models.Skill{
		SkillCategoryID: reqSkill.SkillCategoryID,
		Name:            reqSkill.Name,
		Description:     reqSkill.Description,
	}

	newSkill, err := repositories.InsertSkill(tx, ss.db, modelSkill)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to insert skill")
		return models.Skill{}, skillLevels, err
	}

	for _, level := range reqSkill.LevelExplanation {
		skillLevel := models.SkillLevel{
			SkillID:     newSkill.ID,
			Level:       level.Level,
			Explanation: level.Explanation,
		}

		_, err = repositories.InsertSkillLevel(tx, ss.db, skillLevel)
		if err != nil {
			err = apperrors.InsertDataFailed.Wrap(err, "failed to insert skill level")
			return models.Skill{}, skillLevels, err
		}

		skillLevels = append(skillLevels, skillLevel)
	}

	err = tx.Commit()
	if err != nil {
		err = apperrors.TransactinoFailed.Wrap(err, "failed to commit transaction")
		return models.Skill{}, skillLevels, err
	}

	return newSkill, skillLevels, nil
}
