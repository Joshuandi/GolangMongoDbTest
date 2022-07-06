package util

import (
	"errors"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

//
func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidateEmail(email string) (string, bool) {
	em, err := mail.ParseAddress(email)
	if err != nil {
		return "", false
	}
	return em.Address, true
}

func EmptyValidation(input string) (string error) {
	if input == "" {
		return errors.New("Please input your empty data")
	}
	return nil
}
