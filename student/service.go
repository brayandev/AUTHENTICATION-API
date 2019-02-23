package student

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
)

// Service ...
type Service interface {
	SaveStudent(ctx context.Context, student Student) (string, error)
	GetStudent(ctx context.Context, id string) (*Student, error)
	DeleteStudent(ctx context.Context, id string) error
	UpdateStudent(ctx context.Context, id string, student *UpdateStudent) (*UpdateStudentResult, error)
}

// ServiceImpl ...
type ServiceImpl struct {
	repository Repository
}

// NewService ...
func NewService(repository Repository) *ServiceImpl {
	return &ServiceImpl{repository: repository}
}

// SaveStudent ...
func (s *ServiceImpl) SaveStudent(ctx context.Context, student Student) (string, error) {
	studentID, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	student.ID = studentID.String()
	student.Creation = time.Now().UTC()

	err = s.repository.save(ctx, student)
	if err != nil {
		return "", err
	}

	return student.ID, nil
}

// GetStudent ...
func (s *ServiceImpl) GetStudent(ctx context.Context, id string) (*Student, error) {
	return s.repository.get(ctx, id)
}

// DeleteStudent ...
func (s *ServiceImpl) DeleteStudent(ctx context.Context, id string) error {
	return s.repository.delete(ctx, id)
}

// UpdateStudent ...
func (s *ServiceImpl) UpdateStudent(ctx context.Context, id string, updateStudent *UpdateStudent) (*UpdateStudentResult, error) {
	currentStudent, err := s.repository.get(ctx, id)
	if err != nil {
		return nil, err
	}

	currentStudent.Name = updateStudent.Name
	currentStudent.Email = updateStudent.Email
	currentStudent.Login = updateStudent.Login
	currentStudent.Password = updateStudent.Password
	currentStudent.LastUpdate = time.Now().UTC()

	uErr := s.repository.update(ctx, currentStudent)
	if uErr != nil {
		return nil, uErr
	}

	return &UpdateStudentResult{
		ID: currentStudent.ID,
	}, nil
}
