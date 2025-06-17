package aoc18

import (
	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/ds"
)

func Day01() Solution {
	numbers := data01(true)

	// Part 1
	total := fn.Sum(numbers)

	// Part 2
	limit := len(numbers)
	done := ds.NewSet[int]()
	i, curr := 0, 0
	for {
		curr += numbers[i]
		if done.Contains(curr) {
			break
		}
		done.Add(curr)
		i = (i + 1) % limit
	}

	return NewSolution(total, curr)
}

func data01(full bool) []int {
	return fn.Map(ReadLines(18, 1, full), fn.ParseInt)
}
