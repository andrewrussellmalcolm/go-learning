package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const GB = (1024 * 1024 * 1024)

var arraySizeFlag = flag.Int64("size", 8, "size of the array in GBytes")

/** */
func main() {
	flag.Parse()

	arraySize := *arraySizeFlag * GB
	var f *os.File

	fmt.Printf("Array Size = %dG\n", arraySize/(1024*1024*1024))

	f, err := os.OpenFile("/dev/shm/bf-cache", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("failed to open cache memory file %v", err)
	}

	defer f.Close()

	f.Seek(arraySize, 0)
	f.Write([]byte("bf-cache 1.0"))

	mm, err := mmap(f, arraySize)
	if err != nil {
		log.Fatalf("mapping shared memory %v", err)
	}
	fmt.Println("Starting test")
	timeOperation(func() {
		for i := int64(0); i < (arraySize); i++ {

			//binary.BigEndian.PutUint64(mm[i:i+8], 0x1111222233334444)

			//fmt.Printf("0x%x\n", i+8)
			mm[i] = 'A'
		}
	}, fmt.Sprintf("wrote 0x%x int64s", arraySize/8))

	// binary.BigEndian.PutUint64(mm[0:8], 0)
	// mm[arraySize-1] = 'P'
	//binary.BigEndian.PutUint64(mm[arraySize-8:arraySize], 0)

}

/** */
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

// writeMeasurement(uint64(0), 0x1111222233334444, 0x5555666677778888, mmap)

// v, t := readMeasurement(uint64(0), mmap)

// fmt.Printf("0x%x 0x%x\n", v, t)
// mmap.Flush()
