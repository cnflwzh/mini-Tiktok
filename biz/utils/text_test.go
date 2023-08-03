package utils

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "testPassword"
	hash, err := HashPassword(password)
	hlog.Info("hash", hash)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
		return
	}
	if len(hash) == 0 {
		t.Errorf("Expected hashed password, but got empty string")
		return
	}
}
func TestCheckPasswordHash(t *testing.T) {
	password := "testPassword"
	if !CheckPasswordHash(password, "$2a$14$Z.emErTSQ9fN8PnyaxCR8OJZnDi4Iwlw3Ji6Cp0mtC1mFP7fpuONS") {
		t.Errorf("Expected password to match hash, but got false")
		return
	}
}
