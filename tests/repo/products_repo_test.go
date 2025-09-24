package repo_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arthu/shop-api-go/internal/database"
	"github.com/arthu/shop-api-go/internal/repo"
)

func TestProducts_ListProducts_OK(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil { t.Fatalf("sqlmock: %v", err) }
	database.SetDB(db)

	rows := sqlmock.NewRows([]string{"id","name","description","price","quantity","imageUrl"}).
		AddRow(1, "p1", "d1", 10.0, 5, "img1").
		AddRow(2, "p2", "d2", 20.0, 10, "img2")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, description, price, quantity, imageUrl FROM products")).
		WillReturnRows(rows)

	items, err := repo.ListProducts()
	if err != nil { t.Fatalf("ListProducts err: %v", err) }
	if len(items) != 2 { t.Fatalf("expected 2 items, got %d", len(items)) }
	if err := mock.ExpectationsWereMet(); err != nil { t.Fatalf("sql expectations: %v", err) }
}

func TestProducts_GetProductByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil { t.Fatalf("sqlmock: %v", err) }
	database.SetDB(db)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, description, price, quantity, imageUrl FROM products WHERE id = ?")).
		WithArgs(int64(99)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	p, err := repo.GetProductByID(99)
	if err != nil { t.Fatalf("GetProductByID err: %v", err) }
	if p != nil { t.Fatalf("expected nil product when not found") }
	if err := mock.ExpectationsWereMet(); err != nil { t.Fatalf("sql expectations: %v", err) }
}
