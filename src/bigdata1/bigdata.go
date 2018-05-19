package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

func main() {

	data := make(map[int]string)

	inserts := 100000
	for j := 0; j < 10000; j++ {
		timeOperation(func() {
			for i := 0; i < inserts; i++ {
				u1 := uuid.Must(uuid.NewV4())
				data[i*j] = u1.String()

			}
		}, fmt.Sprintf("INSERT %d ROWS", inserts))

		timeOperation(func() {
			fmt.Printf("\t%s\n", data[12345])
		}, "QUERY SINGLE ROW BY ID")

		timeOperation(func() {
			count := len(data)
			fmt.Printf("\tCOUNT=%d\n", count)
		}, "QUERY ROW COUNT")
	}

	os.Exit(0)
}

func bailOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func timeOperation(operation func(), statement string) {

	t0 := time.Now()
	operation()
	t1 := time.Now()

	time := float32(t1.Sub(t0).Nanoseconds())

	if time < 1000 {
		fmt.Printf("[%s] took %.1fns\n", statement, time)
	} else if time < 1000000 {
		fmt.Printf("[%s] took %.1fus\n", statement, time/1000.0)
	} else if time < 1000000000 {
		fmt.Printf("[%s] took %.1fms\n", statement, time/1000000.0)
	} else if time < 1000000000000 {
		fmt.Printf("[%s] took %.1fs\n", statement, time/1000000000.0)
	}
}
