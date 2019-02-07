package authentication

// Student is structure to persist a student...
type Student struct {
	ID       string `bson:"studentID" json:"studentID"`
	Login    string `bson:"login" json:"login"`
	Password string `bson:"password" json:"password"`
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
}

// Version ...
func (s *Student) Version() string {
	return "vnd.authentication-student.student.v1"
}
