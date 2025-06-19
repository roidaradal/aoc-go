package aoc19

import (
	"slices"
	"strconv"

	"github.com/roidaradal/fn"
)

func commandParts(number int) (string, string) {
	word := strconv.Itoa(number)
	n := len(word)
	if n <= 2 {
		return "", word
	} else {
		return word[:n-2], word[n-2:]
	}
}

func intcodeModes(cmd string, count int) []int {
	m := make([]int, count)
	i := 0
	digits := []rune(cmd)
	slices.Reverse(digits)
	for _, x := range digits {
		m[i] = fn.ParseInt(string(x))
		i += 1
	}
	return m
}

func intcodeParam(x int, mode int, numbers []int) int {
	if mode == 0 {
		return numbers[x]
	}
	return x
}
