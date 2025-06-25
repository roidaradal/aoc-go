package aoc21

import (
	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day02() Solution {
	moves := data02(true)

	// Part 1
	curr := Coords{0, 0}
	for _, d := range moves {
		curr = Move(curr, d)
	}
	y, x := curr[0], curr[1]
	part1 := y * x

	// Part 2
	y, x, a := 0, 0, 0
	for _, d := range moves {
		dy, dx := d[0], d[1]
		if dy == 0 {
			x += dx
			y += a * dx
		} else {
			a += dy
		}
	}
	part2 := y * x

	return NewSolution(part1, part2)
}

func data02(full bool) []Delta {
	return fn.Map(ReadLines(21, 2, full), func(line string) Delta {
		p := fn.SpaceSplit(line)
		cmd, x := p[0], fn.ParseInt(p[1])
		switch cmd {
		case "forward":
			return Delta{0, x}
		case "up":
			return Delta{-x, 0}
		case "down":
			return Delta{x, 0}
		default:
			return X
		}
	})
}
