package aoc16

import (
	"fmt"
	"slices"

	. "github.com/roidaradal/aoc-go/aoc"
)

func Day06() {
	words := data06(true)
	numCols := len(words[0])
	freq := columnFrequency(words, numCols)
	maxMsg := make([]rune, numCols)
	minMsg := make([]rune, numCols)
	for col := range numCols {
		colFreq := make([]CharInt, 0)
		for k, v := range freq[col] {
			colFreq = append(colFreq, CharInt{Char: k, Int: v})
		}
		slices.SortFunc(colFreq, SortCharIntAsc)
		minPair := colFreq[0]
		maxPair := colFreq[len(colFreq)-1]
		minMsg[col] = minPair.Char
		maxMsg[col] = maxPair.Char
	}
	fmt.Println(string(maxMsg))
	fmt.Println(string(minMsg))
}

func data06(full bool) []string {
	return ReadLines(full)
}

func columnFrequency(words []string, numCols int) map[int]map[rune]int {
	freq := make(map[int]map[rune]int)
	for i := range numCols {
		freq[i] = make(map[rune]int)
	}
	for _, word := range words {
		for col, char := range word {
			freq[col][char] += 1
		}
	}
	return freq
}
