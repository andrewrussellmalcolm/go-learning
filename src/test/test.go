package main

import (
	"fmt"
	"unsafe"
)

// Test :
type Test struct {
	header [10]byte
	data   [256]byte
}

func main() {

	test := Test{}

	copy(test.header[:], []byte("1234567890"))
	fmt.Printf("size of test %d\n", unsafe.Sizeof(test))

	p := []byte(unsafe.Pointer(&test))

	fmt.Println(p)
}
