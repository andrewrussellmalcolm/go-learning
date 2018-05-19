package main

import "C"
import (
	"fmt"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

//go build -v -tags gtk_3_18 -gcflags "-N -l"

func main() {

	gtk.Init(nil)

	windowMain, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	//	bailOnError(err)

	windowMain.SetTitle("Plot")
	windowMain.SetDefaultSize(600, 600)
	windowMain.SetBorderWidth(20)
	windowMain.Connect("destroy", func() {
		gtk.MainQuit()
	})

	image, err := gtk.ImageNew()
	if err != nil {
		panic(err)
	}

	windowMain.Container.Add(image)

	windowMain.Container.Connect("configure-event", func(window *gtk.Window, evt *gdk.Event) {

		w, h := window.GetSize()
		fmt.Printf("w=%d h=%d\n", w, h)

	})

	plot, err := plotFromFile("data.txt")
	if err != nil {
		panic(err)
	}

	updateImage(image, plot)
	windowMain.ShowAll()
	gtk.Main()
}

func updateImage(image *gtk.Image, data []byte) {

	loader, err := gdk.PixbufLoaderNew()
	if err != nil {
		panic(err)
	}

	_, err = loader.Write(data)
	if err != nil {
		panic(err)
	}

	pixbuf, err := loader.GetPixbuf()
	if err != nil {
		panic(err)
	}

	image.SetFromPixbuf(pixbuf)

	loader.Close()
}
