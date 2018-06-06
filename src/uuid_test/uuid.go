package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {

	uuid := uuid.New()

	fmt.Printf("%v\n", uuid)

	b, err := uuid.MarshalBinary()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d %d\n", len(uuid.String()), len(b))

}
