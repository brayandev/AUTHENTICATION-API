package authentication

// Repository ...
type Repository interface {
	Authentication(Login, Password string) error
	Save(register Register) error
}

// RegisterDB ...
type RegisterDB struct {
	Server   string
	Database string
}
