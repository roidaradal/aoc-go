package aoc15

import (
	"fmt"
	"strings"

	. "github.com/roidaradal/aoc-go/aoc"
)

func Day04() {
	key := data04(true)

	// Part 1
	idx := findHash(key, 5)
	fmt.Println(idx)

	// Part 2
	idx = findHash(key, 6)
	fmt.Println(idx)
}

func data04(full bool) string {
	return ReadLines(full)[0]
}

func findHash(key string, numZeros int) int {
	goal := strings.Repeat("0", numZeros)
	i := 1
	for {
		word := fmt.Sprintf("%s%d", key, i)
		hash := MD5Hash(word)
		if strings.HasPrefix(hash, goal) {
			return i
		}
		i += 1
	}
}
