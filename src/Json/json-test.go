package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Person :
type Person struct {
	Forename string
	Surname  string
}

func main() {

	me := Person{Forename: "Andrew", Surname: "Malcolm"}

	f, _ := os.Create("test.json")

	w := bufio.NewWriter(f)

	json.NewEncoder(w).Encode(me)

	w.Flush()

	me = Person{}

	r := bufio.NewReader(f)

	json.NewDecoder(r).Decode(me)

	fmt.Print(me)

	
}
