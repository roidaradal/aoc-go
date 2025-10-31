package aoc21

import (
	"math"
	"slices"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/conv"
	"github.com/roidaradal/fn/list"
)

func Day07() Solution {
	numbers := data07(true)
	slices.Sort(numbers)

	// Part 1
	mid := len(numbers) / 2
	median := numbers[mid]
	total := list.Sum(fn.Map(numbers, func(x int) int {
		return conv.Abs(median - x)
	}))

	// Part 2
	start, end := numbers[0], Last(numbers, 1)
	minCost := math.MaxInt
	for target := range NumRange(start, end+1) {
		cost := list.Sum(fn.Map(numbers, func(x int) int {
			return list.Sum(NumRange(0, conv.Abs(target-x)+1))
		}))
		minCost = min(minCost, cost)
	}

	return NewSolution(total, minCost)
}

func data07(full bool) []int {
	line := ReadFirstLine(21, 7, full)
	return ToIntList(line, ",")
}
