package main

import (
	"fmt"
)

type Payload1 struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
	Z float64 `json:"z,omitempty"`
}

type Payload2 struct {
	X     float64 `json:"x,omitempty"`
	Y     float64 `json:"y,omitempty"`
	Theta float64 `json:"theta,omitempty"`
}

type Payload3 struct {
	Data string `json:"data,omitempty"`
}

func main() {

	initDB()
	defer closeDB()

	deleteAllMessagesDB()

	payload1 := Payload1{1, 2, 3}
	payload2 := Payload2{4, 5, 6}
	payload3 := Payload3{"hippity hoppity hip"}

	insertMessageDB(Message{Text: "Payload 1", Payload: payload1})
	insertMessageDB(Message{Text: "Payload 2", Payload: payload2})
	insertMessageDB(Message{Text: "Payload 3", Payload: payload3})
	messages := queryAllMessagesDB()

	for _, message := range messages {
		fmt.Println(message)

		switch t := message.Payload.(type) {

		case Payload1:
			fmt.Println("P1")
		case Payload2:
			fmt.Println("P2")
		case Payload3:
			fmt.Println("P3")
		default:
			fmt.Println(t)
		}

	}

}
