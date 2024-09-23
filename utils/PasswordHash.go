package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordWithHash(storedPasswwordHash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(storedPasswwordHash), []byte(password)) == nil
}
