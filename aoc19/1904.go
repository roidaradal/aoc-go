package aoc19

import (
	"slices"
	"strconv"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day04() Solution {
	start, end := data04(true)
	count1, count2 := 0, 0
	for x := start; x < end; x++ {
		// Part 1
		if isValid(x) {
			count1 += 1
		}

		// Part 2
		if isValid2(x) {
			count2 += 1
		}
	}

	return NewSolution(count1, count2)
}

func data04(full bool) (int, int) {
	pair := ToIntList(ReadFirstLine(19, 4, full), "-")
	return pair[0], pair[1]
}

func isValid(number int) bool {
	x := strconv.Itoa(number)
	if x != SortedString(x) {
		return false
	}

	return HasTwins(x, 0)
}

func isValid2(number int) bool {
	x := strconv.Itoa(number)
	if x != SortedString(x) {
		return false
	}

	sizes := fn.Map(GroupChunks(x), func(chunk string) int {
		return len(chunk)
	})
	return slices.Contains(sizes, 2)
}
