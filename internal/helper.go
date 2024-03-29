package internal

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Generate a salt with a cost factor of 10
	// Cost factor determines the computational cost of the hashing; higher values are slower but more secure
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
