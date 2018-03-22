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

	fmt.Println(client.Reverse("Hello, world"))
	fmt.Println(client.ToUpper("Hello, world"))
	fmt.Println(client.ToLower("Hello, world"))
}
