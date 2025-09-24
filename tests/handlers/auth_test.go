package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"golang.org/x/crypto/bcrypt"

	"github.com/arthu/shop-api-go/internal/config"
	"github.com/arthu/shop-api-go/tests/testutils"
)

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TestLogin_Success(t *testing.T) {
	cfg := &config.Config{JWTSecret: "secret"}
	db, mock, err := sqlmock.New()
	if err != nil { t.Fatalf("sqlmock: %v", err) }
	r := testutils.NewServer(cfg, db)

	pw := "mypassword"
	hash, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	rows := sqlmock.NewRows([]string{"id","username","email","password","firstName","lastName","address","phoneNumber","role"}).
		AddRow(1, "john", "john@example.com", string(hash), "John", "Doe", "123 St", "9999", "admin")
	mock.ExpectQuery("SELECT id, username, email, password").
		WithArgs("john@example.com").
		WillReturnRows(rows)

	body, _ := json.Marshal(loginReq{Email: "john@example.com", Password: pw})
	req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d, body=%s", w.Code, w.Body.String())
	}
	if err := mock.ExpectationsWereMet(); err != nil { t.Fatalf("sql expectations: %v", err) }
}

func TestLogin_Unauthorized(t *testing.T) {
	cfg := &config.Config{JWTSecret: "secret"}
	db, mock, err := sqlmock.New()
	if err != nil { t.Fatalf("sqlmock: %v", err) }
	r := testutils.NewServer(cfg, db)

	mock.ExpectQuery("SELECT id, username, email, password").
		WithArgs("missing@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	body, _ := json.Marshal(loginReq{Email: "missing@example.com", Password: "pw"})
	req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", w.Code)
	}
	if err := mock.ExpectationsWereMet(); err != nil { t.Fatalf("sql expectations: %v", err) }
}
