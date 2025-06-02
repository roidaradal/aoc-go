package aoc

import "github.com/roidaradal/fn"

func ToIntList(line string, sep string) []int {
	return fn.Map(fn.CleanSplit(line, sep), func(x string) int {
		return fn.ParseInt(x)
	})
}

func ToDims3(line string, sep string) Dims3 {
	p := ToIntList(line, sep)
	return Dims3{p[0], p[1], p[2]}
}
