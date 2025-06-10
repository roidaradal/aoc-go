package aoc

import "cmp"

func SortCharIntDesc(a, b CharInt) int {
	cmp1 := cmp.Compare(b.Int, a.Int)
	if cmp1 != 0 {
		return cmp1
	}
	return cmp.Compare(a.Char, b.Char)
}

func SortCharIntAsc(a, b CharInt) int {
	cmp1 := cmp.Compare(a.Int, b.Int)
	if cmp1 != 0 {
		return cmp1
	}
	return cmp.Compare(a.Char, b.Char)
}
