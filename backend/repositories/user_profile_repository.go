package repositories

import (
	"backend/models"
	"backend/models/shared"
	"database/sql"
	"time"
)

func FindProfileByUserID(db *sql.DB, userID int) (models.UserProfile, error) {
	sqlStr := "SELECT * FROM user_profiles WHERE user_id = $1"

	var UserProfile models.UserProfile
	var CreatedAt, UpdatedAt time.Time
	err := db.QueryRow(sqlStr, userID).Scan(&UserProfile.ID, &UserProfile.UserID, &UserProfile.Notes, &UserProfile.Desires, &UserProfile.Dislikes, &CreatedAt, &UpdatedAt)
	if err != nil {
		return models.UserProfile{}, err
	}

	UserProfile.CreatedAt = shared.JstTime{Time: CreatedAt}
	UserProfile.UpdatedAt = shared.JstTime{Time: UpdatedAt}

	return UserProfile, nil
}

func UpsertUserProfile(db *sql.DB, userProfile models.UserProfile) (models.UserProfile, error) {
	sqlStr := "INSERT INTO user_profiles (user_id, notes, desires, dislikes) VALUES ($1, $2, $3, $4) ON CONFLICT (user_id) DO UPDATE SET notes = $2, desires = $3, dislikes = $4, updated_at = now() RETURNING id"

	var id int
	err := db.QueryRow(sqlStr, userProfile.UserID, userProfile.Notes, userProfile.Desires, userProfile.Dislikes).Scan(&id)
	if err != nil {
		return models.UserProfile{}, err
	}

	userProfile.ID = id

	return userProfile, nil
}
