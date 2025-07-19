package aoc23

import (
	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/ds"
)

func Day09() Solution {
	numberLists := data09(true)

	total1, total2 := 0, 0
	for _, numbers := range numberLists {
		// Part 1
		total1 += getNext(numbers)

		// Part 2
		total2 += getPrev(numbers)
	}

	return NewSolution(total1, total2)
}

func data09(full bool) [][]int {
	return fn.Map(ReadLines(23, 9, full), func(line string) []int {
		return ToIntList(line, " ")
	})
}

func getNext(numbers []int) int {
	gap := 0
	diff := numbers
	for !fn.AllEqual(diff, 0) {
		diff = fn.Map(NumRange(1, len(diff)), func(i int) int {
			return diff[i] - diff[i-1]
		})
		gap += Last(diff, 1)
	}
	return Last(numbers, 1) + gap
}

func getPrev(numbers []int) int {
	diff := numbers
	fronts := []int{diff[0]}
	for !fn.AllEqual(diff, 0) {
		diff = fn.Map(NumRange(1, len(diff)), func(i int) int {
			return diff[i] - diff[i-1]
		})
		fronts = append(fronts, diff[0])
	}
	stack := ds.StackFrom(fronts)
	for stack.Len() >= 2 {
		b := stack.Pop()
		a := stack.Pop()
		stack.Push(a - b)
	}
	return stack.Top()
}
