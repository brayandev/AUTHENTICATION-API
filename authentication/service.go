package authentication

import (
	"context"

	"gopkg.in/mgo.v2/bson"
)

const (
	student typeUser = "student"
	teacher typeUser = "teacher"
)

// Service ...
type Service interface {
	Save(ctx context.Context, user User) (*UserInserted, error)
}

// ServiceImpl ...
type ServiceImpl struct {
	repository Repository
}

// UserInserted ...
type UserInserted struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	TypeUser string `json:"typeUser"`
}

// Version ...
func (UserInserted) Version() string {
	return "" //TODO version
}

// NewService ...
func NewService(repository Repository) *ServiceImpl {
	return &ServiceImpl{repository: repository}
}

// Save ...
func (s *ServiceImpl) Save(ctx context.Context, user User) (*UserInserted, error) {
	user = User{
		Login:    user.Login,
		Password: user.Password,
		Name:     user.Name,
		Email:    user.Email,
		TypeUser: user.TypeUser,
	}
	user.ID = bson.NewObjectId()

	err := s.repository.save(user)
	if err != nil {
		return nil, err
	}

	userInserted := &UserInserted{
		Login:    user.Login,
		Email:    user.Email,
		Name:     user.Name,
		TypeUser: user.TypeUser,
	}

	return userInserted, nil
}
