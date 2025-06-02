package aoc15

import (
	"fmt"
	"strings"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day05() {
	words := data05(true)
	count1, count2 := 0, 0
	for _, word := range words {
		if isNice(word) {
			count1 += 1
		}
		if isNice2(word) {
			count2 += 1
		}
	}
	fmt.Println(count1)
	fmt.Println(count2)
}

func data05(full bool) []string {
	return ReadLines(full)
}

var (
	invalids = []string{"ab", "cd", "pq", "xy"}
	vowels   = []rune{'a', 'e', 'i', 'o', 'u'}
)

func isNice(word string) bool {
	for _, invalid := range invalids {
		if strings.Contains(word, invalid) {
			return false
		}
	}

	if !HasTwins(word, 0) {
		return false
	}

	freq := CharFreq(word, nil)
	numVowels := fn.Sum(fn.Map(vowels, func(vowel rune) int {
		return freq[vowel]
	}))
	return numVowels >= 3
}

func isNice2(word string) bool {
	if !HasTwins(word, 1) {
		return false
	}

	pairs := substringPositions(word, 2)
	for _, idxs := range pairs {
		if len(idxs) >= 3 {
			return true
		} else if len(idxs) == 2 && fn.Abs(idxs[0]-idxs[1]) >= 2 {
			return true
		}
	}
	return false
}

func substringPositions(word string, length int) map[string][]int {
	at := make(map[string][]int)
	limit := len(word) - (length - 1)
	for i := range limit {
		sub := word[i : i+length]
		if _, ok := at[sub]; !ok {
			at[sub] = make([]int, 0)
		}
		at[sub] = append(at[sub], i)
	}
	return at
}
