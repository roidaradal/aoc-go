package aoc15

import (
	"fmt"

	. "github.com/roidaradal/aoc-go/aoc"
)

func Day01() {
	line := data01(true)
	// Part 1
	level := elevatorFloor(line, nil)
	fmt.Println(level)
	// Part 2
	goal := -1
	level = elevatorFloor(line, &goal)
	fmt.Println(level)
}

func data01(full bool) string {
	return ReadLines(full)[0]
}

func elevatorFloor(line string, goal *int) int {
	T := map[rune]int{
		'(': 1,
		')': -1,
	}
	level := 0
	for i, x := range line {
		level += T[x]
		if goal != nil && level == *goal {
			return i + 1
		}
	}
	return level
}
