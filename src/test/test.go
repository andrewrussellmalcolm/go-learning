package main

import (
	/*
		        #include <stdio.h>
		        #include <stdlib.h>
				 #include "pcap.h"
				#cgo LDFLAGS: /usr/lib/x86_64-linux-gnu/libpcap.so
	*/
	"C"
)
import (
	"fmt"
	"log"
)

func main() {

	type PcapIf *C.struct_pcap_if
	type PcapPktHdr *C.struct_pcap_pkthdr

	var devlist PcapIf
	var lanDev PcapIf

	res := C.pcap_findalldevs(&devlist, nil)

	if res != 0 {
		log.Fatalf("pcap_findalldevs returned %d", res)
	}

	for dev := devlist; dev != nil; dev = dev.next {

		name := C.GoString(dev.name)

		if name == "enp5s0" {
			lanDev = dev
			break
		}
	}

	fmt.Printf("%s\n", C.GoString(lanDev.name))

	var err string

	handle := C.pcap_open_live(lanDev.name, 1000, 1, 1000, C.CString(err))

	if handle == nil {
		log.Fatalf("pcap_open_live returned %s", err)
	}

	pktHdr := new(C.struct_pcap_pkthdr)

	pktBuf := C.malloc(10000)

	for {

		C.pcap_next_ex(handle, &pktHdr, pktBuf)

		fmt.Printf("len=%d\n", pktHdr.len)
	}

	C.pcap_close(handle)
}
