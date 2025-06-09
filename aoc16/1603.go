package aoc16

import (
	"fmt"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day03() {
	triples1 := data03(true)
	triples2 := readVertical(triples1)

	for _, triples := range [][]Dims3{triples1, triples2} {
		count := 0
		for _, triple := range triples {
			if isValidTriple(triple) {
				count += 1
			}
		}
		fmt.Println(count)
	}
}

func data03(full bool) []Dims3 {
	return fn.Map(ReadLines(full), func(line string) Dims3 {
		return ToDims3(line, " ")
	})
}

func readVertical(t []Dims3) []Dims3 {
	t2 := make([]Dims3, 0)
	for r := 0; r < len(t); r += 3 {
		for c := range 3 {
			t2 = append(t2, Dims3{
				t[r][c],
				t[r+1][c],
				t[r+2][c],
			})
		}
	}
	return t2
}

func isValidTriple(t Dims3) bool {
	a, b, c := t[0], t[1], t[2]
	return (a+b > c) && (b+c > a) && (a+c > b)
}
