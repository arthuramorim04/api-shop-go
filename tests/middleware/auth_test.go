package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/arthu/shop-api-go/api/middleware"
	"github.com/arthu/shop-api-go/internal/config"
	"github.com/arthu/shop-api-go/internal/utils"
)

func TestAuthenticateToken_ValidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	cfg := &config.Config{JWTSecret: "secret"}
	r.GET("/protected", middleware.AuthenticateToken(cfg), func(c *gin.Context) { c.Status(http.StatusOK) })

	token, err := utils.GenerateJWT(cfg.JWTSecret, utils.JwtUser{ID: 1, Email: "a@b.com", Role: "admin", Username: "john"})
	if err != nil { t.Fatalf("token err: %v", err) }

	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestAuthenticateToken_MissingToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	cfg := &config.Config{JWTSecret: "secret"}
	r.GET("/protected", middleware.AuthenticateToken(cfg), func(c *gin.Context) { c.Status(http.StatusOK) })

	req := httptest.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", w.Code)
	}
}

func TestAuthorizeRole_AllowsMatchingRole(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	cfg := &config.Config{JWTSecret: "secret"}
	r.GET("/admin", middleware.AuthenticateToken(cfg), middleware.AuthorizeRole("admin"), func(c *gin.Context) { c.Status(http.StatusOK) })

	token, _ := utils.GenerateJWT(cfg.JWTSecret, utils.JwtUser{ID: 1, Email: "a@b.com", Role: "admin", Username: "john"})
	req := httptest.NewRequest("GET", "/admin", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK { t.Fatalf("expected 200, got %d", w.Code) }
}

func TestAuthorizeRole_RejectsDifferentRole(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	cfg := &config.Config{JWTSecret: "secret"}
	r.GET("/admin", middleware.AuthenticateToken(cfg), middleware.AuthorizeRole("admin"), func(c *gin.Context) { c.Status(http.StatusOK) })

	token, _ := utils.GenerateJWT(cfg.JWTSecret, utils.JwtUser{ID: 1, Email: "a@b.com", Role: "client", Username: "john"})
	req := httptest.NewRequest("GET", "/admin", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden { t.Fatalf("expected 403, got %d", w.Code) }
}
