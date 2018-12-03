package main

import (
	"errors"
	"math/rand"
	"time"
)

var (
	random *rand.Rand
)

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

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

// GetEmptySpaces returns all of the empty spaces in the GameGrid.
func (grid *GameGrid) GetEmptySpaces() (emptySpots []Point) {
	// find all the empty spots on the grid
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if grid[x][y] == 0 {
				emptySpots = append(emptySpots, Point{x, y})
			}
		}
	}
	return emptySpots
}

// AddNumber adds a number to an empty spot on the grid.
func (grid *GameGrid) AddNumber() error {
	emptySpots := grid.GetEmptySpaces()
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
func (grid *GameGrid) rotateCW() GameGrid {
	return GameGrid{
		GridRow{grid[3][0], grid[2][0], grid[1][0], grid[0][0]},
		GridRow{grid[3][1], grid[2][1], grid[1][1], grid[0][1]},
		GridRow{grid[3][2], grid[2][2], grid[1][2], grid[0][2]},
		GridRow{grid[3][3], grid[2][3], grid[1][3], grid[0][3]},
	}
}

// RotateCCW is a lazy counterclockwise rotation algorithm
func (grid *GameGrid) rotateCCW() GameGrid {
	return GameGrid{
		GridRow{grid[0][3], grid[1][3], grid[2][3], grid[3][3]},
		GridRow{grid[0][2], grid[1][2], grid[2][2], grid[3][2]},
		GridRow{grid[0][1], grid[1][1], grid[2][1], grid[3][1]},
		GridRow{grid[0][0], grid[1][0], grid[2][0], grid[3][0]},
	}
}

// Rotate180 is a lazy grid flipping algorithm
func (grid *GameGrid) rotate180() GameGrid {
	return GameGrid{
		GridRow{grid[3][3], grid[3][2], grid[3][1], grid[3][0]},
		GridRow{grid[2][3], grid[2][2], grid[2][1], grid[2][0]},
		GridRow{grid[1][3], grid[1][2], grid[1][1], grid[1][0]},
		GridRow{grid[0][3], grid[0][2], grid[0][1], grid[0][0]},
	}
}

func (grid *GameGrid) merge() (newgrid GameGrid, score int) {
	newgrid = *grid
	for i := 0; i < 4; i++ {
		score += newgrid[i].Merge()
	}
	return newgrid, score
}

// MergeLeft merges all rows in the grid to the left.
func (grid *GameGrid) MergeLeft() (newgrid GameGrid, score int) {
	return grid.merge()
}

// MergeDown merges all rows in the grid downward.
func (grid *GameGrid) MergeDown() (newgrid GameGrid, score int) {
	newgrid = grid.rotateCW()
	newgrid, score = newgrid.merge()
	newgrid = newgrid.rotateCCW()
	return newgrid, score
}

// MergeUp merges all rows in the grid upward.
func (grid *GameGrid) MergeUp() (newgrid GameGrid, score int) {
	newgrid = grid.rotateCCW()
	newgrid, score = newgrid.merge()
	newgrid = newgrid.rotateCW()
	return newgrid, score
}

// MergeRight merges all rows in the grid to the right.
func (grid *GameGrid) MergeRight() (newgrid GameGrid, score int) {
	newgrid = grid.rotate180()
	newgrid, score = newgrid.merge()
	newgrid = newgrid.rotate180()
	return newgrid, score
}
