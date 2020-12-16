package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/rene5394/my-twitter/models"
)

// GenerateJWT generates the encryption with JWT
func GenerateJWT(u models.User) (string, error) {
	myKey := []byte("ReneTorres")
	payload := jwt.MapClaims{
		"email":     u.Email,
		"name":      u.Name,
		"lastname":  u.Lastname,
		"birthdate": u.Birthdate,
		"biography": u.Biography,
		"location":  u.Location,
		"website":   u.Website,
		"_id":       u.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
