package repositories

import (
	"backend/models"
	"backend/models/shared"
	"database/sql"
	"log"
	"time"
)

func InsertSkill(tx *sql.Tx, db *sql.DB, skill models.Skill) (models.Skill, error) {
	sqlStr := "INSERT INTO skills (name, skill_category_id, description) VALUES ($1, $2, $3) RETURNING id"

	var id int
	err := tx.QueryRow(sqlStr, skill.Name, skill.SkillCategoryID, skill.Description).Scan(&id)
	if err != nil {
		return models.Skill{}, err
	}

	skill.ID = id

	return skill, nil
}

func GetAllSkills(db *sql.DB) ([]models.Skill, error) {
	sqlStr := "SELECT * FROM skills"

	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skills []models.Skill
	for rows.Next() {
		var skill models.Skill
		var CreatedAt, UpdatedAt time.Time
		err := rows.Scan(&skill.ID, &skill.SkillCategoryID, &skill.Name, &skill.Description, &CreatedAt, &UpdatedAt)
		if err != nil {
			return nil, err
		}

		skill.CreatedAt = shared.JstTime{Time: CreatedAt}
		skill.UpdatedAt = shared.JstTime{Time: UpdatedAt}

		skills = append(skills, skill)
	}

	return skills, nil
}

func UpSertUserSkill(tx *sql.Tx, db *sql.DB, userSkill models.UserSkill) (models.UserSkill, error) {
	log.Printf("userSkill: %v", userSkill)
	sqlStr := "INSERT INTO user_skills (user_id, skill_id, level, interested) VALUES ($1, $2, $3, $4) ON CONFLICT (user_id, skill_id) DO UPDATE SET level = $3, interested = $4 RETURNING id"

	var id int
	err := tx.QueryRow(sqlStr, userSkill.UserID, userSkill.SkillID, userSkill.Level, userSkill.Interested).Scan(&id)
	if err != nil {
		return models.UserSkill{}, err
	}

	userSkill.ID = id

	return userSkill, nil
}
