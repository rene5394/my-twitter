package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rene5394/my-twitter/middlewares"
	"github.com/rene5394/my-twitter/routers"
	"github.com/rs/cors"
)

// Handlers set the PORT and listen the Server
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.Profile))).Methods("GET")
	router.HandleFunc("/modifyProfile", middlewares.CheckDB(middlewares.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/readTweets", middlewares.CheckDB(middlewares.ValidateJWT(routers.ReadTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8085"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
