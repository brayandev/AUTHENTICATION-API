package authentication

import "gopkg.in/mgo.v2/bson"

// User test usr...
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Login    string        `bson:"login" json:"login"`
	Password string        `bson:"password" json:"password"`
	Name     string        `bson:"name" json:"name"`
	Email    string        `bson:"email" json:"email"`
	TypeUser string        `bson:"typeUser" json:"typeUser"`
}

type typeUser string
