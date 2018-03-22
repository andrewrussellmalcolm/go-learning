package main

import (
	"fmt"
	"testing"
)

// TestMain :
func TestMain(m *testing.M) {

}

// TestInsert :
func TestInsert(t *testing.T) {
	initDB()
	defer closeDB()
	deleteAllMessagesDB()

	messages := queryAllMessagesDB()

	if len(messages) != 1 {
		t.Errorf("delete all failed")
	}

	insertMessageDB(Message{Text: "Hello"})
	messages = queryAllMessagesDB()

	if len(messages) != 1 {
		t.Errorf("insert message failed")
	}

	fmt.Println("done")
}
