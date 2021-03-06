package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	dbinfo := fmt.Sprintf("user=postgres host=citus1")

	db, err := sql.Open("postgres", dbinfo)
	bailOnError(err)
	defer db.Close()

	inserts := 1000

	// delete, err := db.Prepare("DELETE FROM test")
	// bailOnError(err)
	insert, err := db.Prepare("INSERT INTO measurement (sensorlocationid,timestamp,value) VALUES($1, $2, $3)")
	bailOnError(err)
	// querySingleRowByID, err := db.Prepare("SELECT * FROM test WHERE id = ($1)")
	// bailOnError(err)
	queryRowCount, err := db.Prepare("SELECT COUNT(*) FROM measurement")
	bailOnError(err)

	// timeOperation(func() {
	// 	_, err := delete.Exec()
	// 	bailOnError(err)
	// }, "DELETE ALL")

	for j := 0; j < 10000; j++ {

		timeOperation(func() {
			for i := 0; i < inserts; i++ {
				_, err := insert.Exec(i, time.Now().Unix(), rand.Float64())
				bailOnError(err)
			}
		}, fmt.Sprintf("INSERT %d ROWS", inserts))

		// timeOperation(func() {
		// 	_, err := querySingleRowByID.Query(12345)
		// 	bailOnError(err)
		// }, "QUERY SINGLE ROW BY ID")

		timeOperation(func() {
			rows, err := queryRowCount.Query()
			bailOnError(err)
			var count int
			rows.Next()
			err = rows.Scan(&count)
			bailOnError(err)
			fmt.Printf("\tCOUNT=%d\n", count)
		}, "QUERY ROW COUNT")
	}
	// timeOperation(func() {
	// 	_, err := delete.Exec()
	// 	bailOnError(err)
	// }, "DELETE ALL")

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
