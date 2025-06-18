package aoc

import (
	"cmp"
	"slices"
)

func Last[T any](items []T, n int) T {
	return items[len(items)-n]
}

func ArgMax(items []int) int {
	pairs := make([]Int2, len(items))
	for i, x := range items {
		pairs[i] = Int2{x, i}
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

func ArgMin(items []int) int {
	pairs := make([]Int2, len(items))
	for i, x := range items {
		pairs[i] = Int2{x, i}
	}
	best := slices.MinFunc(pairs, func(a, b Int2) int {
		cmp1 := cmp.Compare(a[0], b[0])
		if cmp1 != 0 {
			return cmp1
		}
		return cmp.Compare(a[1], b[1])
	})
	return best[1]
}

func JoinLists[T any](items1 []T, items2 []T) []T {
	items := make([]T, 0)
	items = append(items, items1...)
	items = append(items, items2...)
	return items
}
