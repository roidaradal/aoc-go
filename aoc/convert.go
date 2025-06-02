package aoc

import "github.com/roidaradal/fn"

func ToDims3(line string, sep string) Dims3 {
	p := fn.Map(fn.CleanSplit(line, sep), func(x string) int {
		return fn.ParseInt(x)
	})
	return [3]int{p[0], p[1], p[2]}
}
