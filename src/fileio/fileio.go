package main

import (
	"bufio"
	"encoding/json"
	"os"
)

type Template struct {
	Data1 int64
	Data2 string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	template := Template{4, "Hello"}

	file, err := os.Create("data.json")

	check(err)

	w := bufio.NewWriter(file)

	err = json.NewEncoder(w).Encode(template)

	check(err)

	w.Flush()
}
