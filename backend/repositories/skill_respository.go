package repositories

import (
	"backend/models"
	"database/sql"
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
