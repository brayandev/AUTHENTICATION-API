package authentication

import (
	"context"

	"github.com/gofrs/uuid"
	mgo "gopkg.in/mgo.v2"
)

// Service ...
type Service interface {
	SaveStudent(ctx context.Context, student Student, session *mgo.Session) (string, error)
	GetStudent(ctx context.Context, id string, session *mgo.Session) (*Student, error)
	DeleteStudent(ctx context.Context, id string, session *mgo.Session) error
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
func (s *ServiceImpl) SaveStudent(ctx context.Context, student Student, session *mgo.Session) (string, error) {
	studentID, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	student.ID = studentID.String()
	err = s.repository.save(ctx, student, session)
	if err != nil {
		return "", err
	}

	return student.ID, nil
}

// GetStudent ...
func (s *ServiceImpl) GetStudent(ctx context.Context, id string, session *mgo.Session) (*Student, error) {
	return s.repository.get(ctx, id, session)
}

// DeleteStudent ...
func (s *ServiceImpl) DeleteStudent(ctx context.Context, id string, session *mgo.Session) error {
	return s.repository.delete(ctx, id, session)
}
