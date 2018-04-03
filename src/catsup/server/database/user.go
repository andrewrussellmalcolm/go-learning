package database

import (
	"catsup/shared"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
)

// CheckAuth :
func CheckAuth(name, pass string) *shared.User {
	c := session.DB("test").C("users")

	user := shared.User{}
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
func GetUserByID(id string) *shared.User {
	c := session.DB("test").C("users")

	user := shared.User{}
	err := c.Find(bson.M{"_id": id}).One(&user)

	if err != nil {

		return nil
	}

	return &user
}

// GetUserList :
func GetUserList() []shared.User {
	c := session.DB("test").C("users")

	var users []shared.User
	err := c.Find(nil).All(&users)

	if err != nil {

		return nil
	}

	for _, user := range users {
		fmt.Printf("=====%d\n", user.LastAccess.Minute())
	}
	return users
}

// CreateUser :
func CreateUser(name, email, pass string) bool {

	user := shared.User{Name: name, Email: email, Hash: hashPassword(pass)}

	c := session.DB("test").C("users")

	err := c.Insert(user)

	return err == nil
}

// UpdateUserAccessTime :
func UpdateUserAccessTime(id bson.ObjectId, time time.Time) bool {

	c := session.DB("test").C("users")

	err := c.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"last_access": time}})

	if err != nil {

		panic(err)
	}
	fmt.Printf("%v %v\n", id, time)
	return err == nil
}
