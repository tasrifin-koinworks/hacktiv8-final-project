package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	SALT = 8
)

func HashPassword(pass string) string {
	password := []byte(pass)

	hash, _ := bcrypt.GenerateFromPassword(password, SALT)

	return string(hash)
}

func ComparePassword(hashed, password []byte) bool {
	h, p := []byte(hashed), []byte(password)

	err := bcrypt.CompareHashAndPassword(h, p)

	return err == nil
}
