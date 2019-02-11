package student

// Student is structure to persist a student.
type Student struct {
	ID       string `bson:"studentId" json:"studentId"`
	Login    string `bson:"login" json:"login"`
	Password string `bson:"password" json:"password"`
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
}

// Version is student version.
func (s *Student) Version() string {
	return "vnd.student.student.v1"
}
