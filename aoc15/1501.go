package aoc15

import (
	. "github.com/roidaradal/aoc-go/aoc"
)

func Day01() Solution {
	line := data01(true)

	// Part 1
	level := elevatorFloor(line, nil)

	// Part 2
	goal := -1
	index := elevatorFloor(line, &goal)

	return NewSolution(level, index)
}

func data01(full bool) string {
	return ReadFirstLine(15, 1, full)
}

func elevatorFloor(line string, goal *int) int {
	level := 0
	for i, x := range line {
		if x == '(' {
			level += 1
		} else if x == ')' {
			level -= 1
		}
		if goal != nil && level == *goal {
			return i + 1
		}
	}
	return level
}
