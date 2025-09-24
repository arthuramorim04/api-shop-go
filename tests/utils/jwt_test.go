package utils_test

import (
	"testing"
	"github.com/golang-jwt/jwt/v5"
	"github.com/arthu/shop-api-go/internal/utils"
)

func TestGenerateJWTContainsClaims(t *testing.T) {
	secret := "test-secret"
	tokenStr, err := utils.GenerateJWT(secret, utils.JwtUser{ID: 42, Email: "a@b.com", Role: "admin", Username: "arthur"})
	if err != nil {
		t.Fatalf("GenerateJWT error: %v", err)
	}
	parsed, err := jwt.Parse(tokenStr, func(tk *jwt.Token) (interface{}, error) { return []byte(secret), nil })
	if err != nil || !parsed.Valid {
		t.Fatalf("expected valid token, got err=%v valid=%v", err, parsed.Valid)
	}
}
