package db

import "golang.org/x/crypto/bcrypt"

// EncryptPassword is the function encrypts the password
func EncryptPassword(pass string) (string, error) {
	lenght := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), lenght)
	return string(bytes), err
}
