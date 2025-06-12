package aoc17

import (
	"fmt"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day05() {
	jumps := data05(true)
	increment1 := func(jump int) int {
		return 1
	}
	count := countJumps(jumps, increment1)
	fmt.Println(count)

	jumps = data05(true)
	increment2 := func(jump int) int {
		if jump >= 3 {
			return -1
		}
		return 1
	}
	count = countJumps(jumps, increment2)
	fmt.Println(count)
}

func data05(full bool) []int {
	return fn.Map(ReadLines(full), fn.ParseInt)
}

func countJumps(jumps []int, increment func(int) int) int {
	limit := len(jumps)
	i, count := 0, 0
	for 0 <= i && i < limit {
		jump := jumps[i]
		jumps[i] += increment(jump)
		i += jump
		count += 1
	}
	return count
}
