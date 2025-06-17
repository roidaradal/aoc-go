package aoc16

import (
	"strings"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day09() Solution {
	word := data09(true)

	// Part 1
	size1 := decompressLength(word, true)

	// Part 2
	size2 := decompressLength(word, false)

	return NewSolution(size1, size2)
}

func data09(full bool) string {
	return ReadLines(16, 9, full)[0]
}

func decompressLength(word string, skip bool) int {
	wordLen := len(word)
	count := make([]int, wordLen)
	for i := range wordLen {
		count[i] = 1
	}
	i := 0
	for i < wordLen {
		if word[i] != '(' {
			i += 1
			continue
		}
		end := strings.Index(word[i:], ")") + i
		p := ToIntList(word[i+1:end], "x")
		size, repeat := p[0], p[1]
		start := end + 1
		for j := i; j < start; j++ {
			count[j] = 0
		}
		for j := start; j < start+size; j++ {
			count[j] *= repeat
		}
		if skip {
			i = start + size
		} else {
			i = end + 1
		}
	}
	return fn.Sum(count)
}
