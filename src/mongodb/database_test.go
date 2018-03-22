package main

import (
	"fmt"
	"testing"
)

// TestMain :
func TestMain(m *testing.M) {
	initDB()
	m.Run()
	closeDB()
}

// TestInsert :
func TestInsert(t *testing.T) {

	fmt.Println("TestInsert")

	deleteAllMessagesDB()

	messages := queryAllMessagesDB()

	if len(messages) != 0 {
		t.Errorf("delete all failed")
	}

	insertMessageDB(Message{Text: "Hello"})
	messages = queryAllMessagesDB()

	if len(messages) != 1 {
		t.Errorf("insert message failed")
	}

	fmt.Println("done")
}

func TestDelete(t *testing.T) {

	fmt.Println("TestDelete")

	deleteAllMessagesDB()

	messages := queryAllMessagesDB()

	if len(messages) != 0 {
		t.Errorf("delete all failed")
	}

	insertMessageDB(Message{Text: "Hello"})
	messages = queryAllMessagesDB()

	if len(messages) != 1 {
		t.Errorf("insert message failed")
	}

	deleteMessageDB(messages[0].ID)
	messages = queryAllMessagesDB()

	if len(messages) != 0 {
		t.Errorf("delete message failed")
	}

	fmt.Println("done")
}
