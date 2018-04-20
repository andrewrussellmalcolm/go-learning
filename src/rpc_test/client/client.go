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

type Address struct {
	addr string
	set  bool
}

func (address *Address) Set(addr string) error {
	address.addr = addr
	address.set = true
	return nil
}

func (address *Address) String() string {
	return address.addr
}

var address Address

func init() {
	flag.Var(&address, "s", "service address e.g. localhost:1234")
}

func main() {

	flag.Parse()

	if !address.set {
		log.Fatal("No adress set")
	}

	conn, err := rpc.Dial("tcp", address.addr)

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
