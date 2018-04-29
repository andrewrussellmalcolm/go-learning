package main

import (
	"fmt"
	"os"

	"github.com/blackjack/webcam"
)

var cam *webcam.Webcam

func initWebcam() {
	c, err := webcam.Open("/dev/video0")

	if err != nil {
		panic(err.Error())
	}

	cam = c
}

func listWebcamFormats() {
	formatDesc := cam.GetSupportedFormats()
	var formats []webcam.PixelFormat
	for f := range formatDesc {
		formats = append(formats, f)
	}

	println("Available formats: ")
	for i, value := range formats {
		fmt.Fprintf(os.Stderr, "[%d] %s\n", i+1, formatDesc[value])
	}

}

func closeWebcam() {
	cam.Close()
}

func startCapture(frame chan []byte) {

	err := cam.StartStreaming()
	if err != nil {
		panic(err.Error())
	}

	for {
		err = cam.WaitForFrame(5)

		switch err.(type) {
		case nil:
		case *webcam.Timeout:
			fmt.Fprint(os.Stderr, err.Error())
			continue
		default:
			panic(err.Error())
		}

		f, err := cam.ReadFrame()

		if len(f) != 0 {

			frame <- f
		} else if err != nil {
			panic(err.Error())
		}
	}
}
