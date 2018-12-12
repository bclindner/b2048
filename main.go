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
		moved := false
		grid := GameGrid{}
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				grid, score = game.Grid.MergeUp()
				moved = true
			case termbox.KeyArrowDown:
				grid, score = game.Grid.MergeDown()
				moved = true
			case termbox.KeyArrowLeft:
				grid, score = game.Grid.MergeLeft()
				moved = true
			case termbox.KeyArrowRight:
				grid, score = game.Grid.MergeRight()
				moved = true
			case termbox.KeyCtrlC:
				break Loop
			}
			if !moved {
				switch ev.Ch {
				// r to restart (simply resets the game)
				case 'r':
					game = NewGame()
				// q to quit
				case 'q':
					break Loop
				}
			}
		}
		if moved {
			game.Tick(grid, score)
		}
	}
}
