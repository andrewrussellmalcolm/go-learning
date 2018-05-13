package main

//go build -v -tags gtk_3_18 -gcflags "-N -l"

import (
	"fmt"
	"os"

	"github.com/gotk3/gotk3/gdk"

	"github.com/gotk3/gotk3/cairo"
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

	// custom
	custom, err := CustomNew(DrawCustom)
	custom.SetEvents(custom.GetEvents() | int(gdk.POINTER_MOTION_MASK) | int(gdk.BUTTON_PRESS_MASK))
	custom.Connect("motion-notify-event", func(widget *gtk.DrawingArea, evt *gdk.Event) {

		x, y := gdk.EventMotionNewFromEvent(evt).MotionVal()
		fmt.Printf("MNE %f %f\n", x, y)
	})

	custom.Connect("button-press-event", func(widget *gtk.DrawingArea, evt *gdk.Event) {

		x, y := gdk.EventButtonNewFromEvent(evt).MotionVal()
		fmt.Printf("BPE %f %f\n", x, y)
	})

	// grid
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

// DrawCustom draws th c4 clock
func DrawCustom(w, h float64, ctx *cairo.Context) {

	ctx.SetSourceRGB(0, 0, 0)
	ctx.Rectangle(0, 0, w, h)
	ctx.Fill()

	ctx.SetSourceRGB(1.0, 1.0, 1.0)
	ctx.MoveTo(0, 0)
	ctx.LineTo(w, h)
	ctx.Stroke()

	ctx.MoveTo(w, 0)
	ctx.LineTo(0, h)
	ctx.Stroke()

}

func bailOnError(err error) {
	if err != nil {
		panic(err)
	}
}
