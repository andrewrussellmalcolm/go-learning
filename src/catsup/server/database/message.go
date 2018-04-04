package database

import (
	"catsup/shared"
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
)

// InsertMessage :
func InsertMessage(message shared.Message, to, from bson.ObjectId) (bson.ObjectId, error) {

	message.ID = bson.NewObjectId()
	message.Timestamp = time.Now()
	message.To = to
	message.From = from
	message.Status = shared.WAITING
	c := session.DB("test").C("messages")

	err := c.Insert(message)

	if err != nil {
		log.Println(err)

		return "", err
	}

	return message.ID, nil
}

/** */
func UpdateMessage(message shared.Message) (bson.ObjectId, error) {

	message.Timestamp = time.Now()
	c := session.DB("test").C("messages")
	err := c.UpdateId(message.ID, message)

	if err != nil {
		return "", err
	}

	return message.ID, nil
}

// DeleteMessage :
func DeleteMessage(id bson.ObjectId) error {

	c := session.DB("test").C("messages")
	err := c.RemoveId(id)

	return err
}

// QueryMessage :
func QueryMessage(id bson.ObjectId) (shared.Message, error) {

	return shared.Message{}, nil
}

// GetMessageList :
func GetMessageList(toID, fromID bson.ObjectId) []shared.Message {

	messages := []shared.Message{}
	c := session.DB("test").C("messages")

	toAndFrom := bson.M{"$or": []bson.M{bson.M{"to": toID, "from": fromID}, bson.M{"to": fromID, "from": toID}}}

	err := c.Find(toAndFrom).Sort("-timestamp").All(&messages)

	if err != nil {
		return nil
	}

	return messages
}

/** */
func QueryAllUnreadMessages() ([]shared.Message, error) {

	var messages []shared.Message
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
