package database

import (
	"crypto/sha256"
	"encoding/base64"

	"github.com/globalsign/mgo/bson"
)

// User :
type User struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string
	Email string
	Hash  string
}

// CheckAuth :
func CheckAuth(name, pass string) *User {
	c := session.DB("test").C("users")

	user := User{}
	err := c.Find(bson.M{"name": name}).One(&user)

	if err != nil || user.Hash != hashPassword(pass) {

		return nil
	}

	return &user
}

//
func hashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

// GetUserByID :
func GetUserByID(id string) *User {
	c := session.DB("test").C("users")

	user := User{}
	err := c.Find(bson.M{"_id": id}).One(&user)

	if err != nil {

		return nil
	}

	return &user
}

// GetUserList :
func GetUserList() []User {
	c := session.DB("test").C("users")

	var users []User
	err := c.Find(nil).All(&users)

	if err != nil {

		return nil
	}

	return users
}
