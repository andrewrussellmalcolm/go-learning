package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"time"
)

// Client :
type Client struct {
	conn *rpc.Client
}

var addr = flag.String("addr", "localhost:1234", "service address")

func main() {

	flag.Parse()

	conn, err := rpc.Dial("tcp", *addr)

	if err != nil {
		log.Fatal("Connecting ", err)
	}

	client := &Client{conn: conn}

	const messages = 1000
	for {

		start := time.Now()
		for message := 0; message < messages; message++ {

			client.Reverse("Hello, world")
		}

		fmt.Printf("%.0f messages per second\r", messages/time.Now().Sub(start).Seconds())
	}
}
