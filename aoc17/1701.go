package aoc17

import (
	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day01() Solution {
	word := data01(true)

	// Part 1
	limit := len(word)
	total1 := 0
	for i := range limit {
		j := (i + 1) % limit
		if word[i] == word[j] {
			total1 += fn.ParseInt(string(word[i]))
		}
	}

	// Part 2
	mid := limit / 2
	total2 := 0
	for i := range mid {
		j := mid + i
		if word[i] == word[j] {
			total2 += fn.ParseInt(string(word[i])) * 2
		}
	}

	return NewSolution(total1, total2)
}

func data01(full bool) string {
	return ReadFirstLine(17, 1, full)
}
