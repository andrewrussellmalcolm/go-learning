package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

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
	var frameSizes []webcam.FrameSize
	var formats []webcam.PixelFormat
	for f := range formatDesc {
		formats = append(formats, f)
	}

	println("Available formats: ")
	for i, format := range formats {
		fmt.Fprintf(os.Stderr, "[%d] %s\n", i+1, formatDesc[format])

		frameSizes = cam.GetSupportedFrameSizes(format)

		for i, value := range frameSizes {
			fmt.Fprintf(os.Stderr, "\t[%d] %s\n", i+1, value.GetString())
		}
	}
}

// SelectBestWebcamJpegFormat @:
func selectBestWebcamJpegFormat() error {
	formatDesc := cam.GetSupportedFormats()
	var frameSizes []webcam.FrameSize
	var formats []webcam.PixelFormat
	for f := range formatDesc {
		formats = append(formats, f)
	}

	for _, format := range formats {

		if strings.Contains(strings.ToUpper(formatDesc[format]), "JPEG") {

			frameSizes = cam.GetSupportedFrameSizes(format)

			var height uint32
			var bestIndex int

			for index, size := range frameSizes {

				if size.MaxHeight < height {
					continue
				}
				height = size.MaxHeight
				bestIndex = index
			}

			f, w, h, err := cam.SetImageFormat(format, frameSizes[bestIndex].MaxWidth, frameSizes[bestIndex].MaxHeight)

			if err != nil {
				panic(err.Error())
			} else {
				fmt.Fprintf(os.Stderr, "Selected image format: %s (%dx%d)\n", formatDesc[f], w, h)
			}
			return nil
		}
	}

	return errors.New("no JPEG support from camera")
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
