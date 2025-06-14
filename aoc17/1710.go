package aoc17

import (
	"fmt"
	"strings"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day10() {
	line := data10(true)
	lengths := ToIntList(line, ",")
	numbers := knotHash(lengths, 1)
	fmt.Println(numbers[0] * numbers[1])

	lengths = fn.Map([]byte(line), func(b byte) int {
		return int(b)
	})
	numbers = knotHash(lengths, 64)
	result := make([]string, 0)
	for i := 0; i < knotHashLimit; i += 16 {
		r := numbers[i]
		for j := 1; j < 16; j++ {
			r ^= numbers[i+j]
		}
		result = append(result, hexCode(r))
	}
	fmt.Println(strings.Join(result, ""))
}

func data10(full bool) string {
	return ReadLines(full)[0]
}
