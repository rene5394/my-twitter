package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection us the object of connection to the DB
var Connection = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://dbadmin:Sapojin4@mytwitter.h1gia.mongodb.net/twitter?retryWrites=true&w=majority")

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
