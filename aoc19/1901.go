package aoc19

import (
	"math"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day01() Solution {
	numbers := data01(true)

	// Part 1
	total1 := fn.Sum(fn.Map(numbers, fuel))

	// Part 2
	total2 := fn.Sum(fn.Map(numbers, totalFuel))

	return NewSolution(total1, total2)
}

func data01(full bool) []int {
	return fn.Map(ReadLines(19, 1, full), fn.ParseInt)
}

func fuel(x int) int {
	f := math.Floor(float64(x)/3) - 2
	f = max(f, 0)
	return int(f)
}

func totalFuel(x int) int {
	total := 0
	for x > 0 {
		x = fuel(x)
		total += x
	}
	return total
}
