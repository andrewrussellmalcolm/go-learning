package main

import "fmt"

// Drawer :
type Drawer interface {
	draw()
}

// Origin :
type Origin struct {
	x, y float32
}

// Color :
type Color struct {
	r, g, b uint8
}

// Shape :
type Shape struct {
	Origin
	Color
}

// Square :
type Square struct {
	Shape
	l float32
}

// Circle :
type Circle struct {
	Shape
	r float32
}

// Triangle :
type Triangle struct {
	Shape
	r float32
}

func NewOrigin(x, y float32) Origin {
	return Origin{x, y}
}
func NewColor(r, g, b uint8) Color {
	return Color{r, g, b}
}
func NewShape(origin Origin, color Color) Shape {
	return Shape{origin, color}
}
func (s Square) draw() {
	fmt.Println("Square:", s.x, s.y, s.l, s.Color)
}

func (c Circle) draw() {
	fmt.Println("Circle:", c.x, c.y, c.r, c.Color)
}

func (t Triangle) draw() {
	fmt.Println("Triangle:", t.x, t.y, t.r, t.Color)
}

func main() {

	c := Circle{NewShape(NewOrigin(0, 0), NewColor(0xff, 0x00, 0x00)), 3}
	s := Square{NewShape(NewOrigin(1, 1), NewColor(0x00, 0xff, 0x00)), 4}
	t := Triangle{NewShape(NewOrigin(2, 2), NewColor(0x00, 0x00, 0xff)), 4}

	drawers := []Drawer{c, s, t}

	for index, drawer := range drawers {

		drawer.draw()

		fmt.Println(index)
	}
}
