package db

import (
	"context"
	"time"

	"github.com/rene5394/my-twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertRegister is a function to add the user to the database
func InsertRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := Connection.Database("twitter")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
