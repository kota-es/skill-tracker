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

func GetSkillLevelsBySkillID(db *sql.DB, skillID int) ([]models.SkillLevel, error) {
	sqlStr := "SELECT * FROM skill_levels WHERE skill_id = $1"

	rows, err := db.Query(sqlStr, skillID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skillLevels []models.SkillLevel
	for rows.Next() {
		var skillLevel models.SkillLevel
		err := rows.Scan(&skillLevel.ID, &skillLevel.SkillID, &skillLevel.Level, &skillLevel.Explanation)
		if err != nil {
			return nil, err
		}

		skillLevels = append(skillLevels, skillLevel)
	}

	return skillLevels, nil
}
