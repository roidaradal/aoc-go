package aoc

import (
	"cmp"
	"slices"
)

func Last[T any](items []T) T {
	return items[len(items)-1]
}

func ArgMax(items []int) int {
	pairs := make([]Int2, len(items))
	for i, x := range items {
		pairs = append(pairs, Int2{x, i})
	}
	best := slices.MaxFunc(pairs, func(a, b Int2) int {
		cmp1 := cmp.Compare(a[0], b[0])
		if cmp1 != 0 {
			return cmp1
		}
		return -cmp.Compare(a[1], b[1])
	})
	return best[1]
}
