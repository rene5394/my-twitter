package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/rene5394/my-twitter/db"
	"github.com/rene5394/my-twitter/models"
)

// Email value of Email used in all the endpoints
var Email string

// IDUsuario value of ID return by the model, that it uses in the endpoints
var IDUsuario string

// ProcessToken process the token to extact its values
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("ReneTorres")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Token format invalid")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(tk *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, found, _ := db.CheckUserExist(claims.Email)
		if found == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, found, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Invalid Token")
	}

	return claims, false, string(""), err
}
