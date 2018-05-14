package main

import (
	"fmt"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Point struct {
	x, y float64
}

type Custom struct {
	*gtk.DrawingArea
	points []Point
}

func CustomNew() (*Custom, error) {

	da, err := gtk.DrawingAreaNew()

	if err != nil {
		return nil, err
	}
	return &Custom{da, nil}, nil
}

func (c *Custom) MotionEvent(custom *gtk.DrawingArea, evt *gdk.Event) {

	x, y := gdk.EventMotionNewFromEvent(evt).MotionVal()
	fmt.Printf("MNE %f %f\n", x, y)

	fmt.Printf("%d\n", gdk.EventButtonNewFromEvent(evt).Button())

	c.points = append(c.points, Point{x, y})
	c.QueueDraw()
}

func (c *Custom) ButtonEvent(widget *gtk.DrawingArea, evt *gdk.Event) {

	x, y := gdk.EventButtonNewFromEvent(evt).MotionVal()
	fmt.Printf("BPE %f %f\n", x, y)

	fmt.Printf("%d\n", gdk.EventButtonNewFromEvent(evt).Button())
	c.points = nil
	c.QueueDraw()
}

func (c *Custom) DrawCustom(custom *gtk.DrawingArea, ctx *cairo.Context) {

	w := float64(c.GetAllocatedWidth())
	h := float64(c.GetAllocatedHeight())

	ctx.SetSourceRGB(0, 0, 0)
	ctx.Rectangle(0, 0, w, h)
	ctx.Fill()

	ctx.SetSourceRGB(1.0, 1.0, 1.0)
	// ctx.MoveTo(0, 0)
	// ctx.LineTo(w, h)
	// ctx.Stroke()

	// ctx.MoveTo(w, 0)
	// ctx.LineTo(0, h)
	// ctx.Stroke()

	for _, p := range c.points {

		ctx.Rectangle(p.x-3, p.y-3, 6, 6)
		ctx.Fill()
	}
}
