package routers

import (
	"encoding/json"
	"net/http"

	"github.com/rene5394/my-twitter/db"
	"github.com/rene5394/my-twitter/models"
)

// ModifyProfile allows to change profile data
func ModifyProfile(w http.ResponseWriter, r *http.Request) {

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Data wrong "+err.Error(), 400)
		return
	}

	var status bool

	status, err = db.ModifyProfile(u, userID)
	if err != nil {
		http.Error(w, "Error trying to modify profile, please try again "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "It was not possible to modify the user profile "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
