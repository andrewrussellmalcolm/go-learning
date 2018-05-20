//
// Package cache : a cache for measurement entries backed by shared memory files
//
package cache

//#include <errno.h>
//#include <sys/mman.h>
import "C"
import (
	"fmt"
	"log"
	"os"
	"unsafe"
)

const stripeEntries = 1024 * 1024 * 4
const stripes = 256

// Cache :
type Cache struct {
	stripes []Stripe
}

// Stripe :
type Stripe struct {
	entries        []Entry
	insertPosition int
	rawAddress     unsafe.Pointer
	rawSize        C.ulong
}

// Entry :
type Entry struct {
	locationID  uint32
	timeOffset  uint32
	measurement uint32
}

// SetLocationID :
func (e *Entry) SetLocationID(locationID uint32) {
	e.locationID = locationID
}

// SetTimeOffset :
func (e *Entry) SetTimeOffset(timeOffset uint32) {
	e.timeOffset = timeOffset
}

// SetMeasurement :
func (e *Entry) SetMeasurement(measurement uint32) {
	e.measurement = measurement
}

// GetLocationID :
func (e *Entry) GetLocationID() uint32 {
	return e.locationID
}

// GetTimeOffset :
func (e *Entry) GetTimeOffset() uint32 {
	return e.timeOffset
}

// GetMeasurement :
func (e *Entry) GetMeasurement() uint32 {
	return e.measurement
}

// GetStripes :
func (c *Cache) GetStripes() []Stripe {
	return c.stripes
}

// GetStripe :
func (c *Cache) GetStripe(i int) *Stripe {
	return &c.stripes[i]
}

// GetEntries :
func (s *Stripe) GetEntries() []Entry {
	return s.entries
}

// GetEntry :
func (s *Stripe) GetEntry(i int) *Entry {
	return &s.entries[i]
}

// GetLength :
func (s *Stripe) GetLength() int {
	return len(s.entries)
}

// InitCache :
func InitCache(cache *Cache) error {
	cache.stripes = make([]Stripe, stripes)
	for i := 0; i < stripes; i++ {
		stripe, err := createStripe(i, stripeEntries)
		if err != nil {
			return err
		}
		cache.stripes[i] = stripe
	}
	return nil
}

// DeinitCache :
func DeinitCache(cache *Cache) error {
	for i := 0; i < stripes; i++ {
		_, err := C.munmap(cache.GetStripe(i).rawAddress, cache.GetStripe(i).rawSize)
		if err != nil {
			return err
		}
	}

	return nil
}

/*
this requres some configuration of the OS to make resources available

to enable locking of memory:

IN /etc/security/limits.conf
*        	hard   memlock       	unlimited
*        	soft    memlock       	unlimited

to increase size of tempfs
sudo mount -o remount,size=12G /dev/shm
*/
func createStripe(index, length int) (Stripe, error) {

	stripe := Stripe{}
	var e Entry
	stripe.rawSize = C.ulong(length * int(unsafe.Sizeof(e)))
	//fmt.Printf("Array size in entries= %d\n", length)
	//fmt.Printf("Array size in bytes= %d\n", arraySize)

	// open or create a shared memory file for the stripe
	f, err := os.OpenFile(fmt.Sprintf("/dev/shm/cache-stripe-%04d", index), os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("failed to open cache memory file %v", err)
	}

	// map the shared memory
	stripe.rawAddress, err = C.mmap(unsafe.Pointer(nil), stripe.rawSize, C.PROT_READ|C.PROT_WRITE, C.MAP_SHARED, C.int(f.Fd()), 0)
	if err != nil {
		return stripe, fmt.Errorf("mmap failed %v", err)
	}

	// lock it into main memory, no swapping
	_, err = C.mlock(stripe.rawAddress, stripe.rawSize)
	if err != nil {
		return stripe, fmt.Errorf("mlock failed %v", err)
	}

	// build an array of entries from the raw bytes
	var sl = struct {
		addr uintptr
		len  int
		cap  int
	}{uintptr(stripe.rawAddress), length, length}

	// add the entry arry to the cache stripe
	stripe.entries = *(*[]Entry)(unsafe.Pointer(&sl))

	// return initialised stripe
	return stripe, nil
}
