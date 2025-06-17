package aoc15

import (
	"slices"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day02() Solution {
	items := data02(true)
	total1, total2 := 0, 0
	for _, dims := range items {
		// Part 1
		l, w, h := dims[0], dims[1], dims[2]
		lw, wh, lh := l*w, w*h, l*h
		total1 += (2 * lw) + (2 * wh) + (2 * lh) + min(lw, wh, lh)

		// Part 2
		d := dims[:]
		slices.Sort(d)
		total2 += (2 * (d[0] + d[1])) + (lw * h)
	}
	return NewSolution(total1, total2)
}

func data02(full bool) []Dims3 {
	return fn.Map(ReadLines(15, 2, full), func(line string) Dims3 {
		return ToDims3(line, "x")
	})
}
