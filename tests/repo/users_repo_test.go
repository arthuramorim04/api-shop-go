package repo_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arthu/shop-api-go/internal/database"
	"github.com/arthu/shop-api-go/internal/models"
	"github.com/arthu/shop-api-go/internal/repo"
)

func TestUsers_CreateUser_OK(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil { t.Fatalf("sqlmock: %v", err) }
	database.SetDB(db)

	u := &models.User{Username: "john", Email: "john@example.com", Password: "hash", FirstName: "John", LastName: "Doe", Address: "123", PhoneNumber: "9999", Role: models.RoleAdmin}
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (username, email, password, firstName, lastName, address, phoneNumber, role) VALUES (?,?,?,?,?,?,?,?)")).
		WithArgs(u.Username, u.Email, u.Password, u.FirstName, u.LastName, u.Address, u.PhoneNumber, u.Role).
		WillReturnResult(sqlmock.NewResult(10, 1))

	id, err := repo.CreateUser(u)
	if err != nil { t.Fatalf("CreateUser err: %v", err) }
	if id != 10 { t.Fatalf("expected id=10, got %d", id) }
	if err := mock.ExpectationsWereMet(); err != nil { t.Fatalf("sql expectations: %v", err) }
}

func TestUsers_GetUserByEmail_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil { t.Fatalf("sqlmock: %v", err) }
	database.SetDB(db)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, username, email, password, firstName, lastName, address, phoneNumber, role FROM users WHERE email = ?")).
		WithArgs("missing@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	u, err := repo.GetUserByEmail("missing@example.com")
	if err != nil { t.Fatalf("GetUserByEmail err: %v", err) }
	if u != nil { t.Fatalf("expected nil user when not found") }
	if err := mock.ExpectationsWereMet(); err != nil { t.Fatalf("sql expectations: %v", err) }
}
