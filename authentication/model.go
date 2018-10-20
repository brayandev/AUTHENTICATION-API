package authentication

import (
	"gopkg.in/mgo.v2/bson"
)

// Register ...
type Register struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	User     string        `bson:"user" json:"user"`
	Password string        `bson:"password" json:"password"`
	Name     string        `bson:"name" json:"name"`
	Email    string        `bson:"email" json:"email"`
}

// Login ...
type Login struct {
	User     string `bson:"user" json:"user"`
	Password string `bson:"password" json:"password"`
}
