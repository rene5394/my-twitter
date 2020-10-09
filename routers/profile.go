package routers

import (
	"encoding/json"
	"net/http"

	"github.com/rene5394/my-twitter/db"
)

// Profile allows to get profile's values
func Profile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Should send the ID", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Error finding profile"+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
