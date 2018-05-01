package main

import (
	"fmt"
	"time"

	"../ctime/ctime"
)

func main() {

	fmt.Println(ctime.Strftime(time.Now(), "Full output %Y-%m-%d %H:%M:%S"))
	fmt.Println(ctime.Strftime(time.Now(), "Full output alternate %Y-%b-%d %H:%M:%S"))

	fmt.Println(ctime.Strftime(time.Now(), "Veeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeery looooooooooooooooooooooooooooooong %Y-%b-%d %H:%M:%S"))

	fmt.Println(ctime.Strftime(time.Now(), "Today is %d/%m/%Y"))
	fmt.Println(ctime.Strftime(time.Now(), "The time right now is %H:%M:%S"))
}
