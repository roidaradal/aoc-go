package aoc23

import (
	"math"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/ds"
)

func Day04() Solution {
	games := data04(true)

	// Part 1
	total1 := 0
	for _, game := range games {
		total1 += game.score()
	}

	// Part 2
	total2 := countTotalCards(games)

	return NewSolution(total1, total2)
}

func data04(full bool) []CardGame {
	return fn.Map(ReadLines(23, 4, full), func(line string) CardGame {
		tail := fn.CleanSplit(line, ":")[1]
		t := fn.CleanSplit(tail, "|")
		winners := ToIntList(t[0], " ")
		numbers := ToIntList(t[1], " ")
		return CardGame{
			winners: ds.SetFrom(winners),
			numbers: ds.SetFrom(numbers),
		}
	})
}

type CardGame struct {
	winners *ds.Set[int]
	numbers *ds.Set[int]
}

func (g CardGame) score() int {
	common := g.winners.Intersection(g.numbers).Len()
	if common == 0 {
		return 0
	} else {
		score := math.Pow(2, float64(common-1))
		return int(score)
	}
}

func countTotalCards(games []CardGame) int {
	limit := len(games)
	count := make(map[int]int)
	for i := range limit {
		count[i] = 1
	}
	for i, game := range games {
		common := game.winners.Intersection(game.numbers).Len()
		for j := range common {
			k := i + j + 1
			if k < limit {
				count[k] += count[i]
			}
		}
	}
	return fn.Sum(fn.MapValues(count))
}
