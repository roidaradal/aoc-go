package aoc

import "github.com/roidaradal/fn"

var (
	U Delta = [2]int{-1, 0}
	D Delta = [2]int{1, 0}
	L Delta = [2]int{0, -1}
	R Delta = [2]int{0, 1}
)

var (
	LeftOf  = map[Delta]Delta{U: L, L: D, D: R, R: U}
	RightOf = map[Delta]Delta{U: R, R: D, D: L, L: U}
)

func NewIntGrid(rows, cols, initial int) IntGrid {
	grid := make(IntGrid, rows)
	for row := range rows {
		line := make([]int, cols)
		for col := range cols {
			line[col] = initial
		}
		grid[row] = line
	}
	return grid
}

func Move(c Coords, d Delta) Coords {
	row, col := c[0], c[1]
	dy, dx := d[0], d[1]
	return Coords{row + dy, col + dx}
}

func GridSum(grid IntGrid) int {
	total := 0
	for _, line := range grid {
		for _, x := range line {
			total += x
		}
	}
	return total
}

func Manhattan(c1 Coords, c2 Coords) int {
	y1, x1 := c1[0], c1[1]
	y2, x2 := c2[0], c2[1]
	return fn.Abs(y2-y1) + fn.Abs(x2-x1)
}

func ManhattanOrigin(c Coords) int {
	return Manhattan(c, Coords{0, 0})
}

func InsideBounds(c Coords, maxBounds Dims2) bool {
	return InsideBoundsFull(c, maxBounds, Dims2{0, 0})
}

func InsideBoundsFull(c Coords, maxBounds Dims2, minBounds Dims2) bool {
	row, col := c[0], c[1]
	minRows, minCols := minBounds[0], minBounds[1]
	numRows, numCols := maxBounds[0], maxBounds[1]
	return minRows <= row && row < numRows && minCols <= col && col < numCols
}
