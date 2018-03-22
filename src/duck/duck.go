package main

import (
	"fmt"
)

type Feeder interface {
	Feed(message string)
}

type Cat struct {
}

type Dog struct {
}

func (cat *Cat) Feed(food string) {
	fmt.Printf("cat was fed %s\n", food)
}

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
