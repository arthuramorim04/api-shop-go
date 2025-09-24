package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arthu/shop-api-go/tests/testutils"
	"github.com/arthu/shop-api-go/internal/config"
)

func TestHealth(t *testing.T) {
	cfg := &config.Config{}
	r := testutils.NewServer(cfg, nil)

	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}
