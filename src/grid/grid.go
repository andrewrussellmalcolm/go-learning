package main

// sudo apt-get install libgtk-3-dev
//go build -v -tags gtk_3_18 -gcflags "-N -l"

import (
	"os"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func main() {

	gtk.Init(&os.Args)
	windowMain, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	bailOnError(err)

	windowMain.SetTitle("Controller")
	windowMain.SetDefaultSize(600, 300)
	//windowMain.SetPosition(gtk.WIN_POS_CENTER)
	windowMain.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// some buttons
	buttonStart, err := gtk.ButtonNewWithLabel("Start")
	buttonStop, err := gtk.ButtonNewWithLabel("Stop")
	buttonPanic, err := gtk.ButtonNewWithLabel("PANIC!!")

	custom, err := CustomNew()

	custom.SetEvents(custom.GetEvents() | int(gdk.POINTER_MOTION_MASK) | int(gdk.BUTTON_PRESS_MASK) | int(gdk.BUTTON_RELEASE_MASK))
	custom.Connect("draw", custom.DrawCustom)
	custom.Connect("motion-notify-event", custom.MotionEvent)
	custom.Connect("button-press-event", custom.ButtonEvent, true)
	custom.Connect("button-release-event", custom.ButtonEvent, false)

	grid, err := gtk.GridNew()
	bailOnError(err)
	grid.SetBorderWidth(10)
	grid.SetRowSpacing(10)
	grid.SetColumnSpacing(10)
	grid.SetColumnHomogeneous(true)
	grid.Attach(buttonStart, 0, 0, 1, 1)
	grid.Attach(buttonStop, 0, 1, 1, 1)
	grid.Attach(buttonPanic, 0, 2, 1, 1)

	// vertical split pane
	pane, err := gtk.PanedNew(gtk.ORIENTATION_HORIZONTAL)
	pane.SetBorderWidth(10)
	pane.Add1(grid)
	pane.Add2(custom)

	windowMain.Add(pane)
	windowMain.ShowAll()

	gtk.Main()
}

func bailOnError(err error) {
	if err != nil {
		panic(err)
	}
}
