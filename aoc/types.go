package aoc

import "fmt"

type (
	Dims2    [2]int
	Dims3    [3]int
	Delta    [2]int
	Coords   [2]int
	Int2     [2]int
	Int3     [3]int
	Str2     [2]string
	Solution [2]string
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

func NewSolution[T any, V any](part1 T, part2 V) Solution {
	sol1 := fmt.Sprintf("%v", part1)
	sol2 := fmt.Sprintf("%v", part2)
	return Solution{sol1, sol2}
}

func (c CharInt) Tuple() (rune, int) {
	return c.Char, c.Int
}

func (s StrInt) Tuple() (string, int) {
	return s.Str, s.Int
}

func (d Dims2) Tuple() (int, int) {
	return d[0], d[1]
}

func (d Dims3) Tuple() (int, int, int) {
	return d[0], d[1], d[2]
}

func (d Delta) Tuple() (int, int) {
	return d[0], d[1]
}

func (c Coords) Tuple() (int, int) {
	return c[0], c[1]
}

func (p Int2) Tuple() (int, int) {
	return p[0], p[1]
}

func (p Int3) Tuple() (int, int, int) {
	return p[0], p[1], p[2]
}
