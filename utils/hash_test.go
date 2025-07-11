package utils

import "testing"

func TestHashPassword(t *testing.T) {
	password := "secure123"
	hashedPass, err := HashPassword(password)

	if err != nil {
		t.Errorf("Failed to hash password: %v", err)
	}

	if !CheckPasswordHash(password, hashedPass) {
		t.Errorf("Password and hash do not match")
	}
}
