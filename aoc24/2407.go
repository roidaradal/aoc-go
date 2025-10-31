package aoc24

import (
	"fmt"
	"slices"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/conv"
	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
)

func Day07() Solution {
	pairs := day07(true)

	// Part 1
	score1 := scoreFn(false)
	total1 := list.Sum(fn.Map(pairs, score1))

	// Part 2
	score2 := scoreFn(true)
	total2 := list.Sum(fn.Map(pairs, score2))

	return NewSolution(total1, total2)
}

func day07(full bool) []Pair {
	return fn.Map(ReadLines(24, 7, full), func(line string) Pair {
		parts := str.CleanSplit(line, ":")
		return Pair{
			goal:    conv.ParseInt(parts[0]),
			numbers: ToIntList(parts[1], " "),
		}
	})
}

type Pair struct {
	goal    int
	numbers []int
}

func scoreFn(useConcat bool) func(Pair) int {
	return func(pair Pair) int {
		return fn.Ternary(isPossible(&pair, useConcat), pair.goal, 0)
	}
}

func isPossible(pair *Pair, useConcat bool) bool {
	q := []int{pair.numbers[0]}
	for _, y := range pair.numbers[1:] {
		q2 := make([]int, 0)
		for _, x := range q {
			q2 = append(q2, x+y)
			q2 = append(q2, x*y)
			if useConcat {
				digits := conv.ParseInt(fmt.Sprintf("%d%d", x, y))
				q2 = append(q2, digits)
			}
		}
		q = q2
	}
	return slices.Contains(q, pair.goal)
}
