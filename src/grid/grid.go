package main

import "C"
import (
	"github.com/gotk3/gotk3/gtk"
)

func main() {

	gtk.Init(nil)

	windowMain, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	bailOnError(err)

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	bailOnError(err)
	windowMain.Container.Add(box)

	windowMain.SetTitle("Controller")
	windowMain.SetDefaultSize(600, 300)
	windowMain.SetPosition(gtk.WIN_POS_CENTER)
	windowMain.Connect("destroy", func() {
		gtk.MainQuit()
	})

	grid, err := gtk.GridNew()
	bailOnError(err)
	box.PackStart(grid, false, false, 0)

	buttonStart, err := gtk.ButtonNewWithLabel("Start")
	buttonStop, err := gtk.ButtonNewWithLabel("Stop")
	buttonPanic, err := gtk.ButtonNewWithLabel("PANIC!!")

	grid.SetBorderWidth(10)
	grid.SetRowSpacing(10)
	grid.SetColumnSpacing(10)
	grid.SetColumnHomogeneous(true)
	grid.Attach(buttonStart, 0, 0, 1, 1)
	grid.Attach(buttonStop, 3, 0, 1, 1)
	grid.Attach(buttonPanic, 0, 1, 4, 1)

	grid.Attach(drawingArea, 2, 2, 3, 3)

	gtk.Main()
}
