package aoc15

import (
	"fmt"
	"strings"

	. "github.com/roidaradal/aoc-go/aoc"
)

func Day10() {
	text := data10(true)

	// Part 1
	length := repeatExpand(text, 40)
	fmt.Println(length)

	// Part 2
	length = repeatExpand(text, 50)
	fmt.Println(length)
}

func data10(full bool) string {
	return ReadLines(full)[0]
}

func repeatExpand(text string, count int) int {
	curr := text
	for range count {
		next := make([]string, 0)
		d, r := curr[0], 1
		for i := range len(curr) {
			if i == 0 {
				continue
			}
			x := curr[i]
			if x == d {
				r += 1
			} else {
				next = append(next, fmt.Sprintf("%d", r))
				next = append(next, string(d))
				d, r = x, 1
			}
		}
		next = append(next, fmt.Sprintf("%d", r))
		next = append(next, string(d))
		curr = strings.Join(next, "")
	}
	return len(curr)
}
