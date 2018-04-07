package main

import (
	"singleton/device"
)

func main() {

	device := device.GetDevice()

	device.Print()
}
