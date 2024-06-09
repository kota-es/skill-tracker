package repositories_test

import (
	"fmt"
	"testing"
	"time"

	"backend/models"
	"backend/repositories"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsertUser(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := models.User{
		Email:     "test@example.com",
		Password:  "password123",
		LastName:  "Doe",
		FirstName: "John",
		Role:      "user",
	}

	now := time.Now()

	mock.ExpectQuery("INSERT INTO users \\(email, password, lastname, firstname, role\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\) returning id, created_at, updated_at").
		WithArgs(user.Email, sqlmock.AnyArg(), user.LastName, user.FirstName, user.Role).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
			AddRow(1, now, now))

	result, err := repositories.InsertUser(db, user)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if result.ID != 1 {
		t.Errorf("expected ID to be 1, but got %d", result.ID)
	}
	if !result.CreatedAt.Time.Equal(now) {
		t.Errorf("expected CreatedAt to be %v, but got %v", now, result.CreatedAt)
	}
	if !result.UpdatedAt.Time.Equal(now) {
		t.Errorf("expected UpdatedAt to be %v, but got %v", now, result.UpdatedAt)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertUser_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := models.User{
		Email:     "test@example.com",
		Password:  "password123",
		LastName:  "Doe",
		FirstName: "John",
		Role:      "user",
	}

	// クエリ実行時のエラーを設定
	mock.ExpectQuery("INSERT INTO users \\(email, password, lastname, firstname, role\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\) returning id, created_at, updated_at").
		WithArgs(user.Email, sqlmock.AnyArg(), user.LastName, user.FirstName, user.Role).
		WillReturnError(fmt.Errorf("some query error"))

	result, err := repositories.InsertUser(db, user)
	if err == nil {
		t.Fatalf("expected an error but got none")
	}
	if result.ID != 0 {
		t.Errorf("expected ID to be 0, but got %d", result.ID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertUser_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	user := models.User{
		Email:     "test@example.com",
		Password:  "password123",
		LastName:  "Doe",
		FirstName: "John",
		Role:      "user",
	}

	// スキャン時のエラーをシミュレートするため、不正なデータ型を返す
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
		AddRow("invalid_id", "invalid_created_at", "invalid_updated_at")

	mock.ExpectQuery("INSERT INTO users \\(email, password, lastname, firstname, role\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\) returning id, created_at, updated_at").
		WithArgs(user.Email, sqlmock.AnyArg(), user.LastName, user.FirstName, user.Role).
		WillReturnRows(rows)

	result, err := repositories.InsertUser(db, user)
	if err == nil {
		t.Fatalf("expected an error but got none")
	}
	if result.ID != 0 {
		t.Errorf("expected ID to be 0, but got %d", result.ID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
