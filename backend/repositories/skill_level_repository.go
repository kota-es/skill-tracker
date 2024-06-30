package repositories

import (
	"backend/models"
	"database/sql"
)

func InsertSkillLevel(tx *sql.Tx, db *sql.DB, skillLevel models.SkillLevel) (models.SkillLevel, error) {
	sqlStr := "INSERT INTO skill_levels (skill_id, level, explanation) VALUES ($1, $2, $3) RETURNING id"

	var id int
	err := tx.QueryRow(sqlStr, skillLevel.SkillID, skillLevel.Level, skillLevel.Explanation).Scan(&id)
	if err != nil {
		return models.SkillLevel{}, err
	}

	skillLevel.ID = id

	return skillLevel, nil
}
