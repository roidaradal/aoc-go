package aoc

import (
	"crypto/md5"
	"fmt"
	"slices"
)

func MD5Hash(word string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(word)))
}

func HasTwins(word string, gap int) bool {
	back := gap + 1
	limit := len(word)
	for i := back; i < limit; i++ {
		if word[i] == word[i-back] {
			return true
		}
	}
	return false
}

func CharFreq(word string, skip []rune) map[rune]int {
	freq := make(map[rune]int)
	for _, char := range word {
		if skip != nil && slices.Contains(skip, char) {
			continue
		}
		freq[char] += 1
	}
	return freq
}
