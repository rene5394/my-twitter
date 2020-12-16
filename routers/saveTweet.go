package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rene5394/my-twitter/db"
	"github.com/rene5394/my-twitter/models"
)

// SaveTweet is a function
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	register := models.SaveTweet{
		UserID:  userID,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.InsertTweet(register)
	if err != nil {
		http.Error(w, "Error trying to insert the register, please try again"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Not was possible to insert tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
