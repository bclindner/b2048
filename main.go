package main

import (
	"github.com/nsf/termbox-go"
)

// Point is a simple array of ints with X and Y values.
type Point [2]int

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err) // excuse me but what
	}
	defer termbox.Close()
	game := NewGame()
Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				game.Score += game.Grid.MergeUp()
			case termbox.KeyArrowDown:
				game.Score += game.Grid.MergeDown()
			case termbox.KeyArrowLeft:
				game.Score += game.Grid.MergeLeft()
			case termbox.KeyArrowRight:
				game.Score += game.Grid.MergeRight()
			case termbox.KeyCtrlC:
				break Loop
			}
		}
		game.Tick()
	}
}
