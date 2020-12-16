package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rene5394/my-twitter/db"
	"github.com/rene5394/my-twitter/jwt"
	"github.com/rene5394/my-twitter/models"
)

// Login is a function
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Wrong user or password "+err.Error(), 400)
		return
	}
	if len(u.Email) == 0 {
		http.Error(w, "User email is required", 400)
		return
	}
	document, exist := db.TryLogin(u.Email, u.Password)
	if exist == false {
		http.Error(w, "Wrong user or password", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Error trying to generate token "+err.Error(), 400)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
