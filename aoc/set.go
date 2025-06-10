package aoc

func SetIntersection[T comparable](set1, set2 map[T]bool) []T {
	common := make([]T, 0)
	for k := range set2 {
		if _, ok := set1[k]; ok {
			common = append(common, k)
		}
	}
	return common
}
