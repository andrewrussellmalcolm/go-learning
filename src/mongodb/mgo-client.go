package main

import (
	"fmt"
	"time"
)

func main() {

	initDB()
	defer closeDB()

	deleteAllMessagesDB()

	insertMessageDB(Message{Text: "Hello"})
	messages := queryAllMessagesDB()
	fmt.Println("== should be one message")
	fmt.Println(messages)

	deleteAllMessagesDB()
	insertMessageDB(Message{Text: "Hello"})
	insertMessageDB(Message{Text: "World"})
	messages = queryAllMessagesDB()
	fmt.Println("== should be two messages")
	fmt.Println(messages)

	deleteAllMessagesDB()
	messages = queryAllMessagesDB()
	fmt.Println("== should be no messages")
	fmt.Println(messages)

	insertMessageDB(Message{Text: "Hello"})
	messages = queryAllMessagesDB()
	fmt.Println("== should be 1 message 'hello'")
	fmt.Println(messages)
	time.Sleep(2 * time.Second)
	messages[0].Text = "World"
	updateMessageDB(messages[0])
	messages = queryAllMessagesDB()
	fmt.Println("== should be 1 message 'world'")
	fmt.Println(messages)
}
