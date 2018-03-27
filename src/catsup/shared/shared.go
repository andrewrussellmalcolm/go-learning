package shared

import (
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

// User :
type User struct {
	ID    bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string        `json:"name,omitempty"`
	Email string        `json:"email,omitempty"`
	Hash  string        `json:"hash,omitempty"`
}

func init() {

}
