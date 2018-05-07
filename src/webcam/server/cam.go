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

func startStreaming() error {

	err := cam.StartStreaming()
	if err != nil {
		return err
	}

	return nil
}

func waitForFrame() error {

	err := cam.WaitForFrame(5)
	if err != nil {
		return err
	}

	return nil
}

func getFrame() ([]byte, error) {

	f, err := cam.ReadFrame()

	if err != nil {
		return nil, err
	}

	return f, nil
}
