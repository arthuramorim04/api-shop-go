package repo_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arthu/shop-api-go/internal/database"
	"github.com/arthu/shop-api-go/internal/models"
	"github.com/arthu/shop-api-go/internal/repo"
)

func TestOrders_InsertOrder_OK(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil { t.Fatalf("sqlmock: %v", err) }
	database.SetDB(db)

	o := &models.Order{ID: "ord-1", UserID: 1, Status: models.OrderPending}
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO orders (id, user_id, order_date, status, payment_url) VALUES (?,?,?,?,?)")).
		WithArgs(o.ID, o.UserID, o.OrderDate, o.Status, o.PaymentURL).
		WillReturnResult(sqlmock.NewResult(0, 1))

	if err := repo.InsertOrder(o); err != nil { t.Fatalf("InsertOrder err: %v", err) }
	if err := mock.ExpectationsWereMet(); err != nil { t.Fatalf("sql expectations: %v", err) }
}

func TestOrders_UpdateOrderStatus_OK(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil { t.Fatalf("sqlmock: %v", err) }
	database.SetDB(db)

	mock.ExpectExec(regexp.QuoteMeta("UPDATE orders SET status = ? WHERE id = ?")).
		WithArgs(models.OrderApproved, "ord-1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	ok, err := repo.UpdateOrderStatus("ord-1", models.OrderApproved)
	if err != nil { t.Fatalf("UpdateOrderStatus err: %v", err) }
	if !ok { t.Fatalf("expected ok=true") }
	if err := mock.ExpectationsWereMet(); err != nil { t.Fatalf("sql expectations: %v", err) }
}
