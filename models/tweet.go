package models

// Tweet captures the body, the message we recieved
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
