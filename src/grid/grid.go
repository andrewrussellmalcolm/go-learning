package main

// sudo apt-get install libgtk-3-dev
//go build -v -tags gtk_3_18 -gcflags "-N -l"

import (
	"fmt"
	"os"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func main() {

	gtk.Init(&os.Args)
	windowMain, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	bailOnError(err)

	windowMain.SetTitle("Controller")
	windowMain.SetDefaultSize(1600, 800)
	//windowMain.SetPosition(gtk.WIN_POS_CENTER)
	windowMain.Connect("destroy", func() {
		gtk.MainQuit()
	})

	custom, err := CustomNew()

	custom.SetEvents(custom.GetEvents() | int(gdk.POINTER_MOTION_MASK) | int(gdk.BUTTON_PRESS_MASK) | int(gdk.BUTTON_RELEASE_MASK))
	custom.Connect("draw", custom.DrawCustom)
	custom.Connect("motion-notify-event", custom.MotionEvent)
	custom.Connect("button-press-event", custom.ButtonEvent, true)
	custom.Connect("button-release-event", custom.ButtonEvent, false)

	toolbar, err := gtk.ToolbarNew()
	toolbar.SetStyle(gtk.TOOLBAR_TEXT)

	colorButton, err := gtk.ToolButtonNew(nil, "Colour")
	colorButton.Connect("clicked", func() {
		colorChooser, err := gtk.ColorChooserDialogNew("Choose a color", windowMain)
		if err != nil {
			panic(err)
		}

		colorChooser.Connect("color-activated", func() {
		})
		response := gtk.ResponseType(colorChooser.Run())

		fmt.Printf("respone %d\n", response)

		switch response {
		case gtk.RESPONSE_OK:
			fmt.Printf("color chosen %v\n", colorChooser.GetRGBA())

			custom.color = colorChooser.GetRGBA()
		case gtk.RESPONSE_CANCEL:
		}
		colorChooser.Close()
	})
	toolbar.Insert(colorButton, 0)

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	bailOnError(err)
	windowMain.Container.Add(box)
	box.PackStart(toolbar, false, false, 0)
	box.PackStart(custom, true, true, 0)

	windowMain.ShowAll()

	gtk.Main()
}

func bailOnError(err error) {
	if err != nil {
		panic(err)
	}
}
