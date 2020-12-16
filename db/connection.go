package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection us the object of connection to the DB
var Connection = ConnectDB()
var uri = LoadDBVars()
var clientOptions = options.Client().ApplyURI(uri)

// LoadDBVars is a function that loads .env global variables
func LoadDBVars() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return ""
	}
	var dbURL = os.Getenv("DB_URL")
	var dbUser = os.Getenv("DB_USER")
	var dbPass = os.Getenv("DB_PASS")
	return "mongodb://" + dbUser + ":" + dbPass + "@" + dbURL
}

// ConnectDB is the functiont to connect to the DB
func ConnectDB() *mongo.Client {
	client, error := mongo.Connect(context.TODO(), clientOptions)
	if error != nil {
		log.Fatal(error.Error())
		return client
	}
	error = client.Ping(context.TODO(), nil)
	if error != nil {
		log.Fatal(error.Error())
		return client
	}
	log.Println("Sucessfull connection to the database")
	return client
}

// CheckConnection is the Ping to the DB
func CheckConnection() int {
	error := Connection.Ping(context.TODO(), nil)
	if error != nil {
		return 0
	}
	return 1
}
