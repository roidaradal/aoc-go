package aoc17

import (
	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn/ds"
)

func Day09() Solution {
	stream := data09(true)

	// Part 1 and 2
	i, limit := 0, len(stream)
	count, total := 0, 0
	garbage := false
	stack := ds.NewStack[int]()
	for i < limit {
		char := stream[i]
		if garbage {
			if char == '!' {
				i += 1
			} else if char == '>' {
				garbage = false
			} else {
				count += 1
			}
		} else if char == '{' {
			score := 1
			if !stack.IsEmpty() {
				score = stack.Top() + 1
			}
			stack.Push(score)
		} else if char == '}' {
			total += stack.Pop()
		} else if char == '<' {
			garbage = true
		}
		i += 1
	}
	return NewSolution(total, count)
}

func data09(full bool) string {
	return ReadFirstLine(17, 9, full)
}
