package authentication

import (
	"os"

	mgo "gopkg.in/mgo.v2"
)

const (
	collection = "User"
	database   = "CourseDB"
)

// Repository ...
type Repository interface {
	authentication(Login, Password string) error
	save(user User) error
}

// RepositoryImpl ...
type RepositoryImpl struct {
	collectionName string
	session        *mgo.Session
}

// Session ...
type Session struct {
	session  *mgo.Session
	Database *mgo.Database
}

// NewRepository ...
func NewRepository(collectionName string, session *mgo.Session) *RepositoryImpl {
	return &RepositoryImpl{
		collectionName: collectionName,
		session:        session,
	}
}

// NewMongoDB ...
func NewMongoDB(endpoint string) (*mgo.Session, error) {
	session, err := mgo.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	return session, err
}

// Save ...
func (r *RepositoryImpl) save(user User) error {
	session, err := NewMongoDB(os.Getenv("HOST_MONGODB"))
	if err != nil {
		return err
	}
	c := session.DB(database).C(collection)
	return c.Insert(user)
}

// Authentication ...
func (r *RepositoryImpl) authentication(Login, Password string) error {
	return nil
}
