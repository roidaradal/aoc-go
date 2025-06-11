package aoc17

import (
	"fmt"
	"slices"
	"sort"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day02() {
	numbersList := data02(true)
	total1, total2 := 0, 0
	for _, numbers := range numbersList {
		total1 += slices.Max(numbers) - slices.Min(numbers)

		for _, pair := range Combinations(numbers, 2) {
			sort.Ints(pair)
			a, b := pair[0], pair[1]
			if b%a == 0 {
				total2 += b / a
				break
			}
		}
	}
	fmt.Println(total1)
	fmt.Println(total2)
}

func data02(full bool) [][]int {
	return fn.Map(ReadLines(full), func(line string) []int {
		return ToIntList(line, " ")
	})
}
