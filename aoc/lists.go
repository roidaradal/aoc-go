package aoc

func Last[T any](items []T) T {
	return items[len(items)-1]
}
