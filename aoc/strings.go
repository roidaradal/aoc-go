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

func ReverseString(word string) string {
	chars := make([]byte, 0)
	i := len(word) - 1
	for i >= 0 {
		chars = append(chars, word[i])
		i -= 1
	}
	return string(chars)
}
