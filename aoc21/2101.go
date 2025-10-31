package aoc21

import (
	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/conv"
	"github.com/roidaradal/fn/list"
)

func Day01() Solution {
	numbers := data01(true)
	limit := len(numbers)
	count1, count2 := 0, 0
	for i := range limit {
		// Part 1
		if i < limit-1 && numbers[i+1] > numbers[i] {
			count1 += 1
		}

		// Part 2
		if i < 3 {
			continue
		}
		curr := list.Sum(numbers[i-2 : i+1])
		prev := list.Sum(numbers[i-3 : i])
		if curr > prev {
			count2 += 1
		}
	}
	return NewSolution(count1, count2)
}

func data01(full bool) []int {
	return fn.Map(ReadLines(21, 1, full), conv.ParseInt)
}
