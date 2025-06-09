package aoc16

import (
	"fmt"
	"iter"
	"slices"
	"strings"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day05() {
	door := data05(true)
	pwdLength := 8

	// Part 1
	hashGen := md5HashGenerator(door, "00000", 0)
	next, stop := iter.Pull2(hashGen)
	pwdChars := make([]byte, pwdLength)
	for i := range pwdLength {
		_, hash, ok := next()
		if ok {
			pwdChars[i] = hash[5]
		}
	}
	stop()
	pwd := string(pwdChars)
	fmt.Println(pwd)

	// Part 2
	hashGen = md5HashGenerator(door, "00000", 0)
	next, stop = iter.Pull2(hashGen)
	defer stop()
	indexes := make(map[byte]bool)
	for x := range pwdLength {
		b := fmt.Sprintf("%d", x)[0]
		indexes[b] = true
	}
	pwdChars = slices.Repeat([]byte{'.'}, pwdLength)
	for {
		_, hash, ok := next()
		if !ok {
			break
		}
		if !indexes[hash[5]] {
			continue
		}
		idx := fn.ParseInt(string(hash[5]))
		if pwdChars[idx] == '.' {
			pwdChars[idx] = hash[6]
			fmt.Println(string(pwdChars))
		}
		allSet := fn.All(pwdChars, func(char byte) bool {
			return char != '.'
		})
		if allSet {
			break
		}
	}
}

func data05(full bool) string {
	return ReadLines(full)[0]
}

func md5HashGenerator(key string, goal string, start int) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		i := start
		for {
			word := fmt.Sprintf("%s%d", key, i)
			hash := MD5Hash(word)
			if strings.HasPrefix(hash, goal) {
				if !yield(i, hash) {
					return
				}
			}
			i += 1
		}
	}
}
