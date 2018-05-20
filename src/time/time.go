package main

import (
	"fmt"
	"time"
)

var baseTime = time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)

func main() {

	fmt.Printf("%s\n", baseTime.String())
	fmt.Printf("%d\n", baseTime.Unix())

	fmt.Printf("%d\n", time.Now().Unix())

	nowOffset := time.Now().Unix() - baseTime.Unix()
	fmt.Printf("0x%x\n", nowOffset)

	fmt.Printf("%s\n", time.Unix(baseTime.Unix()+nowOffset, 0).String())

}
