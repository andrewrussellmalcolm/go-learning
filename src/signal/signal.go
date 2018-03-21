package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	defer fmt.Println("\nStopped by signal")

	ticker := time.NewTicker(200 * time.Millisecond)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		for range ticker.C {
			fmt.Print(".")
		}
	}()

	// wait for control-C signal
	<-stop
	ticker.Stop()

}
