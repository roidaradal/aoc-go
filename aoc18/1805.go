package aoc18

import (
	"fmt"
	"math"
	"slices"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/ds"
)

func Day05() Solution {
	word := data05(true)

	// Part 1
	word1 := fullyCompress(word)
	length := len(word1)

	// Part 2
	// charSet := ds.NewSet[byte]()
	charSet := ds.SetFrom(fn.Map([]byte(word), LowerChar))
	chars := charSet.Items()
	slices.Sort(chars)
	numChars := len(chars)
	minLength := math.MaxInt
	for i, skipChar := range chars {
		chars2 := fn.Filter([]byte(word), func(char byte) bool {
			return LowerChar(char) != skipChar
		})
		word2 := fullyCompress(string(chars2))
		wordLen := len(word2)
		minLength = min(minLength, wordLen)
		fmt.Printf("%.02d / %.02d - %c - %d\n", i+1, numChars, skipChar, wordLen)
	}

	return NewSolution(length, minLength)
}

func data05(full bool) string {
	return ReadFirstLine(18, 5, full)
}

func fullyCompress(word string) string {
	ok := true
	for ok {
		word, ok = compress(word)
	}
	return word
}

func compress(word string) (string, bool) {
	for i := range len(word) - 1 {
		x, y := word[i], word[i+1]
		if x != y && LowerChar(x) == LowerChar(y) {
			word = word[:i] + word[i+2:]
			return word, true
		}
	}
	return word, false
}
