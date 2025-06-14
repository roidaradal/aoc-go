package aoc18

import (
	"fmt"
	"slices"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day02() {
	words := data02(true)
	count2, count3 := 0, 0
	for _, word := range words {
		freq := fn.MapValues(CharFreq(word, nil))
		if slices.Contains(freq, 2) {
			count2 += 1
		}
		if slices.Contains(freq, 3) {
			count3 += 1
		}
	}
	fmt.Println(count2 * count3)

	for _, pair := range Combinations(words, 2) {
		word1, word2 := pair[0], pair[1]
		diff := strDiff(word1, word2)
		if len(diff) != 1 {
			continue
		}
		idx := diff[0]
		word := word1[:idx] + word1[idx+1:]
		fmt.Println(word)
		break
	}
}

func data02(full bool) []string {
	return ReadLines(full)
}

func strDiff(word1, word2 string) []int {
	diff := make([]int, 0)
	for i := range len(word1) {
		if word1[i] != word2[i] {
			diff = append(diff, i)
		}
	}
	return diff
}
