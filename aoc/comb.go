package aoc

import (
	"iter"

	"gonum.org/v1/gonum/stat/combin"
)

func Permutations[T any](items []T, take int) iter.Seq2[int, []T] {
	return comboIterator(items, take, combin.Permutations)
}

func Combinations[T any](items []T, take int) iter.Seq2[int, []T] {
	return comboIterator(items, take, combin.Combinations)
}

func comboIterator[T any](items []T, take int, comboFn func(int, int) [][]int) iter.Seq2[int, []T] {
	return func(yield func(int, []T) bool) {
		for i, indexes := range comboFn(len(items), take) {
			combo := createCombo(items, indexes)
			if !yield(i, combo) {
				return
			}
		}
	}
}

func createCombo[T any](items []T, indexes []int) []T {
	combo := make([]T, len(indexes))
	for i, idx := range indexes {
		combo[i] = items[idx]
	}
	return combo
}
