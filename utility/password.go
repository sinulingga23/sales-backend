package utility

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%s", passwordBytes), nil
}

func CheckPasswordMatch(hashedPassword string, currentPassword string) (bool, error) {
	err := bcrypt. CompareHashAndPassword([]byte(hashedPassword), []byte(currentPassword))
	if err != nil {
		return false, err
	}
	return true, nil
}
