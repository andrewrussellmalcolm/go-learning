package main

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Point struct {
	x, y  float64
	color *gdk.RGBA
}

type Custom struct {
	*gtk.DrawingArea
	points []Point
	draw   bool
	color  *gdk.RGBA
}

func CustomNew() (*Custom, error) {

	da, err := gtk.DrawingAreaNew()

	if err != nil {
		return nil, err
	}
	return &Custom{da, nil, false, gdk.NewRGBA(1, 1, 1, 1)}, nil
}

func (c *Custom) MotionEvent(custom *gtk.DrawingArea, evt *gdk.Event) {

	x, y := gdk.EventMotionNewFromEvent(evt).MotionVal()

	if c.draw {
		c.points = append(c.points, Point{x, y, c.color})
		c.QueueDraw()
	}
}

func (c *Custom) ButtonEvent(widget *gtk.DrawingArea, evt *gdk.Event) {

	if gdk.EventButtonNewFromEvent(evt).Button() == 1 && gdk.EventButtonNewFromEvent(evt).State() == 0 {

		c.draw = true
		c.QueueDraw()
	}

	if gdk.EventButtonNewFromEvent(evt).Button() == 1 && gdk.EventButtonNewFromEvent(evt).State() != 0 {

		c.draw = false
	}

	if gdk.EventButtonNewFromEvent(evt).Button() == 3 && gdk.EventButtonNewFromEvent(evt).State() == 0 {

		c.points = nil

		c.QueueDraw()
	}
}

func (c *Custom) DrawCustom(custom *gtk.DrawingArea, ctx *cairo.Context) {

	w := float64(c.GetAllocatedWidth())
	h := float64(c.GetAllocatedHeight())

	ctx.SetSourceRGB(0, 0, 0)
	ctx.Rectangle(0, 0, w, h)
	ctx.Fill()

	ctx.SetSourceRGB(0, 0, 0)
	ctx.Rectangle(0, 0, w, h)
	ctx.Fill()

	for _, p := range c.points {

		rgb := p.color.Floats()
		ctx.SetSourceRGB(rgb[0], rgb[1], rgb[2])

		ctx.Rectangle(p.x-3, p.y-3, 6, 6)
		ctx.Fill()
	}
}
