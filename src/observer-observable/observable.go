package main

import (
	"time"
)

// Observable :
type Observable struct {
	obervers map[string]func(int)
	tick     int
}

// NewObservable :
func NewObservable() Observable {

	observable := Observable{make(map[string]func(int)), 0}
	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {

		for range ticker.C {

			for _, observer := range observable.obervers {
				observer(observable.tick)
			}
			observable.tick++
		}
	}()

	return observable
}

// RegisterObserver :
func (o Observable) RegisterObserver(name string, observer func(int)) {

	o.obervers[name] = observer
}
