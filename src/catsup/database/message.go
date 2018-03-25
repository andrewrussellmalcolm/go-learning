package database

import (
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
)

// Status :
type Status int

const (
	// WAITING :
	WAITING = 0
	// SENT :
	SENT = 1
	//RECEIVED :
	RECEIVED = 2
	// READ :
	READ = 3
)

// Message :
type Message struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Text      string
	Timestamp time.Time
	To        bson.ObjectId
	From      bson.ObjectId
	Status    Status
}

// InsertMessage :
func InsertMessage(message Message, to, from bson.ObjectId) (bson.ObjectId, error) {

	message.ID = bson.NewObjectId()
	message.Timestamp = time.Now()
	message.To = to
	message.From = from
	message.Status = WAITING
	c := session.DB("test").C("messages")

	err := c.Insert(message)

	if err != nil {
		log.Println(err)

		return "", err
	}

	return message.ID, nil
}

/** */
func UpdateMessage(message Message) (bson.ObjectId, error) {

	message.Timestamp = time.Now()
	c := session.DB("test").C("messages")
	err := c.UpdateId(message.ID, message)

	if err != nil {
		return "", err
	}

	return message.ID, nil
}

/** */
func DeleteMessage(id bson.ObjectId) error {

	c := session.DB("test").C("messages")
	err := c.RemoveId(id)

	return err
}

/** */
func QueryMessage(id bson.ObjectId) (Message, error) {

	return Message{}, nil
}

// GetMessageList :
func GetMessageList(toID, fromID bson.ObjectId) []Message {

	var messages []Message
	c := session.DB("test").C("messages")
	err := c.Find(bson.M{"to": toID, "from": fromID}).Sort("-timestamp").All(&messages)

	if err != nil {
		return nil
	}

	return messages
}

/** */
func QueryAllUnreadMessages() ([]Message, error) {

	var messages []Message
	c := session.DB("test").C("messages")
	err := c.Find(nil).Sort("-timestamp").All(&messages)

	if err != nil {
		return nil, err
	}

	return messages, nil
}

/** */
func DeleteAllMessages() error {
	c := session.DB("test").C("messages")

	_, err := c.RemoveAll(bson.M{})

	return err
}
