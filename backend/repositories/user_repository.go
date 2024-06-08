package repositories

import (
	"backend/models"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func InsertUser(db *sql.DB, user models.User) (models.User, error) {
	sqlStr := "INSERT INTO users (email, password, lastname, firstname, role) VALUES ($1, $2, $3, $4, $5) returning id, created_at, updated_at"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	row := db.QueryRow(sqlStr, user.Email, hashedPassword, user.LastName, user.FirstName, user.Role)
	if row.Err() != nil {
		return models.User{}, row.Err()
	}

	var userID int
	var createdAt, updatedAt time.Time
	err = row.Scan(&userID, &createdAt, &updatedAt)
	if err != nil {
		return models.User{}, err
	}

	user.CreatedAt = models.JstTime{Time: createdAt}
	user.UpdatedAt = models.JstTime{Time: updatedAt}

	return user, nil
}
