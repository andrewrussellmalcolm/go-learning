package main

import (
	"fmt"
)

// Feeder :
type Feeder interface {
	Feed(message string)
}

// Cat :
type Cat struct {
}

// Dog :
type Dog struct {
}

// Feed :
func (cat *Cat) Feed(food string) {
	fmt.Printf("cat was fed %s\n", food)
}

// Feed :
func (dog *Dog) Feed(food string) {
	fmt.Printf("dog was fed %s\n", food)
}

func main() {

	dogFeeder := Feeder(&Dog{})
	catFeeder := Feeder(&Cat{})

	feeders := []Feeder{dogFeeder, catFeeder}

	for _, feeder := range feeders {
		feeder.Feed("meat")
	}
}
