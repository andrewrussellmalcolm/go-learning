package main

import (
	"math"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

// Custom is a custom widget
type Custom struct {
	*gtk.DrawingArea
	sweep float64
}

// CustomNew creates a new widget
func CustomNew() (*Custom, error) {

	drawingArea, err := gtk.DrawingAreaNew()

	if err != nil {
		return nil, err
	}

	return &Custom{drawingArea, 0.0}, nil
}

/** */
func (c *Custom) DrawCustom(custom *gtk.DrawingArea, ctx *cairo.Context) {

	w := float64(c.GetAllocatedWidth())
	h := float64(c.GetAllocatedHeight())

	// fil background
	ctx.SetSourceRGB(0, 0, 0)
	ctx.Rectangle(0, 0, w, h)
	ctx.Fill()

	// calculate center and radius
	x0 := w / 2
	y0 := h / 2
	r0 := math.Min(x0, y0)

	ctx.SetSourceRGB(0, 0.5, 0)
	ctx.SetLineWidth(2)
	ctx.MoveTo(0, y0)
	ctx.LineTo(w, y0)
	ctx.Stroke()
	ctx.MoveTo(x0, 0)
	ctx.LineTo(x0, h)
	ctx.Stroke()

	for r := r0 / 6; r < r0; r += r0 / 6 {
		ctx.Arc(x0, y0, r, 0, math.Pi*2)
		ctx.Stroke()
	}

	ctx.SetSourceRGB(0, 1, 0)
	ctx.SetLineWidth(3)
	x1 := x0 + r0*math.Sin(c.sweep)*0.95
	y1 := y0 + r0*math.Cos(c.sweep)*0.95
	ctx.MoveTo(x0, y0)
	ctx.LineTo(x1, y1)
	ctx.Stroke()

	ctx.Save()
	ctx.Arc(x0, y0, r0, 0, math.Pi/15)
	ctx.ClosePath()
	ctx.SetSourceRGBA(0, 0, 0.8, 0.6)
	ctx.FillPreserve()
	ctx.Restore()
	ctx.Stroke()

	c.sweep -= math.Pi / 300
}
