package aoc22

import (
	"cmp"
	"slices"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day01() Solution {
	calories := data01(true)

	// Part 1
	maxCal := slices.Max(calories)

	// Part 2
	slices.SortFunc(calories, func(a, b int) int {
		return cmp.Compare(b, a)
	})
	top3 := fn.Sum(calories[0:3])

	return NewSolution(maxCal, top3)
}

func data01(full bool) []int {
	calories := make([]int, 0)
	curr := 0
	for _, line := range ReadRawLines(22, 1, full, true) {
		if line == "" {
			calories = append(calories, curr)
			curr = 0
		} else {
			curr += fn.ParseInt(line)
		}
	}
	calories = append(calories, curr)
	return calories
}
