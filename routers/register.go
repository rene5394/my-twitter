package routers

import (
	"encoding/json"
	"net/http"

	"github.com/rene5394/my-twitter/db"
	"github.com/rene5394/my-twitter/models"
)

// Register is the function to create the user in the database
func Register(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Error in the data "+err.Error(), 400)
		return
	}

	if len(u.Email) == 0 {
		http.Error(w, "Email user is requerid", 400)
		return
	}
	if len(u.Password) < 6 {
		http.Error(w, "Password should have at least 6 characters", 400)
		return
	}

	_, found, _ := db.CheckUserExist(u.Email)
	if found == true {
		http.Error(w, "There is already a user with that email account", 400)
		return
	}

	_, status, err := db.InsertRegister(u)
	if err != nil {
		http.Error(w, "There is an error tryiong to register the user "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Failed to insert user record, please try later", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
