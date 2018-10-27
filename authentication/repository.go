package authentication

import (
	mgo "gopkg.in/mgo.v2"
)

const (
	// Collection ...
	Collection = "users"
)

// Repository ...
type Repository interface {
	Authentication(Login, Password string) error
	Save(register Register) error
}

// RepositoryImpl ...
type RepositoryImpl struct {
	collectionName string
	session        *mgo.Session
}

// Session ...
type Session struct {
	session *mgo.Session
}

// NewRepository ...
func NewRepository(collectionName string, session *mgo.Session) *RepositoryImpl {
	return &RepositoryImpl{
		collectionName: collectionName,
		session:        session,
	}
}

// NewMongoDB ...
func NewMongoDB(url string) (*Session, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	return &Session{session}, err
}

// Save ...
func Save(register Register) error {
	return nil
}
