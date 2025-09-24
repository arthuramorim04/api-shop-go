package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtUser struct {
	ID       int64
	Email    string
	Role     string
	Username string
}

func GenerateJWT(secret string, u JwtUser) (string, error) {
	claims := jwt.MapClaims{
		"id":       u.ID,
		"email":    u.Email,
		"role":     u.Role,
		"username": u.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
