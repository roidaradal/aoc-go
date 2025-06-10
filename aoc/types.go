package aoc

type (
	Dims2  = [2]int
	Dims3  = [3]int
	Delta  = [2]int
	Coords = [2]int
	Int2   = [2]int
	Str2   = [2]string
)

type (
	IntGrid = [][]int
)

type CharInt struct {
	Char rune
	Int  int
}

type StrInt struct {
	Str string
	Int int
}
