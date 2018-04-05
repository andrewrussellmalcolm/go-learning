package main

import (
	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

func main() {
	termbox.Init()
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	//const coldef = termbox.ColorDefault
	//termbox.Clear(coldef, coldef)
	//w, h := termbox.Size()

	tbprint(0, 0, termbox.ColorGreen, termbox.ColorBlack, "Hello")
	termbox.Flush()
	termbox.PollEvent()
	termbox.PollEvent()
	termbox.PollEvent()
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func tbfill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}
