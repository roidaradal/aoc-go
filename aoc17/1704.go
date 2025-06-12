package aoc17

import (
	"fmt"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day04() {
	phrases := data04(true)
	count1, count2 := 0, 0
	for _, phrase := range phrases {
		if isValidPhrase(phrase, false) {
			count1 += 1
		}
		if isValidPhrase(phrase, true) {
			count2 += 1
		}
	}
	fmt.Println(count1)
	fmt.Println(count2)
}

func data04(full bool) [][]string {
	return fn.Map(ReadLines(full), func(line string) []string {
		return fn.SpaceSplit(line)
	})
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
