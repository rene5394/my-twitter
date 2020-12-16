package main

import (
	"log"

	"github.com/rene5394/my-twitter/db"
	"github.com/rene5394/my-twitter/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Without DB Connection")
		return
	}
	handlers.Handlers()
}
