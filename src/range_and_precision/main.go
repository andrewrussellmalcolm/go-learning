package main

import (
	"fmt"
	"math"
)

func main() {

	oneTeraWatt := float64(1e12)
	oneGigaWatt := float64(1e9)
	oneMegaWatt := float64(1e6)
	oneKiloWatt := float64(1e3)
	oneWatt := float64(1)
	oneMilliwatt := float64(1e-3)
	oneMicroWatt := float64(1e-6)
	oneNanoWatt := float64(1e-9)
	onePicoWatt := float64(1e-12)
	oneFemtoWatt := float64(1e-15)
	oneAttoWatt := float64(1e-18)
	oneZeptoWatt := float64(1e-21)
	oneYoctoWatt := float64(1e-24)

	aLotOfWatts := oneTeraWatt + oneGigaWatt + oneMegaWatt + oneKiloWatt + oneWatt + oneMilliwatt +
		oneMicroWatt + oneNanoWatt + onePicoWatt + oneFemtoWatt + oneAttoWatt + oneZeptoWatt + oneYoctoWatt

	fmt.Printf("aLotOfWatts=%.20F\n", aLotOfWatts)
	fmt.Printf("------------T\n")
	fmt.Printf("---------------G\n")
	fmt.Printf("------------------M\n")
	fmt.Printf("---------------------K\n")
	fmt.Printf("------------------------U\n")
	fmt.Printf("---------------------------m\n")
	fmt.Printf("------------------------------u\n")
	fmt.Printf("---------------------------------n\n")

	fmt.Printf("52 bits              = %.0f\n", math.Pow(2, 52))
	fmt.Printf("5 orders of magntude = %.0f\n", math.Pow(10, 15))
}
