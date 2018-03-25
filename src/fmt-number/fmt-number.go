package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	p := message.NewPrinter(language.English)

	p.Printf("%d\n", 1)
	p.Printf("%d\n", 10)
	p.Printf("%d\n", 100)
	p.Printf("%d\n", 1000)
	p.Printf("%d\n", 10000)
	p.Printf("%d\n", 100000)
	p.Printf("%d\n", 1000000)
	p.Printf("%d\n", 10000000)
	p.Printf("%d\n", 100000000)
	p.Printf("%d\n", 1000000000)
	p.Printf("%d\n", 10000000000)
	p.Printf("%d\n", 100000000000)
	p.Printf("%d\n", 1000000000000)
	p.Printf("%d\n", 10000000000000)
}
