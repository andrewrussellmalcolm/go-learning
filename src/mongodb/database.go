package main

import (
	"time"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Message :
type Message struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Text      string        `json:"text,omitempty"`
	Payload   interface{}   `bson:"payload,omitempty"`
	Timestamp time.Time     `json:"timestamp,omitempty"`
}

var session *mgo.Session

/** */
func initDB() {

	var err error
	session, err = mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
}

/** */
func closeDB() {
	session.Close()
}

/** */
func insertMessageDB(message Message) bson.ObjectId {

	message.ID = bson.NewObjectId()
	message.Timestamp = time.Now()
	c := session.DB("test").C("messages")
	err := c.Insert(message)

	if err != nil {
		panic(err)
	}

	return message.ID
}

/** */
func updateMessageDB(message Message) bson.ObjectId {

	message.Timestamp = time.Now()
	c := session.DB("test").C("messages")
	err := c.UpdateId(message.ID, message)

	if err != nil {
		panic(err)
	}

	return message.ID
}

/** */
func deleteMessageDB(id bson.ObjectId) {

	c := session.DB("test").C("messages")
	err := c.RemoveId(id)

	if err != nil {
		panic(err)
	}
}

/** */
func queryAllMessagesDB() []Message {

	var messages []Message
	c := session.DB("test").C("messages")
	err := c.Find(nil).Sort("-timestamp").All(&messages)

	if err != nil {
		panic(err)
	}

	return messages
}

/** */
func deleteAllMessagesDB() {
	c := session.DB("test").C("messages")

	_, err := c.RemoveAll(bson.M{})

	if err != nil {
		panic(err)
	}
}
