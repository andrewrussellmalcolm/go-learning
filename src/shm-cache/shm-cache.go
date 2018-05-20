package main

import (
	"fmt"
	"log"
	ce "shm-cache/cache"
	"time"
)

/** */
func main() {

	var cache ce.Cache

	timeOperation(func() {
		err := ce.InitCache(&cache)
		if err != nil {
			log.Fatalf("failed to allocate entry array%v", err)
		}

	}, fmt.Sprintf("creating cache"))

	defer ce.DeinitCache(&cache)

	timeOperation(func() {

		stripe := cache.GetStripe(1)

		for j := range stripe.GetEntries() {

			stripe.GetEntry(j).SetLocationID(uint32(j))
			stripe.GetEntry(j).SetMeasurement(uint32(j))
			stripe.GetEntry(j).SetTimeOffset(uint32(j))
		}

	}, fmt.Sprintf("initialising single cache stripe"))

	timeOperation(func() {

		stripe := cache.GetStripe(1)

		for j := range stripe.GetEntries() {

			if stripe.GetEntry(j).GetLocationID() == 100 {

				fmt.Printf("%v\n", stripe.GetEntry(j))
			}
		}

	}, fmt.Sprintf("searching single cache stripe for a locationID"))

	timeOperation(func() {

		entries1 := cache.GetStripe(1).GetEntries()
		entriesX := make([]ce.Entry, len(entries1))

		fmt.Printf("entry %d LocationID%d\n", 100, entries1[100])

		// copy - modify -replace
		copy(entriesX, entries1)
		entriesX[100].SetLocationID(555)
		copy(entries1, entriesX)

		fmt.Printf("entry %d LocationID%d\n", 100, entries1[100])

	}, fmt.Sprintf("copy - modify - replace a cache stripe"))

	timeOperation(func() {

		var locationID = int32(0x10007700)

		for i := 0; i < len(cache.GetStripes()); i++ {

			if byte(i) == byte(locationID>>8) {
				fmt.Printf("locationID is 0x%x in cache entry 0x%x\n", locationID, i)
			}
		}

		//fmt.Printf("entry %d LocationID %d\n", 100, entryCache[1][100])

	}, fmt.Sprintf("locate cache entry for LocationID"))
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
