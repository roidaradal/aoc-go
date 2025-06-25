package aoc16

import (
	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/ds"
)

func Day01() Solution {
	moves := data01(true)

	// Part 1
	hq := findHQ(moves, false)
	dist1 := ManhattanOrigin(hq)

	// Part 2
	hq = findHQ(moves, true)
	dist2 := ManhattanOrigin(hq)

	return NewSolution(dist1, dist2)
}

const (
	left  int = -1
	right int = 1
)

func data01(full bool) []Int2 {
	T := map[byte]int{
		'L': left,
		'R': right,
	}
	line := ReadFirstLine(16, 1, full)
	return fn.Map(fn.CleanSplit(line, ","), func(move string) Int2 {
		turn := T[move[0]]
		steps := fn.ParseInt(move[1:])
		return Int2{turn, steps}
	})
}

func findHQ(moves []Int2, atVisitedTwice bool) Coords {
	curr := Coords{0, 0}
	d := X
	visited := ds.NewSet[Coords]()
	for _, move := range moves {
		turn, steps := move[0], move[1]
		if d == X {
			d = fn.Ternary(turn == left, L, R)
		} else if turn == left {
			d = LeftOf[d]
		} else if turn == right {
			d = RightOf[d]
		}
		for range steps {
			curr = Move(curr, d)
			if atVisitedTwice && visited.Contains(curr) {
				return curr
			}
			visited.Add(curr)
		}
	}
	return curr
}
