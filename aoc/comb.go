package aoc

import "gonum.org/v1/gonum/stat/combin"

func Permutations[T any](items []T, take int) [][]T {
	combos := make([][]T, 0)
	for _, indexes := range combin.Permutations(len(items), take) {
		combo := make([]T, len(indexes))
		for i, idx := range indexes {
			combo[i] = items[idx]
		}
		combos = append(combos, combo)
	}
	return combos
}
