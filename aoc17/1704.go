package aoc17

import (
	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day04() Solution {
	phrases := data04(true)
	count1, count2 := 0, 0
	for _, phrase := range phrases {
		// Part 1
		if isValidPhrase(phrase, false) {
			count1 += 1
		}

		// Part 2
		if isValidPhrase(phrase, true) {
			count2 += 1
		}
	}
	return NewSolution(count1, count2)
}

func data04(full bool) [][]string {
	return fn.Map(ReadLines(17, 4, full), fn.SpaceSplit)
}

func isValidPhrase(phrase []string, sortWord bool) bool {
	freq := make(map[string]int)
	for _, word := range phrase {
		if sortWord {
			word = SortedString(word)
		}
		freq[word] += 1
	}
	return fn.All(fn.MapValues(freq), func(count int) bool {
		return count == 1
	})
}
