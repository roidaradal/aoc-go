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

func RemoveIndex[T any](items []T, index int) []T {
	return append(items[:index], items[index+1:]...)
}

func UniqueAppend[T comparable](items []T, item T) []T {
	if !slices.Contains(items, item) {
		items = append(items, item)
	}
	return items
}

func NumRange(start, end int) []int {
	numbers := make([]int, 0)
	for i := start; i < end; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func NumRangeInc(start, end, inc int) []int {
	numbers := make([]int, 0)
	for i := start; i < end; i += inc {
		numbers = append(numbers, i)
	}
	return numbers
}

func RevRange(start, end int) []int {
	numbers := make([]int, 0)
	for i := start; i > end; i-- {
		numbers = append(numbers, i)
	}
	return numbers
}

func CountFreq[T comparable](items []T) map[T]int {
	freq := make(map[T]int)
	for _, item := range items {
		freq[item] += 1
	}
	return freq
}

func MergeRanges(ranges []Int2) []Int2 {
	slices.SortFunc(ranges, SortInt2)
	result := make([]Int2, 0)
	currStart, currEnd := ranges[0].Tuple()
	for _, r := range ranges[1:] {
		nextStart, nextEnd := r.Tuple()
		if currStart <= nextStart && nextEnd <= currEnd {
			continue // subset
		}

		if nextStart <= currEnd {
			currEnd = nextEnd // overlap
		} else {
			result = append(result, Int2{currStart, currEnd})
			currStart, currEnd = nextStart, nextEnd
		}
	}
	result = append(result, Int2{currStart, currEnd})
	return result
}
