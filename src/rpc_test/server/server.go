package main

import (
	//"errors"

	"net"
	"net/rpc"
)

func main() {
	stringOps := new(StringOps)

	server := rpc.NewServer()

	server.RegisterName("StringOps", stringOps)

	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	server.Accept(listener)
}
