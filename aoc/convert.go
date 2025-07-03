package aoc

import "github.com/roidaradal/fn"

func ToIntList(line string, sep string) []int {
	var parts []string
	if sep == " " {
		parts = fn.SpaceSplit(line)
	} else {
		parts = fn.CleanSplit(line, sep)
	}
	return fn.Map(parts, fn.ParseInt)
}

func ToIntLine(line string) []int {
	digits := make([]string, 0)
	for i := range len(line) {
		digits = append(digits, string(line[i:i+1]))
	}
	return fn.Map(digits, fn.ParseInt)
}

func ToDims3(line string, sep string) Dims3 {
	p := ToIntList(line, sep)
	return Dims3{p[0], p[1], p[2]}
}

func ToInt2(line string, sep string) Int2 {
	p := ToIntList(line, sep)
	return Int2{p[0], p[1]}
}

func ToCharInt(line string) CharInt {
	chars := []rune(line)
	return CharInt{
		Char: chars[0],
		Int:  fn.ParseInt(string(chars[1:])),
	}
}
