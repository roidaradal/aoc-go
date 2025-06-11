package aoc17

import (
	"fmt"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day01() {
	word := data01(true)
	limit := len(word)
	total := 0
	for i := range limit {
		j := (i + 1) % limit
		if word[i] == word[j] {
			total += fn.ParseInt(string(word[i]))
		}
	}
	fmt.Println(total)

	mid := limit / 2
	total = 0
	for i := range mid {
		j := mid + i
		if word[i] == word[j] {
			total += fn.ParseInt(string(word[i])) * 2
		}
	}
	fmt.Println(total)
}

func data01(full bool) string {
	return ReadLines(full)[0]
}
