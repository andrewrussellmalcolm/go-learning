package main

import "C"
import (
	"fmt"
	"math"
	"time"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
)

func main() {

	gtk.Init(nil)

	windowMain, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	bailOnError(err)

	windowMain.SetTitle("Clock")
	windowMain.SetDefaultSize(600, 600)
	windowMain.Connect("destroy", func() {
		gtk.MainQuit()
	})

	custom, err := CustomNew(DrawClock)

	windowMain.Container.Add(custom)

	windowMain.ShowAll()

	ticker := time.NewTicker(time.Second)
	custom.Draw()
	custom.QueueDraw()

	go func() {

		for range ticker.C {
			custom.Draw()

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

// DrawClock draws th c4 clock
func DrawClock(w, h float64, ctx *cairo.Context) {
	// calculate center and radius
	x0 := w / 2
	y0 := h / 2
	r := 2 * math.Min(x0, y0) / 3

	hours := float64(time.Now().Hour())
	mins := float64(time.Now().Minute())
	secs := float64(time.Now().Second())

	// fil background
	ctx.SetSourceRGB(0, 0, 0)
	ctx.Rectangle(0, 0, w, h)
	ctx.Fill()

	// draw trhe cardinals
	ctx.SetSourceRGB(1.0, 1.0, 1.0)
	ctx.SetLineWidth(10)
	for theta0 := 0.0; theta0 < 2*math.Pi; theta0 += math.Pi / 6 {
		x1 := x0 + r*math.Sin(theta0)*0.9
		y1 := y0 + r*math.Cos(theta0)*0.9
		x2 := x0 + r*math.Sin(theta0)*1.1
		y2 := y0 + r*math.Cos(theta0)*1.1
		ctx.MoveTo(x1, y1)
		ctx.LineTo(x2, y2)
		ctx.Stroke()
	}

	// draw hour hand
	ctx.SetSourceRGB(0, 0, 1.0)
	ctx.SetLineWidth(16)
	theta1 := -math.Pi * 2.0 * (hours + (mins / 60.0)) / 12.0
	x3 := x0 - r*math.Sin(theta1)*0.75
	y3 := y0 - r*math.Cos(theta1)*0.75
	x4 := x0 + r*math.Sin(theta1)*0.2
	y4 := y0 + r*math.Cos(theta1)*0.2
	ctx.MoveTo(x4, y4)
	ctx.LineTo(x3, y3)
	ctx.Stroke()

	// draw minute hand
	ctx.SetSourceRGB(1.0, 0.0, 0.0)
	ctx.SetLineWidth(10)
	theta2 := -math.Pi * 2.0 * (mins + secs/60.0) / 60.0
	x5 := x0 - r*math.Sin(theta2)*1.1
	y5 := y0 - r*math.Cos(theta2)*1.1
	x6 := x0 + r*math.Sin(theta2)*0.2
	y6 := y0 + r*math.Cos(theta2)*0.2
	ctx.MoveTo(x6, y6)
	ctx.LineTo(x5, y5)
	ctx.Stroke()

	// draw second hand
	ctx.SetSourceRGB(1.0, 1.0, 0.0)
	ctx.SetLineWidth(1)
	theta3 := -math.Pi * 2.0 * secs / 60.0
	x7 := x0 - r*math.Sin(theta3)*1.1
	y7 := y0 - r*math.Cos(theta3)*1.1
	x8 := x0 + r*math.Sin(theta3)*0.2
	y8 := y0 + r*math.Cos(theta3)*0.2
	ctx.MoveTo(x8, y8)
	ctx.LineTo(x7, y7)
	ctx.Stroke()

	// draw central boss
	//ctx.Scale(1, 0.7)
	ctx.Arc(x0, y0, r/15, 0, math.Pi*2)
	ctx.Fill()
}

//go build -v -tags gtk_3_18 -gcflags "-N -l"
