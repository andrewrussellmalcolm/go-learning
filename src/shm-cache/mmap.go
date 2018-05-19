package main

//#include <sys/mman.h>
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

func mmap(file *os.File, length int64) (mm []byte, err error) {

	fmt.Printf("%x\n", unsafe.Sizeof(uintptr(length)))

	//addr, x, e1 := syscall.RawSyscall6(syscall.SYS_MMAP, 0, uintptr(length), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED, file.Fd(), 0)

	addr := C.mmap(unsafe.Pointer(nil), C.ulong(length), C.PROT_READ|C.PROT_WRITE, C.MAP_SHARED, C.int(file.Fd()), 0)

	fmt.Printf("0x%x\n", length)
	// if x == 0 {
	// 	return nil, fmt.Errorf("mmap failed with errno %s", e1)go build
	// }
	// Slice memory layout
	var sl = struct {
		addr uintptr
		len  int64
		cap  int64
	}{uintptr(addr), length, length}

	// Use unsafe to turn sl into a []byte.
	return *(*[]byte)(unsafe.Pointer(&sl)), nil
}
