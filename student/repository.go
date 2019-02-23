package student

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"

	mgo "gopkg.in/mgo.v2"
)

const (
	studentID = "studentId"
)

// Repository is repository ...
type Repository interface {
	authentication(Login, Password string) error
	save(ctx context.Context, student Student) error
	get(ctx context.Context, id string) (*Student, error)
	delete(ctx context.Context, id string) error
	update(ctx context.Context, student *Student) error
}

// RepositoryImpl implements repository...
type RepositoryImpl struct {
	session         *mgo.Session
	dbName          string
	mongoCollection string
}

// NewRepository ...
func NewRepository(session *mgo.Session, dbName, mongoCollection string) *RepositoryImpl {
	return &RepositoryImpl{
		session:         session,
		dbName:          dbName,
		mongoCollection: mongoCollection,
	}
}

func (r *RepositoryImpl) save(ctx context.Context, student Student) error {
	c := r.session.DB(r.dbName).C(r.mongoCollection)
	return c.Insert(student)
}

func (r *RepositoryImpl) get(ctx context.Context, id string) (*Student, error) {
	c := r.session.DB(r.dbName).C(r.mongoCollection)
	var student Student
	err := c.Find(bson.M{studentID: id}).One(&student)
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *RepositoryImpl) delete(ctx context.Context, id string) error {
	c := r.session.DB(r.dbName).C(r.mongoCollection)
	err := c.Remove(bson.M{studentID: id})
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryImpl) update(ctx context.Context, student *Student) error {
	c := r.session.DB(r.dbName).C(r.mongoCollection)
	return c.Update(bson.M{studentID: student.ID}, bson.M{"$set": student})
}

// Authentication ...
func (r *RepositoryImpl) authentication(Login, Password string) error {
	return nil
}
