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
		score := 0
		grid := GameGrid{}
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				grid, score = game.Grid.MergeUp()
			case termbox.KeyArrowDown:
				grid, score = game.Grid.MergeDown()
			case termbox.KeyArrowLeft:
				grid, score = game.Grid.MergeLeft()
			case termbox.KeyArrowRight:
				grid, score = game.Grid.MergeRight()
			case termbox.KeyCtrlC:
				break Loop
			}
		}
		game.Tick(grid, score)
	}
}
