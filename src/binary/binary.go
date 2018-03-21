package main

import (
    "bytes"
	"encoding/binary"
	"fmt"
)

/**  */
func main() {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, []byte("12345678"))
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.Bytes())
}
