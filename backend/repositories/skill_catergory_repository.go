package repositories

import (
	"backend/models"
	"backend/models/shared"
	"database/sql"
	"log"
	"time"
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

func GetSkillCategories(db *sql.DB) ([]models.SkillCategory, error) {
	sqlStr := "SELECT * FROM skill_categories"

	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]models.SkillCategory, 0)
	for rows.Next() {
		var category models.SkillCategory
		var createdAt, updatedAt time.Time
		err := rows.Scan(&category.ID, &category.Name, &createdAt, &updatedAt)
		if err != nil {
			log.Printf("failed to scan: %v", err.Error())
			return nil, err
		}

		category.CreatedAt = shared.JstTime{Time: createdAt}
		category.UpdatedAt = shared.JstTime{Time: updatedAt}
		categories = append(categories, category)
	}

	return categories, nil
}
