package main

import (
	"time"
)

type informer struct {
	ticker time.Ticker
	x      int
}

func NewInformer(inform func(int) int) informer {

	i := informer{*time.NewTicker(500 * time.Millisecond), 0}

	go func() {

		for _ = range i.ticker.C {
			i.x = inform(i.x + 1)
		}

	}()
	return i
}
