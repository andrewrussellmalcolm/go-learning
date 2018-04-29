package main

import (
	"fmt"
	"time"
)

func main() {

	observable := NewObservable()

	observable.RegisterObserver("Tick observer 1", func(tick int) {

		fmt.Println(tick)
	})

	observable.RegisterObserver("Tick observer 2", func(tick int) {

		fmt.Println("\t", tick)
	})
	time.Sleep(10 * time.Second)
}
