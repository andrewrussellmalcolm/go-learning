package main

import "C"
import (
	"fmt"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

//go build -v -tags gtk_3_18 -gcflags "-N -l"

func main() {

	gtk.Init(nil)

	windowMain, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	bailOnError(err)

	windowMain.SetTitle("Radar")
	windowMain.SetDefaultSize(600, 600)
	windowMain.Connect("destroy", func() {
		gtk.MainQuit()
	})

	custom, err := CustomNew()

	custom.Connect("draw", custom.DrawCustom)

	windowMain.Container.Add(custom)

	windowMain.ShowAll()

	ticker := time.NewTicker(time.Millisecond * 10)

	custom.QueueDraw()

	go func() {

		for range ticker.C {
			custom.QueueDraw()
		}
	}()

	fmt.Println("gtk.Main")
	gtk.Main()
}

func bailOnError(err error) {
	if err != nil {
		panic(err)
	}
}
