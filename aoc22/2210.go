package aoc22

import (
	"fmt"
	"slices"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

type CRT = [][]rune

func Day10() Solution {
	commands := data10(true)

	// Part 1
	interest := []int{20, 60, 100, 140, 180, 220}
	crt := createCRT()
	_, cols := GridBounds(crt).Tuple()
	x, t, total := 1, 0, 0
	for _, pair := range commands {
		cmd, param := pair.Tuple()
		steps := fn.Ternary(cmd == "addx", 2, 1)
		for range steps {
			t += 1
			drawPixel(crt, t, x, cols)
			if slices.Contains(interest, t) {
				total += t * x
			}
		}
		x += param
	}

	// Part 2
	for _, line := range crt {
		fmt.Println(string(line))
	}

	return NewSolution(total, "")
}

func data10(full bool) []StrInt {
	return fn.Map(ReadLines(22, 10, full), func(line string) StrInt {
		p := fn.SpaceSplit(line)
		cmd := p[0]
		value := 0
		if cmd == "addx" {
			value = fn.ParseInt(p[1])
		}
		return StrInt{Str: cmd, Int: value}
	})
}

func createCRT() CRT {
	rows, cols := 6, 40
	crt := make(CRT, 0)
	for range rows {
		line := slices.Repeat([]rune{'.'}, cols)
		crt = append(crt, line)
	}
	return crt
}

func drawPixel(crt CRT, t int, x int, cols int) {
	row := (t - 1) / cols
	col := (t - 1) % cols
	pixel := fn.Ternary(fn.Abs(x-col) <= 1, '#', '.')
	crt[row][col] = pixel
}
