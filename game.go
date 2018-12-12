package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

var (
	validNewValues = []int{2, 2, 2, 4}
)

// Game is the main object on which 2048 is played.
type Game struct {
	Grid  GameGrid
	Score int
}

// NewGame creates a new game
func NewGame() (g Game) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	// zero out score
	g.Score = 0
	// zero out the grid
	g.Grid = GameGrid{}
	// add two numbers
	g.Grid.AddNumber()
	g.Grid.AddNumber()
	g.Draw()
	return g
}

// Draw draws the grid into Termbox.
func (g *Game) Draw() {
	for y, row := range g.Grid {
		for x, cell := range row {
			fg := termbox.ColorDefault
			bg := termbox.ColorDefault
			// determine the color of the cell
			switch cell {
			case 2:
				bg = termbox.ColorRed
			case 4:
				bg = termbox.ColorGreen
			case 8:
				bg = termbox.ColorBlue
			case 16:
				bg = termbox.ColorYellow
			case 32:
				bg = termbox.ColorMagenta
			case 64:
				bg = termbox.ColorCyan
			case 128:
				bg = termbox.ColorRed
			case 256:
				bg = termbox.ColorGreen
			case 512:
				bg = termbox.ColorBlue
			case 1024:
				bg = termbox.ColorYellow
			case 2048:
				bg = termbox.ColorMagenta
			}
			blitbox(Point{x*5 + x, y*5 + y}, Point{x*5 + 5 + x, y*5 + 5 + y}, fg, bg)
			termprint(Point{x*5 + x, y*5 + 2 + y}, fmt.Sprintf("%d", cell), fg, bg)
		}
	}
	termprint(Point{0, 24}, fmt.Sprintf("Score: %d", g.Score), termbox.ColorDefault|termbox.AttrBold, termbox.ColorDefault)
	termbox.Flush()
}

// Tick advances the game, adding a number and drawing the screen.
func (g *Game) Tick(grid GameGrid, score int) {
	if grid != g.Grid {
		g.Grid = grid
		g.Score += score
		err := g.Grid.AddNumber()
		if err != nil {
			if g.IsOver() {
				termprint(Point{0, 25}, "Game Over", termbox.ColorDefault|termbox.AttrBold, termbox.ColorDefault)
				termprint(Point{0, 26}, "Press R to Restart", termbox.ColorDefault|termbox.AttrBold, termbox.ColorDefault)
			}
		}
	}
	emptyspaces := g.Grid.GetEmptySpaces()
	if len(emptyspaces) == 0 && g.IsOver() {
		termprint(Point{0, 25}, "Game Over", termbox.ColorDefault|termbox.AttrBold, termbox.ColorDefault)
		termprint(Point{0, 26}, "Press R to Restart", termbox.ColorDefault|termbox.AttrBold, termbox.ColorDefault)
	}
	g.Draw()
}

// IsOver checks if the game is in a failure state (read: the player has no more usable moves).
func (g *Game) IsOver() bool {
	// check for failure state
	up, _ := g.Grid.MergeUp()
	if g.Grid != up {
		return false
	}
	down, _ := g.Grid.MergeDown()
	if g.Grid != down {
		return false
	}
	left, _ := g.Grid.MergeLeft()
	if g.Grid != left {
		return false
	}
	right, _ := g.Grid.MergeRight()
	if g.Grid != right {
		return false
	}
	return true
}
