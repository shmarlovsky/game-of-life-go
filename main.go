package main

import (
	"fmt"
	"math/rand"
	"time"
)

type State uint

const (
	Dead  State = 0
	Alive State = 1
)

func (s State) String() string {
	if s == Alive {
		return "0"
	} else if s == Dead {
		return "."
	}
	panic("Unknown cell state")
}

const (
	H = 10
	W = 10
)

type Grid [H][W]State

func initialGridManual() Grid {
	var grid Grid = Grid{
		[W]State{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[W]State{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
		[W]State{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		[W]State{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[W]State{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[W]State{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[W]State{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[W]State{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[W]State{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[W]State{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	return grid
}

func initialGridRandom() Grid {
	var grid Grid
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			grid[h][w] = State(rand.Intn(2))
		}
	}
	return grid
}

func drawLine(element string) {
	fmt.Print(" ")
	for i := 0; i < H; i++ {
		fmt.Print(element)
	}
	fmt.Print("\n")
}

func drawGrid(grid *Grid) {
	drawLine("___")
	for i := 0; i < H; i++ {
		fmt.Printf("|")
		for j := 0; j < W; j++ {
			fmt.Printf(" %s ", grid[i][j])
		}
		fmt.Print("|\n")
	}
	drawLine("\u2594\u2594\u2594")
}

func nextGeneration(grid *Grid) Grid {
	var nextGen Grid

	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {

			// count alive neighbours
			aliveNeighbours := 0
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					if (h+i >= 0 && h+i < H) && (w+j >= 0 && w+j < W) {
						aliveNeighbours += int(grid[h+i][w+j])
					}
				}
			}

			aliveNeighbours -= int(grid[h][w])

			// implement rules

			// Any live cell with fewer than two live neighbors dies as if caused by underpopulation.
			if grid[h][w] == Alive && aliveNeighbours < 2 {
				nextGen[h][w] = Dead
			} else

			// Any live cell with more than three live neighbors dies, as if by overpopulation.
			if grid[h][w] == Alive && aliveNeighbours > 3 {
				nextGen[h][w] = Dead
			} else

			// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
			if grid[h][w] == Dead && aliveNeighbours == 3 {
				nextGen[h][w] = Alive
			} else {
				// Any live cell with two or three live neighbors lives on to the next generation.
				nextGen[h][w] = grid[h][w]
			}

		}
	}

	return nextGen
}

func main() {
	// grid := initialGridManual()
	grid := initialGridRandom()
	drawGrid(&grid)

	for {
		grid = nextGeneration(&grid)
		drawGrid(&grid)
		time.Sleep(300 * time.Millisecond)
		// break
	}

}
