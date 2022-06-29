package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) string {
	salt := 10
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)
	return string(hash)
}

func ComparePassword(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}