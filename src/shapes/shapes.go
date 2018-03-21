package main

import "fmt"

// Shape :
type Shape interface {
	draw()
}

// Square :
type Square struct {
	x, y, l float32
}

// Circle :
type Circle struct {
	x, y, r float32
}

// Triangle :
type Triangle struct {
	x, y, r float32
}

func (s Square) draw() {
	fmt.Println("Square:", s.x, s.y, s.l)
}

func (c Circle) draw() {
	fmt.Println("Circle:", c.x, c.y, c.r)
}

func (c Triangle) draw() {
	fmt.Println("Triangle:", c.x, c.y, c.r)
}

func main() {

	c := Circle{0, 0, 3}
	s := Square{1, 1, 4}
	t := Triangle{1, 1, 4}

	shapes := []Shape{c, s, t}

	for index, shape := range shapes {

		shape.draw()

		fmt.Println(index)
	}
}
