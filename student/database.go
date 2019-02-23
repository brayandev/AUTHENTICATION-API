package student

import mgo "gopkg.in/mgo.v2"

// NewMongoDB ...
func NewMongoDB(endpoint string) (*mgo.Session, error) {
	var mgoSession *mgo.Session
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(endpoint)
		if err != nil {
			return nil, err
		}
	}

	return mgoSession.Clone(), nil
}
