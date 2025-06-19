package aoc19

import (
	"slices"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day03() Solution {
	wires := data03(true)
	cross := crossingPoints(wires)

	// Part 1
	distances := fn.Map(fn.MapKeys(cross), ManhattanOrigin)
	closest1 := slices.Min(distances)

	// Part 2
	closest2 := slices.Min(fn.MapValues(cross))

	return NewSolution(closest1, closest2)
}

func data03(full bool) [][]CharInt {
	return fn.Map(ReadLines(19, 3, full), func(line string) []CharInt {
		return fn.Map(fn.CleanSplit(line, ","), ToCharInt)
	})
}

func wire(moves []CharInt) map[Coords]int {
	T := map[rune]Delta{
		'U': U,
		'D': D,
		'L': L,
		'R': R,
	}
	visited := make(map[Coords]int)
	c := Coords{0, 0}
	i := 0
	for _, m := range moves {
		d := T[m.Char]
		steps := m.Int
		for range steps {
			c = Move(c, d)
			i += 1
			if !fn.HasKey(visited, c) {
				visited[c] = i
			}
		}
	}
	return visited
}

func crossingPoints(wires [][]CharInt) map[Coords]int {
	steps := make(map[Coords][]int)
	for _, moves := range wires {
		visited := wire(moves)
		for c, x := range visited {
			steps[c] = append(steps[c], x)
		}
	}

	cross := make(map[Coords]int)
	for c, s := range steps {
		if len(s) > 1 {
			cross[c] = fn.Sum(s)
		}
	}
	return cross
}
