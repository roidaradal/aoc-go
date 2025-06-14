package aoc18

import (
	"fmt"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day01() {
	numbers := data01(true)
	total := fn.Sum(numbers)
	fmt.Println(total)

	limit := len(numbers)
	done := make(map[int]bool)
	i, curr := 0, 0
	for {
		curr += numbers[i]
		if done[curr] {
			break
		}
		done[curr] = true
		i = (i + 1) % limit
	}
	fmt.Println(curr)
}

func data01(full bool) []int {
	return fn.Map(ReadLines(full), fn.ParseInt)
}
