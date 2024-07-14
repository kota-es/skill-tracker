package services

import (
	"backend/apperrors"
	"backend/models"
	"backend/models/requests"
	"backend/repositories"
	"database/sql"
	"log"
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

func (ss *SkillService) GetCategories() ([]models.SkillCategory, error) {
	categories, err := repositories.GetSkillCategories(ss.db)
	if err != nil {
		log.Printf("Failed to fetch skill categories: %v", err.Error())
		err = apperrors.GetDataFailed.Wrap(err, "failed to fetch skill categories")
		return nil, err
	}

	return categories, nil
}

func (ss *SkillService) GetAllSkills() ([]models.Skill, error) {
	skills, err := repositories.GetAllSkills(ss.db)
	if err != nil {
		log.Printf("Failed to fetch skills: %v", err.Error())
		err = apperrors.GetDataFailed.Wrap(err, "failed to fetch skills")
		return nil, err
	}

	return skills, nil
}

func (ss *SkillService) GetUserSkills(userID int) ([]models.UserSkill, error) {
	userSkills, err := repositories.GetUserSkills(ss.db, userID)
	if err != nil {
		log.Printf("Failed to fetch user skills: %v", err.Error())
		err = apperrors.GetDataFailed.Wrap(err, "failed to fetch user skills")
		return nil, err
	}

	return userSkills, nil
}

func (ss *SkillService) UpdateUserSkill(skillRequest requests.PostUserSkillRequest) error {
	tx, err := ss.db.Begin()
	if err != nil {
		apperrors.TransactinoFailed.Wrap(err, "failed to start transaction")
		log.Printf("Failed to start transaction: %v", err.Error())
		return err
	}
	defer tx.Rollback()

	for _, skill := range skillRequest.Skills {
		userSkill := models.UserSkill{
			UserID:     skillRequest.UserID,
			SkillID:    skill.SkillID,
			Level:      skill.Level,
			Interested: skill.Interested,
		}

		_, err := repositories.UpSertUserSkill(tx, ss.db, userSkill)
		if err != nil {
			apperrors.InsertDataFailed.Wrap(err, "failed to update user skill")
			log.Printf("Failed to update user skill: %v", err.Error())
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		apperrors.TransactinoFailed.Wrap(err, "failed to commit transaction")
		log.Printf("Failed to commit transaction: %v", err.Error())
		return err
	}

	return nil
}
