package helpers

import "golang.org/x/crypto/bcrypt"

const salt = 8

func HashPassword(p string) (string, error) {
	password := []byte(p)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, salt)
	return string(hashedPassword), err
}

func ValidatePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
