package authentication

// Service ...
type Service interface {
	Save(user, password, name, email string) error
}

// ServiceImpl ...
type ServiceImpl struct {
	repository Repository
}

// NewService ...
func NewService(rep Repository) *ServiceImpl {
	return &ServiceImpl{repository: rep}
}

// Save ...
func (s *ServiceImpl) Save(user, password, name, email string) error {
	register := Register{
		User:     user,
		Password: password,
		Name:     name,
		Email:    email,
	}
	return s.repository.Save(register)
}
