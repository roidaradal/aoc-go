package aoc20

import (
	"slices"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/conv"
	"github.com/roidaradal/fn/list"
)

func Day10() Solution {
	numbers := data10(true)
	slices.Sort(numbers)

	// Part 1
	diffs := make(map[int]int)
	for i, curr := range numbers {
		if i == 0 {
			continue
		}
		d := curr - numbers[i-1]
		diffs[d] += 1
	}
	part1 := diffs[1] * diffs[3]

	// Part 2
	count := make(map[int]int)
	count[Last(numbers, 1)] = 1
	for i := len(numbers) - 2; i >= 0; i -= 1 {
		curr := numbers[i]
		valid := fn.Filter(numbers[i+1:i+4], func(x int) bool {
			return x-curr <= 3
		})
		count[curr] = list.Sum(fn.Map(valid, func(x int) int {
			return count[x]
		}))
	}
	part2 := count[numbers[0]]

	return NewSolution(part1, part2)
}

func data10(full bool) []int {
	numbers := fn.Map(ReadLines(20, 10, full), conv.ParseInt)
	numbers = append(numbers, 0)
	numbers = append(numbers, slices.Max(numbers)+3)
	return numbers
}
