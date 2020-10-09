package db

import (
	"github.com/rene5394/my-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

// TryLogin check the login information with the DB
func TryLogin(email string, password string) (models.User, bool) {
	user, found, _ := CheckUserExist(email)
	if found == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true
}
