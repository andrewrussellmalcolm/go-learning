package main

import (
	"fmt"
	"time"
)

func main() {

	timeout := make(chan bool, 1)
	channel := make(chan bool, 1)
	go func() {
		time.Sleep(500 * time.Millisecond)
		channel <- true
		time.Sleep(500 * time.Millisecond)
		timeout <- true
	}()

	select {
	case <-channel:
		fmt.Println("read")
	case <-timeout:
		fmt.Println("timeout")

	}

}
