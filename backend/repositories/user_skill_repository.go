package repositories

import (
	"backend/models"
	"backend/models/shared"
	"database/sql"
	"time"
)

func GetUserSkills(db *sql.DB, UserID int) ([]models.UserSkill, error) {
	sqlStr := "SELECT * FROM user_skills WHERE user_id = $1"

	rows, err := db.Query(sqlStr, UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userSkills []models.UserSkill
	for rows.Next() {
		var userSkill models.UserSkill
		var CreatedAt, UpdatedAt time.Time
		err := rows.Scan(&userSkill.ID, &userSkill.UserID, &userSkill.SkillID, &userSkill.Level, &userSkill.Interested, &CreatedAt, &UpdatedAt)
		if err != nil {
			return nil, err
		}

		userSkill.CreatedAt = shared.JstTime{Time: CreatedAt}
		userSkill.UpdatedAt = shared.JstTime{Time: UpdatedAt}

		userSkills = append(userSkills, userSkill)
	}

	return userSkills, nil
}
