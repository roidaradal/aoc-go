package aoc

import "cmp"

func SortIntDesc(a, b int) int {
	return cmp.Compare(b, a)
}

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

func SortInt2(a, b Int2) int {
	cmp0 := cmp.Compare(a[0], b[0])
	if cmp0 != 0 {
		return cmp0
	}
	return cmp.Compare(a[1], b[1])
}

func SortInt3(a, b Int3) int {
	cmp0 := cmp.Compare(a[0], b[0])
	if cmp0 != 0 {
		return cmp0
	}
	cmp1 := cmp.Compare(a[1], b[1])
	if cmp1 != 0 {
		return cmp1
	}
	return cmp.Compare(a[2], b[2])
}

func SortCoordsX(a, b Coords) int {
	return cmp.Compare(a[1], b[1])
}

func SortCoordsY(a, b Coords) int {
	return cmp.Compare(a[0], b[0])
}
