package repositories

import (
	"backend/models"
	"database/sql"
)

func InsertSkillCategory(tx *sql.Tx, db *sql.DB, skillCategory models.SkillCategory) (models.SkillCategory, error) {
	sqlStr := "INSERT INTO skill_categories (name) VALUES ($1) RETURNING id"

	var id int
	err := tx.QueryRow(sqlStr, skillCategory.Name).Scan(&id)
	if err != nil {
		return models.SkillCategory{}, err
	}

	skillCategory.ID = id

	return skillCategory, nil
}
