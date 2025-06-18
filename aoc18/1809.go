package aoc18

import (
	"slices"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/ds"
)

func Day09() Solution {
	numPlayers, lastMarble := data09(true)

	// Part 1
	score1 := maxScore(numPlayers, lastMarble)

	// Part 2
	score2 := maxScore(numPlayers, lastMarble*100)

	return NewSolution(score1, score2)
}

func data09(full bool) (int, int) {
	line := ReadFirstLine(18, 9, full)
	p := fn.SpaceSplit(line)
	numPlayers := fn.ParseInt(p[0])
	lastMarble := fn.ParseInt(Last(p, 2))
	return numPlayers, lastMarble
}

func maxScore(numPlayers, lastMarble int) int {
	score := make([]int, numPlayers)
	player := 0
	curr := ds.NewDLLNode(0)

	for m := 1; m <= lastMarble; m++ {
		if m%23 == 0 {
			prev7 := curr
			for range 7 {
				prev7 = prev7.Prev
			}
			score[player] += m + prev7.Value
			curr = prev7.Next
			prev7.Remove()
		} else {
			next1 := curr.Next
			next2 := next1.Next
			curr = next1.AddBetween(next2, m)
		}
		player = (player + 1) % numPlayers
	}
	return slices.Max(score)
}
