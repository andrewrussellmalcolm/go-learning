package main

import "fmt"

func willError() (int, int, error) {

	return 0, 1, nil

}

func wontError() (int, int, error) {

	return 0, 1, nil

}
func main() {

	i, j, err := check(mightError())

	fmt.Printf("%d %d,%v\n", i, j, err)
}

func check(args ...interface{}) (... ret) {

	err := args[len(args)-1]
	if err != nil {
		panic(err)
	}

}
