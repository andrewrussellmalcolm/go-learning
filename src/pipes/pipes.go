package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	ticker := time.NewTicker(200 * time.Millisecond)

	pipeReader, pipeWriter, _ := os.Pipe()

	go func() {

		for {
			buffer := make([]byte, 64)
			n, _ := pipeReader.Read(buffer)

			fmt.Print(string(buffer[:n]))
		}
	}()

	go func() {

		for range ticker.C {
			pipeWriter.WriteString("hello ")
			pipeWriter.WriteString("Andrew Malcolm ")
		}
	}()

	// wait for control-C signal
	<-stop
	ticker.Stop()
}
