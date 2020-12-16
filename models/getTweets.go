package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetTweets is the structure to return tweets
type GetTweets struct {
	ID      primitive.ObjectID `bson:"id" json:"_id,omitempty"`
	UserID  string             `bson:"userid" json:"userid,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}
