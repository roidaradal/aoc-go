package aoc17

import (
	"fmt"

	. "github.com/roidaradal/aoc-go/aoc"
)

func Day09() {
	stream := data09(true)
	i, limit := 0, len(stream)
	count, total := 0, 0
	garbage := false
	stack := NewStack[int]()
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
	fmt.Println(total)
	fmt.Println(count)
}

func data09(full bool) string {
	return ReadLines(full)[0]
}
