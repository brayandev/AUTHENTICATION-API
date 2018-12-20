package authentication

import (
	"errors"
	"os"

	mgo "gopkg.in/mgo.v2"
)

// Repository is repo ...
type Repository interface {
	authentication(Login, Password string) error
	save(user User) error
}

// RepositoryImpl implements repository...
type RepositoryImpl struct {
	session *mgo.Session
}

// Session ...
type Session struct {
	session  *mgo.Session
	Database *mgo.Database
}

// NewRepository ...
func NewRepository(session *mgo.Session) *RepositoryImpl {
	return &RepositoryImpl{
		session: session,
	}
}

// NewMongoDB ...
func NewMongoDB(endpoint string) (*mgo.Session, error) {
	var mgoSession *mgo.Session
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(endpoint)
		if err != nil {
			return nil, errors.New("Failed to start the mongo session")
		}
	}

	return mgoSession.Clone(), nil
}

// Save ...
func (r *RepositoryImpl) save(user User) error {
	session, err := NewMongoDB(os.Getenv("HOST_MONGODB"))
	if err != nil {
		return err
	}
	c := session.DB(os.Getenv("MONGO_DB_NAME")).C(os.Getenv("MOND_DB_COLLECTION"))
	return c.Insert(user)
}

// Authentication ...
func (r *RepositoryImpl) authentication(Login, Password string) error {
	return nil
}
