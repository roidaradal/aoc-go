package aoc16

import (
	"fmt"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day08() Solution {
	grid, commands := data08(true)
	for _, cmd := range commands {
		if cmd.Type == "rect" {
			turnOn(grid, cmd.Int2)
		} else if cmd.Type == "row" {
			rotateRow(grid, cmd.Int2)
		} else if cmd.Type == "col" {
			rotateCol(grid, cmd.Int2)
		}
	}
	// Part 1
	count := GridSum(grid)
	// Part 2
	displayGrid(grid)
	return NewSolution(count, "")
}

type Command struct {
	Type string
	Int2
}

func data08(full bool) (IntGrid, []Command) {
	commands := make([]Command, 0)
	for _, line := range ReadLines(16, 8, full) {
		p := fn.SpaceSplit(line)
		lastIdx := len(p) - 1
		if p[0] == "rect" {
			cmd := Command{
				Type: "rect",
				Int2: ToInt2(p[1], "x"),
			}
			commands = append(commands, cmd)
		} else if p[1] == "column" {
			col := fn.ParseInt(fn.CleanSplit(p[2], "=")[1])
			rot := fn.ParseInt(p[lastIdx])
			cmd := Command{
				Type: "col",
				Int2: Int2{col, rot},
			}
			commands = append(commands, cmd)
		} else if p[1] == "row" {
			row := fn.ParseInt(fn.CleanSplit(p[2], "=")[1])
			rot := fn.ParseInt(p[lastIdx])
			cmd := Command{
				Type: "row",
				Int2: Int2{row, rot},
			}
			commands = append(commands, cmd)
		}
	}
	grid := NewIntGrid(6, 50, 0)
	return grid, commands
}

func turnOn(grid IntGrid, p Int2) {
	cols, rows := p[0], p[1]
	for row := range rows {
		for col := range cols {
			grid[row][col] = 1
		}
	}
}

func rotateRow(grid IntGrid, p Int2) {
	rowIdx, steps := p[0], p[1]
	row := grid[rowIdx]
	last := len(row) - 1
	for range steps {
		row = JoinLists(row[last:], row[:last])
	}
	grid[rowIdx] = row
}

func rotateCol(grid IntGrid, p Int2) {
	numRows := len(grid)
	colIdx, steps := p[0], p[1]
	col := make([]int, numRows)
	for i, line := range grid {
		col[i] = line[colIdx]
	}
	last := len(col) - 1
	for range steps {
		col = JoinLists(col[last:], col[:last])
	}
	for row := range numRows {
		grid[row][colIdx] = col[row]
	}
}

func displayGrid(grid IntGrid) {
	for _, line := range grid {
		out := make([]byte, len(line))
		for i, x := range line {
			if x == 0 {
				out[i] = ' '
			} else {
				out[i] = '#'
			}
		}
		fmt.Println(string(out))
	}
}
