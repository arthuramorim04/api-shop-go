package handlers_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/DATA-DOG/go-sqlmock"
    "github.com/arthu/shop-api-go/internal/config"
    "github.com/arthu/shop-api-go/internal/mp"
    "github.com/arthu/shop-api-go/tests/mocks"
    "github.com/arthu/shop-api-go/tests/testutils"
)

func TestPaymentNotification_Approved(t *testing.T) {
    cfg := &config.Config{MPAccessToken: "token"}
    mpMock := &mocks.MockMPClient{
        GetPaymentFunc: func(accessToken, paymentID string) (string, error) {
            if accessToken != cfg.MPAccessToken || paymentID != "pay-1" {
                t.Fatalf("unexpected args: %s %s", accessToken, paymentID)
            }
            return "approved", nil
        },
    }
    old := mp.ClientImpl
    mp.ClientImpl = mpMock
    t.Cleanup(func() { mp.ClientImpl = old })

    db, sqlMock, err := testutils.NewSQLMock()
    if err != nil { t.Fatalf("sqlmock: %v", err) }
    r := testutils.NewServer(cfg, db)

    // Expect update order status
    sqlMock.ExpectExec("UPDATE orders SET status").WithArgs("approved", "order-123").WillReturnResult(sqlmock.NewResult(0, 1))

    body := map[string]interface{}{
        "data": map[string]string{"id": "pay-1"},
    }
    b, _ := json.Marshal(body)
    req := httptest.NewRequest("POST", "/payment/notification?orderID=order-123", bytes.NewReader(b))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d, body=%s", w.Code, w.Body.String())
    }
    if err := sqlMock.ExpectationsWereMet(); err != nil {
        t.Fatalf("sql expectations: %v", err)
    }
}
