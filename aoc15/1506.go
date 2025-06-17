package aoc15

import (
	"strings"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day06() Solution {
	side := 1000

	// Part 1
	mask := map[string]int{
		on:     1,
		off:    0,
		toggle: -1,
	}
	commands := data06(mask, true)
	grid := NewIntGrid(side, side, 0)
	for _, cmd := range commands {
		b, y1, x1, y2, x2 := cmd[0], cmd[1], cmd[2], cmd[3], cmd[4]
		for y := y1; y < y2; y++ {
			for x := x1; x < x2; x++ {
				if b == -1 {
					grid[y][x] ^= 1 // flip bit
				} else {
					grid[y][x] = b
				}
			}
		}
	}
	part1 := GridSum(grid)

	// Part 2
	mask = map[string]int{
		on:     1,
		off:    -1,
		toggle: 2,
	}
	commands = data06(mask, true)
	grid = NewIntGrid(side, side, 0)
	for _, cmd := range commands {
		b, y1, x1, y2, x2 := cmd[0], cmd[1], cmd[2], cmd[3], cmd[4]
		for y := y1; y < y2; y++ {
			for x := x1; x < x2; x++ {
				value := grid[y][x] + b
				grid[y][x] = max(value, 0)
			}
		}
	}
	part2 := GridSum(grid)

	return NewSolution(part1, part2)
}

var (
	on     string = "turn on"
	off    string = "turn off"
	toggle string = "toggle"
)

type command = [5]int

func data06(mask map[string]int, full bool) []command {
	return fn.Map(ReadLines(15, 6, full), func(line string) command {
		b := 0
		if strings.HasPrefix(line, on) {
			b = mask[on]
		} else if strings.HasPrefix(line, off) {
			b = mask[off]
		} else if strings.HasPrefix(line, toggle) {
			b = mask[toggle]
		}
		p := fn.CleanSplit(line, "through")
		head := Last(fn.SpaceSplit(p[0]))
		tail := p[1]
		c1 := ToIntList(head, ",")
		c2 := ToIntList(tail, ",")
		x1, y1 := c1[0], c1[1]
		x2, y2 := c2[0], c2[1]
		return command{b, y1, x1, y2 + 1, x2 + 1}
	})
}
