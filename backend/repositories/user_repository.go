package repositories

import (
	"backend/models"
	"backend/models/shared"

	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func InsertUser(db *sql.DB, user models.User) (models.User, error) {
	sqlStr := "INSERT INTO users (email, password, lastname, firstname, lastname_kana, firstname_kana,role) VALUES ($1, $2, $3, $4, $5, $6, $7) returning id, created_at, updated_at"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	row := db.QueryRow(sqlStr, user.Email, hashedPassword, user.LastName, user.FirstName, user.LastNameKana, user.FirstNameKana, user.Role)
	if row.Err() != nil {
		return models.User{}, row.Err()
	}

	var userID int
	var createdAt, updatedAt time.Time
	err = row.Scan(&userID, &createdAt, &updatedAt)
	if err != nil {
		return models.User{}, err
	}

	user.ID = userID
	user.CreatedAt = shared.JstTime{Time: createdAt}
	user.UpdatedAt = shared.JstTime{Time: updatedAt}

	return user, nil
}

func FindUserByEmail(db *sql.DB, email string) (models.User, error) {
	sqlStr := "SELECT * FROM users WHERE email = $1"

	row := db.QueryRow(sqlStr, email)
	if row.Err() != nil {
		return models.User{}, row.Err()
	}

	var user models.User
	var createdAt, updatedAt time.Time
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.LastName, &user.FirstName, &user.LastNameKana, &user.FirstNameKana, &user.Role, &createdAt, &updatedAt)
	if err != nil {
		return models.User{}, err
	}

	user.CreatedAt = shared.JstTime{Time: createdAt}
	user.UpdatedAt = shared.JstTime{Time: updatedAt}

	return user, nil
}

func FindUserByID(db *sql.DB, id int) (models.User, error) {
	sqlStr := "SELECT * FROM users WHERE id = $1"

	row := db.QueryRow(sqlStr, id)
	if row.Err() != nil {
		return models.User{}, row.Err()
	}

	var user models.User
	var createdAt, updatedAt time.Time
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.LastName, &user.FirstName, &user.LastNameKana, &user.FirstNameKana, &user.Role, &createdAt, &updatedAt)
	if err != nil {
		return models.User{}, err
	}

	user.CreatedAt = shared.JstTime{Time: createdAt}
	user.UpdatedAt = shared.JstTime{Time: updatedAt}

	return user, nil
}

func SearchUsers(db *sql.DB) ([]models.SearchedUser, error) {
	sqlStr := `
			SELECT 
					users.id, users.firstname, users.lastname, users.firstname_kana, users.lastname_kana, 
					user_profiles.id, user_profiles.notes, user_profiles.desires, user_profiles.dislikes, 
					user_profiles.created_at, user_profiles.updated_at,
					user_skills.id, user_skills.skill_id, user_skills.level, user_skills.interested, 
					user_skills.created_at, user_skills.updated_at
			FROM users 
			LEFT JOIN user_profiles ON users.id = user_profiles.user_id 
			LEFT JOIN user_skills ON users.id = user_skills.user_id
	`

	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userMap := make(map[int]*models.SearchedUser)

	for rows.Next() {
		var (
			userID           int
			firstname        string
			lastname         string
			firstnameKana    string
			lastnameKana     string
			profileID        sql.NullInt64
			notes            sql.NullString
			desires          sql.NullString
			dislikes         sql.NullString
			profileCreatedAt sql.NullTime
			profileUpdatedAt sql.NullTime
			skillID          sql.NullInt64
			skillSkillID     sql.NullInt64
			level            sql.NullInt64
			interested       sql.NullBool
			skillCreatedAt   sql.NullTime
			skillUpdatedAt   sql.NullTime
		)

		err := rows.Scan(
			&userID, &firstname, &lastname, &firstnameKana, &lastnameKana,
			&profileID, &notes, &desires, &dislikes, &profileCreatedAt, &profileUpdatedAt,
			&skillID, &skillSkillID, &level, &interested, &skillCreatedAt, &skillUpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		user, exists := userMap[userID]
		if !exists {
			user = &models.SearchedUser{
				UserID:        userID,
				Firstname:     firstname,
				Lastname:      lastname,
				FirstnameKana: firstnameKana,
				LastnameKana:  lastnameKana,
				Profile: models.UserProfile{
					ID:        int(profileID.Int64),
					UserID:    userID,
					Notes:     notes.String,
					Desires:   desires.String,
					Dislikes:  dislikes.String,
					CreatedAt: shared.JstTime{Time: profileCreatedAt.Time},
					UpdatedAt: shared.JstTime{Time: profileUpdatedAt.Time},
				},
				Skills: []models.UserSkill{},
			}
			userMap[userID] = user
		}

		if skillID.Valid {
			skill := models.UserSkill{
				ID:         int(skillID.Int64),
				UserID:     userID,
				SkillID:    int(skillSkillID.Int64),
				Level:      int(level.Int64),
				Interested: interested.Bool,
				CreatedAt:  shared.JstTime{Time: skillCreatedAt.Time},
				UpdatedAt:  shared.JstTime{Time: skillUpdatedAt.Time},
			}
			user.Skills = append(user.Skills, skill)
		}
	}

	result := make([]models.SearchedUser, 0, len(userMap))
	for _, user := range userMap {
		result = append(result, *user)
	}

	return result, nil
}
