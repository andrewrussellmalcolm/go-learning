package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//m := make(map[string]int)
	m := make(map[rune]rune)

	for ch := 'A'; ch <= 'Z'+30; ch++ {
		m[ch] = rune(rand.Intn(255))
	}

	var count int

	for key, value := range m {
		fmt.Printf("Key: %c Value: %c %2x\n", key, value, value)
		count++
	}

	fmt.Println(count)
}
