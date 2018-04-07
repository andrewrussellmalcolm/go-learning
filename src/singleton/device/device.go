package device

import "fmt"

type Device struct {
	c int
}

var device Device

func init() {
	device = Device{1}
}

func GetDevice() Device {
	return device
}

func (d Device) Print() {
	fmt.Printf("%v\n", d)
}
