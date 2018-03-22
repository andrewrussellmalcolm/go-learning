package main

import (
	"fmt"
	"time"
)

func ternary(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func main() {

	a := true

	b := true

	c := ternary(a == b, 100, 500)

	fmt.Println(a, b, c)

	d := ternary(a == b, time.Now(), time.Now())

	fmt.Println(d)

	e := ternary(a != b, 0x100, 0x200)

	fmt.Println(e)

	f := ternary(a != b, 'A', 'B')

	fmt.Printf("%c\n", f)
}
