package database

import (
	"github.com/globalsign/mgo"
)

var session *mgo.Session

// Init :
func Init(database string) error {

	s, err := mgo.Dial(database)

	if err != nil {
		return err
	}

	session = s
	return nil
}

// Close :
func Close() {
	session.Close()
}
