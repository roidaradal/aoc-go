package aoc

import "github.com/roidaradal/fn"

func ToIntList(line string, sep string) []int {
	var parts []string
	if sep == " " {
		parts = fn.SpaceSplit(line)
	} else {
		parts = fn.CleanSplit(line, sep)
	}
	return fn.Map(parts, func(x string) int {
		return fn.ParseInt(x)
	})
}

func ToDims3(line string, sep string) Dims3 {
	p := ToIntList(line, sep)
	return Dims3{p[0], p[1], p[2]}
}
