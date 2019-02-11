package student

import (
	"context"

	"github.com/gofrs/uuid"
)

// Service ...
type Service interface {
	SaveStudent(ctx context.Context, student Student) (string, error)
	GetStudent(ctx context.Context, id string) (*Student, error)
	DeleteStudent(ctx context.Context, id string) error
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
