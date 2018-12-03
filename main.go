package main

import (
	"errors"
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

// Game is the main object on which 2048 is played.
type Game struct {
	Grid  GameGrid
	Score int
}

// Point is a simple array of ints with X and Y values.
type Point [2]int

// GameGrid is the grid of integers 2048 is played on.
type GameGrid [4]GridRow

// GridRow is a row on the GameGrid.
type GridRow [4]int

// Merge performs the usual 2048 merge: combine similar numbers, then push all numbers to the left
func (row *GridRow) Merge() (score int) {
	newrow := GridRow{0, 0, 0, 0}
	i := 0
	for _, cell := range row {
		switch {
		case cell == 0: // skip it
		case newrow[i] == 0:
			newrow[i] = cell
		case cell == newrow[i]:
			newrow[i] += cell
			score += newrow[i]
			i++
		case cell != newrow[i]:
			i++
			newrow[i] += cell
		}
	}
	*row = newrow
	return score
}

// AddNumber adds a number to an empty spot on the grid.
func (grid *GameGrid) AddNumber() error {
	// find all the empty spots on the grid
	var emptySpots []Point
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if grid[x][y] == 0 {
				emptySpots = append(emptySpots, Point{x, y})
			}
		}
	}
	if len(emptySpots) == 0 {
		return errors.New("no empty spaces")
	}
	// get a random spot on the grid
	spot := emptySpots[random.Int()%len(emptySpots)]
	// determine whether to add a 2 or a 4
	number := validNewValues[random.Int()%len(validNewValues)]
	// set the spot accordingly
	grid[spot[0]][spot[1]] = number
	return nil
}

// RotateCW is a lazy clockwise rotation algorithm
func (grid *GameGrid) rotateCW() {
	*grid = GameGrid{
		GridRow{grid[3][0], grid[2][0], grid[1][0], grid[0][0]},
		GridRow{grid[3][1], grid[2][1], grid[1][1], grid[0][1]},
		GridRow{grid[3][2], grid[2][2], grid[1][2], grid[0][2]},
		GridRow{grid[3][3], grid[2][3], grid[1][3], grid[0][3]},
	}
}

// RotateCCW is a lazy counterclockwise rotation algorithm
func (grid *GameGrid) rotateCCW() {
	*grid = GameGrid{
		GridRow{grid[0][3], grid[1][3], grid[2][3], grid[3][3]},
		GridRow{grid[0][2], grid[1][2], grid[2][2], grid[3][2]},
		GridRow{grid[0][1], grid[1][1], grid[2][1], grid[3][1]},
		GridRow{grid[0][0], grid[1][0], grid[2][0], grid[3][0]},
	}
}

// Rotate180 is a lazy grid flipping algorithm
func (grid *GameGrid) rotate180() {
	*grid = GameGrid{
		GridRow{grid[3][3], grid[3][2], grid[3][1], grid[3][0]},
		GridRow{grid[2][3], grid[2][2], grid[2][1], grid[2][0]},
		GridRow{grid[1][3], grid[1][2], grid[1][1], grid[1][0]},
		GridRow{grid[0][3], grid[0][2], grid[0][1], grid[0][0]},
	}
}

func (grid *GameGrid) merge() (score int) {
	for i := 0; i < 4; i++ {
		score += grid[i].Merge()
	}
	return score
}

// MergeLeft merges all rows in the grid to the left.
func (grid *GameGrid) MergeLeft() (score int) {
	return grid.merge()
}

// MergeDown merges all rows in the grid downward.
func (grid *GameGrid) MergeDown() (score int) {
	grid.rotateCW()
	score = grid.merge()
	grid.rotateCCW()
	return score
}

// MergeUp merges all rows in the grid upward.
func (grid *GameGrid) MergeUp() (score int) {
	grid.rotateCCW()
	score = grid.merge()
	grid.rotateCW()
	return score
}

// MergeRight merges all rows in the grid to the right.
func (grid *GameGrid) MergeRight() (score int) {
	grid.rotate180()
	score = grid.merge()
	grid.rotate180()
	return score
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

var (
	random         *rand.Rand
	validNewValues = []int{2, 2, 2, 4}
)

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err) // excuse me but what
	}
	defer termbox.Close()
	game := Game{
		Grid: GameGrid{},
	}
	game.Grid.AddNumber()
	game.Grid.AddNumber()
	game.Draw()
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
		err := game.Grid.AddNumber()
		if err != nil {
			termprint(Point{0, 25}, "Game Over", termbox.ColorDefault|termbox.AttrBold, termbox.ColorDefault)
		}
		game.Draw()
		termbox.Flush()
	}
}
