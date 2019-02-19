package main

// Versionable ...
type Versionable interface {
	Version() string
}

type postStudentCreatedResponse struct {
	ID string `json:"studentId"`
}

func (pr *postStudentCreatedResponse) Version() string {
	return "vnd.post-student.v1"
}
