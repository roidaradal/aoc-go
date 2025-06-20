package aoc20

import (
	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day01() Solution {
	numbers := data01(true)

	// Part 1
	value1 := find2020Combo(numbers, 2)

	// Part 2
	value2 := find2020Combo(numbers, 3)

	return NewSolution(value1, value2)
}

func data01(full bool) []int {
	return fn.Map(ReadLines(20, 1, full), fn.ParseInt)
}

func find2020Combo(numbers []int, take int) int {
	for _, combo := range Combinations(numbers, take) {
		if fn.Sum(combo) == 2020 {
			return fn.Product(combo)
		}
	}
	return 0
}
