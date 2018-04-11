package main

import (
	"flag"
	"fmt"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	fmt.Printf("%v\n", *addr)
}
