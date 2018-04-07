package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	i := NewInformer(func(x int) int {

		fmt.Println(x)
		return x
	})

	i.ticker.Stop()

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
