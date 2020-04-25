package util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	pStr := string(hashedPassword)
	return pStr, nil
}
