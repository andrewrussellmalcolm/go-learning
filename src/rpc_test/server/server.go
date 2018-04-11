package main

import (
	//"errors"

	"net"
	"net/rpc"
)

func main() {

	server := rpc.NewServer()

	server.Register(new(StringOps))

	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	server.Accept(listener)
}
