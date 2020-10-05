package middlewares

import (
	"net/http"

	"github.com/rene5394/my-twitter/db"
)

// CheckDB is the middleware allows to know the database status
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Lost connection with the database", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
