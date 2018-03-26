package main

import (
	"fmt"
	"strings"
)

func main() {

	line := "s andrew  this is a test"

	var cmd rune
	var user string

	_, err := fmt.Sscanf(line, "%c %s", &cmd, &user)

	if err != nil {
		panic(err)
	}

	offset := strings.LastIndex(line, user) + len(user) + 2
	fmt.Println(cmd, user, line[offset:])
}
