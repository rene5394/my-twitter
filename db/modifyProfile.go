package db

import (
	"context"
	"time"

	"github.com/rene5394/my-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ModifyProfile allows to change profile data
func ModifyProfile(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := Connection.Database("twitter")
	col := db.Collection("users")

	profile := make(map[string]interface{})
	if len(u.Name) > 0 {
		profile["name"] = u.Name
	}
	if len(u.Lastname) > 0 {
		profile["lastname"] = u.Lastname
	}
	profile["birthdate"] = u.Birthdate
	if len(u.Avatar) > 0 {
		profile["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		profile["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		profile["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		profile["location"] = u.Location
	}
	if len(u.Website) > 0 {
		profile["website"] = u.Website
	}
	if len(u.Website) > 0 {
		profile["website"] = u.Website
	}

	upadateString := bson.M{
		"$set": profile,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, upadateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
