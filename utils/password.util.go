package utils

import "golang.org/x/crypto/bcrypt"

const (
	hashCost = 15
)

func HashPassword(password string) (string, error) {
	if hash, err := bcrypt.GenerateFromPassword([]byte(password), hashCost); err != nil {
		return "", err
	} else {
		return string(hash), nil
	}
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

