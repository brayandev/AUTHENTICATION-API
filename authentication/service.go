package authentication

import "context"

// Service ...
type Service interface {
	Save(ctx context.Context, user *User) (*UserInserted, error)
}

// ServiceImpl ...
type ServiceImpl struct {
	repository Repository
}

// ParamUser ...
type ParamUser struct {
	ID       string `bson:"_id" json:"id"`
	Login    string `bson:"user" json:"user"`
	Password string `bson:"password" json:"password"`
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
	TypeUser string `bson:"typeUser" json:"typeUser"`
}

// UserInserted ...
type UserInserted struct {
	ID       string `json:"id"`
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
func (s *ServiceImpl) Save(ctx context.Context, paramUser *ParamUser) (*UserInserted, error) {
	user := User{
		ID:       paramUser.ID,
		Login:    paramUser.Login,
		Password: paramUser.Password,
		Name:     paramUser.Name,
		Email:    paramUser.Email,
		TypeUser: paramUser.TypeUser,
	}
	err := s.repository.save(user)
	if err != nil {
		return nil, err
	}

	userInserted := &UserInserted{
		ID:       user.ID,
		Name:     user.Name,
		TypeUser: user.TypeUser,
	}
	return userInserted, nil
}
