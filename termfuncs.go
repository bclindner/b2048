package main

import (
	"github.com/nsf/termbox-go"
)

func termprint(starting Point, msg string, fg, bg termbox.Attribute) {
	for i, char := range msg {
		termbox.SetCell(starting[0]+i, starting[1], char, fg|termbox.AttrBold, bg)
	}
}

func blitbox(starting, ending Point, fg, bg termbox.Attribute) {
	for x := starting[0]; x < ending[0]; x++ {
		for y := starting[1]; y < ending[1]; y++ {
			termbox.SetCell(x, y, ' ', fg, bg)
		}
	}
}
