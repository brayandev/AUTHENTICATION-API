package authentication

// User ...
type User struct {
	ID       string `bson:"_id" json:"id"`
	Login    string `bson:"user" json:"user"`
	Password string `bson:"password" json:"password"`
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
	TypeUser string `bson:"typeUser" json:"typeUser"`
}

type typeUser string
