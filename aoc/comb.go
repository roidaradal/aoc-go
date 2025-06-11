package aoc

import "gonum.org/v1/gonum/stat/combin"

func Permutations[T any](items []T, take int) [][]T {
	combos := make([][]T, 0)
	for _, indexes := range combin.Permutations(len(items), take) {
		combo := createCombo(items, indexes)
		combos = append(combos, combo)
	}
	return combos
}

func Combinations[T any](items []T, take int) [][]T {
	combos := make([][]T, 0)
	for _, indexes := range combin.Combinations(len(items), take) {
		combo := createCombo(items, indexes)
		combos = append(combos, combo)
	}
	return combos
}

func createCombo[T any](items []T, indexes []int) []T {
	combo := make([]T, len(indexes))
	for i, idx := range indexes {
		combo[i] = items[idx]
	}
	return combo
}
