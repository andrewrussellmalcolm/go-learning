package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

// Client :
type Client struct {
	conn *rpc.Client
}

func main() {

	if len(os.Args) != 2 {
		panic("Usage: client SERVER_ADDR:SERVER PORT")
	}

	conn, err := rpc.Dial("tcp", os.Args[1])

	if err != nil {
		log.Fatal("Connecting", err)
	}

	client := &Client{conn: conn}

	for i := 0; ; i++ {
		client.Reverse("Hello, world")
		client.ToUpper("Hello, world")
		client.ToLower("Hello, world")

		if i%1000 == 0 {
			fmt.Print(".")
		}
	}
}
