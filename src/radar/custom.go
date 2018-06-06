package main

import (
	"math"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

type Point struct {
	x, y float64
}

type Object struct {
	point Point
	b     float64
}

// Custom is a custom widget
type Custom struct {
	*gtk.DrawingArea
	sweep   float64
	objects []*Object
}

// CustomNew creates a new widget
func CustomNew() (*Custom, error) {

	drawingArea, err := gtk.DrawingAreaNew()

	if err != nil {
		return nil, err
	}

	ob0 := Object{Point{150, -150}, 1}
	ob1 := Object{Point{-150, -150}, 1}
	ob2 := Object{Point{-150, 150}, 1}
	ob3 := Object{Point{150, 150}, 1}
	return &Custom{drawingArea, 0.0, []*Object{&ob0, &ob1, &ob2, &ob3}}, nil
}

/** */
func (c *Custom) DrawCustom(custom *gtk.DrawingArea, ctx *cairo.Context) {

	// calculate center and radius
	w := float64(c.GetAllocatedWidth())
	h := float64(c.GetAllocatedHeight())
	r0 := math.Min(w/2, h/2)
	s0 := math.Pi / 2

	// fill background
	ctx.SetSourceRGB(0, 0, 0)
	ctx.Rectangle(0, 0, w, h)
	ctx.Fill()

	ctx.Translate(w/2, h/2)
	ctx.SetSourceRGB(0, 0.5, 0)
	ctx.SetLineWidth(2)
	ctx.MoveTo(-w/2, 0)
	ctx.LineTo(w/2, 0)
	ctx.Stroke()
	ctx.MoveTo(0, -h/2)
	ctx.LineTo(0, h/2)
	ctx.Stroke()

	for r := r0 / 6; r < r0; r += r0 / 6 {
		ctx.Arc(0, 0, r, 0, math.Pi*2)
		ctx.Stroke()
	}

	x1 := r0 * math.Cos(c.sweep)
	y1 := r0 * math.Sin(c.sweep)
	x2 := r0 * math.Cos(c.sweep-s0/2)
	y2 := r0 * math.Sin(c.sweep-s0/2)
	x3 := r0 * math.Cos(c.sweep-s0)
	y3 := r0 * math.Sin(c.sweep-s0)

	// debugging
	// ctx.MoveTo(x1, y1)
	// ctx.LineTo(x2, y2)
	// ctx.Stroke()

	// draw leading edge
	ctx.SetSourceRGB(0, 1, 0)
	ctx.MoveTo(0, 0)
	ctx.LineTo(x1, y1)
	ctx.Stroke()

	// set the fade pattern
	pat, _ := cairo.NewPatternLinear(x1, y1, x2, y2)
	pat.AddColorStopRGBA(0, 0, 1, 0, 0.5)
	pat.AddColorStopRGBA(1, 0, 0, 0, 0.5)
	ctx.SetSource(pat)

	// draw the persistence sector
	ctx.MoveTo(0, 0)
	ctx.Arc(0, 0, r0, c.sweep-s0/2, c.sweep)
	ctx.ClosePath()
	ctx.Fill()

	// move to next position
	c.sweep += math.Pi / 300
	if c.sweep > 2*math.Pi {
		c.sweep = 0
	}

	// debugging
	// ctx.SetSourceRGB(0, 0, 1)
	// ctx.MoveTo(0, 0)
	// ctx.LineTo(x3, y3)
	// ctx.Stroke()

	// paint the objects if they are in the sector
	for _, ob := range c.objects {
		if pointInPolygon(ob.point, []Point{{0, 0}, {x1, y1}, {x2, y2}, {x3, y3}}) {
			ob.b -= 0.005
			ctx.SetSourceRGB(0, ob.b, 0)
			ctx.Arc(ob.point.x, ob.point.y, 5, 0, math.Pi*2)
			ctx.Fill()
		} else {
			ob.b = 1
		}

		ob.point.y += (h / 2400.0)
		if ob.point.y > h/2 {
			ob.point.y = -h / 2
		}
	}
}

func pointInPolygon(pt Point, py []Point) bool {

	x := pt.x
	y := pt.y

	inside := false

	i := 0
	for j := len(py) - 1; i < len(py); i++ {
		xi := py[i].x
		yi := py[i].y
		xj := py[j].x
		yj := py[j].y

		intersect := ((yi > y) != (yj > y)) && (x < (xj-xi)*(y-yi)/(yj-yi)+xi)
		if intersect {

			inside = !inside
		}
		j = i

	}

	return inside
}
