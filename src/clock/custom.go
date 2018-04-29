package main

//#cgo pkg-config: cairo cairo-gobject gtk+-3.0
//#include <stdlib.h>
//#include <cairo.h>
//#include <cairo-gobject.h>
//#include <gdk/gdk.h>
//#include <gtk/gtk.h>
//#include </home/andrew/go-learning/src/github.com/gotk3/gotk3/gdk/gdk.go.h>
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

// Custom is a custom widget
type Custom struct {
	gtk.DrawingArea
	drawFunc func(width, height float64, ctx *cairo.Context)
	width    int
	height   int
	surface  *cairo.Surface
}

// CustomNew creates a new widget
func CustomNew(drawFunc func(width, height float64, ctx *cairo.Context)) (*Custom, error) {

	drawingArea, err := gtk.DrawingAreaNew()

	if err != nil {
		return nil, err
	}

	custom := Custom{*drawingArea, drawFunc, 0, 0, nil}

	drawingArea.Connect("draw", custom.drawEvent)
	drawingArea.Connect("configure-event", custom.configureEvent)

	return &custom, nil
}

// ConfigureEvent :
func (c *Custom) configureEvent(drawingArea *gtk.DrawingArea, evt *gdk.Event) bool {
	fmt.Printf("ConfigureEvent\n")

	if c.surface != nil {

		surfaceNative := c.surface.Native()

		s := (*C.cairo_surface_t)(unsafe.Pointer(surfaceNative))
		C.cairo_surface_destroy(s)
	}

	parent, err := drawingArea.GetParent()
	if err != nil {
		return false
	}

	p, err := parent.GetWindow()
	if err != nil {
		return false
	}

	c.height = drawingArea.GetAllocatedHeight()
	c.width = drawingArea.GetAllocatedWidth()

	//fmt.Printf("w=%d h=%d\n", c.width, c.height)
	c.surface = createSimilarSurface(p, cairo.CONTENT_COLOR, c.width, c.height)

	if err != nil {
		return false
	}

	c.Clear()
	return true
}

// Clear the surfate to whhite
func (c *Custom) Clear() {
	//fmt.Printf("Clear\n")
	ctx := cairo.Create(c.surface)
	ctx.SetSourceRGB(1, 1, 1)
	ctx.Paint()
	c.drawFunc(float64(c.width), float64(c.height), ctx)
	destroyContext(ctx)
}

// Clear the surfate to whhite
func (c *Custom) Draw() {
	ctx := cairo.Create(c.surface)
	ctx.SetSourceRGB(1, 1, 1)

	c.drawFunc(float64(c.width), float64(c.height), ctx)

	c.QueueDrawArea(0, 0, c.width, c.height)
}

// DrawEvent :
func (c *Custom) drawEvent(drawingArea *gtk.DrawingArea, ctx *cairo.Context) bool {
	//	fmt.Printf("drawEvent\n")
	ctx.SetSourceSurface(c.surface, 0, 0)

	ctx.Paint()
	return false
}

// CreateSimilarSurface is a wrapper around gdk_window_create_similar_surface
func createSimilarSurface(window *gdk.Window, content cairo.Content, width, height int) *cairo.Surface {

	gdkWin := C.toGdkWindow(unsafe.Pointer(window.GObject))

	surfaceNative := C.gdk_window_create_similar_surface(gdkWin, C.cairo_content_t(content), C.int(width), C.int(height))

	surface := cairo.NewSurface(uintptr(unsafe.Pointer(surfaceNative)), false)
	return surface
}

func destroyContext(ctx *cairo.Context) {
	ctxNative := ctx.Native()
	cx := (*C.cairo_t)(unsafe.Pointer(ctxNative))
	C.cairo_destroy(cx)
}
