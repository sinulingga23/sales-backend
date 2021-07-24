package model


import (
	"fmt"
	"errors"
	"crypto/subtle"
	"sales-backend/utility"

	"golang.org/x/crypto/bcrypt"
)

type Login struct {
}

func (l Login) EncryptPassword(password string) (string, error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%s", passwordBytes), nil
}

func (l Login) CheckPasswordMatch(currentPassword string, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(hashedPassword))
	if err != nil {
		return false, err
	}
	return true, nil
}


func (l Login) Login(email string, password string) (bool, error) {
	db, err := utility.ConnectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var expectedEmail 	string
	var expectedPassword 	string
	err = db.QueryRow("SELECT email, password FROM users WHERE email = ?", email).Scan(&expectedEmail, &expectedPassword)
	if err != nil {
		return false, err
	}

	isEmailMatch := (subtle.ConstantTimeCompare([]byte(email)[:], []byte(expectedEmail)[:]) == 1)
	err = bcrypt.CompareHashAndPassword([]byte(expectedPassword), []byte(password))
	if err != nil {
		return false, errors.New("Somethings wrong!")
	}

	if isEmailMatch && err == nil {
		return true, nil
	}
	return false, errors.New("Somethings wrong!")
}
