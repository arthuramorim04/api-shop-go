package utils_test

import (
	"testing"

	"github.com/arthu/shop-api-go/internal/utils"
)

func TestHashAndCheckPassword(t *testing.T) {
	pw := "secret123"
	h, err := utils.HashPassword(pw)
	if err != nil {
		t.Fatalf("unexpected error hashing password: %v", err)
	}
	if h == "" || h == pw {
		t.Fatalf("hash should not be empty or equal to the original password")
	}
	if ok := utils.CheckPassword(h, pw); !ok {
		t.Fatalf("expected CheckPassword to return true with correct password")
	}
	if ok := utils.CheckPassword(h, "wrong"); ok {
		t.Fatalf("expected CheckPassword to return false with wrong password")
	}
}
