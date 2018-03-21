package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/exp/shiny/widget"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

// TODO: scrolling, such as when images are larger than the window.
func decode(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	m, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("could not decode %s: %v", filename, err)
	}
	return m, nil
}
func main() {
	log.SetFlags(0)
	driver.Main(func(s screen.Screen) {
		if len(os.Args) < 2 {
			log.Fatal("no image file specified")
		}
		// TODO: view multiple images.
		src, err := decode(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		w := widget.NewSheet(widget.NewImage(src, src.Bounds()))
		if err := widget.RunWindow(s, w, &widget.RunWindowOptions{
			NewWindowOptions: screen.NewWindowOptions{
				Title: "ImageView Shiny Example",
			},
		}); err != nil {
			log.Fatal(err)
		}
	})
}
