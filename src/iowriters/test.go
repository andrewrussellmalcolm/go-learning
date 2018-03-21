package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// Test :
type Test struct {
	Forename string
	Surname  string
	secret   string
}

// UpperWriter :
type UpperWriter struct{}

func (writer UpperWriter) Write(p []byte) (n int, err error) {
	return fmt.Print(strings.ToUpper(string(p)))
}

// LowerWriter :
type LowerWriter struct{}

func (writer LowerWriter) Write(p []byte) (n int, err error) {
	return fmt.Print(strings.ToLower(string(p)))
}

func main1() {

	d := Test{"andrew", "Malcolm", "nuts"}

	w := io.MultiWriter(UpperWriter{}, LowerWriter{}, os.Stdout, os.Stderr)

	json.NewEncoder(w).Encode(d)
	r := strings.NewReader("You'll see this string several times!!\n")
	io.Copy(w, r)

}
