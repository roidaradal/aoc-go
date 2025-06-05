package aoc16

import (
	"fmt"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day01() {
	moves := data01(true)

	// Part 1
	hq := findHQ(moves, false)
	fmt.Println(ManhattanOrigin(hq))

	// Part 2
	hq = findHQ(moves, true)
	fmt.Println(ManhattanOrigin(hq))
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
	line := ReadLines(full)[0]
	return fn.Map(fn.CleanSplit(line, ","), func(move string) Int2 {
		turn := T[move[0]]
		steps := fn.ParseInt(move[1:])
		return Int2{turn, steps}
	})
}

func findHQ(moves []Int2, atVisitedTwice bool) Coords {
	curr := Coords{0, 0}
	d := Delta{0, 0}
	visited := make(map[Coords]bool)
	for _, move := range moves {
		turn, steps := move[0], move[1]
		if d[0] == 0 && d[1] == 0 {
			if turn == left {
				d = L
			} else {
				d = R
			}
		} else if turn == left {
			d = LeftOf[d]
		} else if turn == right {
			d = RightOf[d]
		}
		for range steps {
			curr = Move(curr, d)
			if atVisitedTwice && visited[curr] {
				return curr
			}
			visited[curr] = true
		}
	}
	return curr
}
