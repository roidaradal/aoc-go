package aoc23

import (
	"strconv"
	"strings"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/conv"
	"github.com/roidaradal/fn/str"
)

func Day06() Solution {
	times, bests := data06(true)

	// Part 1
	total1 := 1
	for i, limit := range times {
		breakers := fn.Filter(computeOutcomes(limit), func(d int) bool {
			return d > bests[i]
		})
		total1 *= len(breakers)
	}

	// Part 2
	newLimit := strings.Join(fn.Map(times, strconv.Itoa), "")
	newBest := strings.Join(fn.Map(bests, strconv.Itoa), "")
	limit, best := conv.ParseInt(newLimit), conv.ParseInt(newBest)
	breakers := fn.Filter(computeOutcomes(limit), func(d int) bool {
		return d > best
	})
	total2 := len(breakers)

	return NewSolution(total1, total2)
}

func data06(full bool) ([]int, []int) {
	lines := ReadLines(23, 6, full)
	times := fn.Map(str.SpaceSplit(lines[0])[1:], conv.ParseInt)
	bests := fn.Map(str.SpaceSplit(lines[1])[1:], conv.ParseInt)
	return times, bests
}

func computeOutcomes(limit int) []int {
	return fn.Map(NumRange(0, limit+1), func(hold int) int {
		return hold * (limit - hold)
	})
}
