package database

import (
	"catsup/shared"
	"log"
	"time"

	"github.com/globalsign/mgo/bson"
)

// CreateMessage :
func CreateMessage(message shared.Message, to, from bson.ObjectId) (bson.ObjectId, error) {

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

// QueryMessageList :
func QueryMessageList(toID, fromID bson.ObjectId) []shared.Message {

	messages := []shared.Message{}
	c := session.DB("test").C("messages")

	toAndFrom := bson.M{"$or": []bson.M{bson.M{"to": toID, "from": fromID}, bson.M{"to": fromID, "from": toID}}}

	err := c.Find(toAndFrom).Sort("-timestamp").All(&messages)

	if err != nil {
		return nil
	}

	_, err = c.UpdateAll(bson.M{}, bson.M{"$set": bson.M{"status": shared.READ}})

	if err != nil {
		panic(err)
	}

	return messages
}

// QueryNewMessageList :
func QueryNewMessageList(toID, fromID bson.ObjectId) []shared.Message {

	messages := []shared.Message{}
	c := session.DB("test").C("messages")

	query := bson.M{"$and": []bson.M{bson.M{"status": bson.M{"$ne": shared.READ}},
		bson.M{"$or": []bson.M{bson.M{"to": toID, "from": fromID}, bson.M{"to": fromID, "from": toID}}}}}

	err := c.Find(query).Sort("-timestamp").All(&messages)

	if err != nil {
		return nil
	}

	return messages
}

// DeleteMessage :
func DeleteMessage(id bson.ObjectId) error {

	c := session.DB("test").C("messages")
	err := c.RemoveId(id)

	return err
}

// DeleteAllMessages :
func DeleteAllMessages() error {
	c := session.DB("test").C("messages")

	_, err := c.RemoveAll(bson.M{})

	return err
}
