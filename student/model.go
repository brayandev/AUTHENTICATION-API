package student

import "time"

// Student is structure to persist a student.
type Student struct {
	ID         string    `bson:"studentId" json:"studentId"`
	Login      string    `bson:"login" json:"login"`
	Password   string    `bson:"password" json:"password"`
	Name       string    `bson:"name" json:"name"`
	Email      string    `bson:"email" json:"email"`
	Creation   time.Time `bson:"creation" json:"creation"`
	LastUpdate time.Time `bson:"lastUpdate" json:"lastUpdate"`
}

// Version is student version.
func (s *Student) Version() string {
	return "vnd.student.v1"
}

// UpdateStudent is structure to persist a student.
type UpdateStudent struct {
	Login    string `bson:"login" json:"login"`
	Password string `bson:"password" json:"password"`
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
}

// Version is student version.
func (UpdateStudent) Version() string {
	return ""
}

// UpdateStudentResult ...
type UpdateStudentResult struct {
	ID string `json:"id"`
}

// Version is student version.
func (UpdateStudentResult) Version() string {
	return ""
}
