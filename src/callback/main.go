package main

import (
	"bufio"
	"callback/informer"
	"fmt"
	"os"
)

func main() {

	i := informer.NewInformer(func(x int) int {

		fmt.Println(x)
		return x
	})

	i.Stop()

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
