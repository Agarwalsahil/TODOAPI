package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(passwrod string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(passwrod), bcrypt.DefaultCost)

	return string(hashedPass), err
}

func CheckPasswordHash(password, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	return err == nil
}