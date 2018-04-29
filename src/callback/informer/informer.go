package informer

import (
	"time"
)

// Informer :
type Informer struct {
	ticker time.Ticker
	x      int
}

// NewInformer :
func NewInformer(inform func(int) int) Informer {

	i := Informer{*time.NewTicker(500 * time.Millisecond), 0}

	go func() {

		for _ = range i.ticker.C {
			i.x = inform(i.x + 1)
		}

	}()
	return i
}

// Stop :
func (i Informer) Stop() {

	i.ticker.Stop()
}
