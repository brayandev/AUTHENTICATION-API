package authentication

import (
	"fmt"

	"github.com/API-AUTENTICATION/config"
	mgo "gopkg.in/mgo.v2"
)

// Repository ...
type Repository interface {
	Authentication(Login, Password string) error
	Save(register *Register) error
}

// Session ...
type Session struct {
	session *mgo.Session
}

// NewSession ...
func NewSession(url string) (*Session, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	return &Session{session}, err
}

// Save ...
func Save(register *Register) error {
	config := config.NewConfig()
	session, err := NewSession(config.MongoDBEndpoint)
	if err != nil {
		return err
	}
	fmt.Println(session)
	return nil
}
