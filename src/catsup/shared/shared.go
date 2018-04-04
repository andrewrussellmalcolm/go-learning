package shared

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// MessageStatus :
type MessageStatus int

// UserStatus :
type UserStatus int

const (
	// WAITING :
	WAITING = 0
	// SENT :
	SENT = 1
	//RECEIVED :
	RECEIVED = 2
	// READ :
	READ = 3
	// ONLINE :
	ONLINE = 4
	// OFFLINE :
	OFFLINE = 5
)

// Message :
type Message struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	To        bson.ObjectId `json:"to,omitempty"`
	From      bson.ObjectId `json:"from,omitempty"`
	Timestamp time.Time     `json:"timestamp,omitempty"`
	Status    MessageStatus `json:"message_status,omitempty"`
	Text      string        `json:"text,omitempty"`
}

// User :
type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Timestamp time.Time     `json:"timestamp,omitempty"`
	Name      string        `json:"name,omitempty"`
	Email     string        `json:"email,omitempty"`
	Hash      string        `json:"hash,omitempty"`
}

func init() {

}
