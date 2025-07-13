package aoc23

import (
	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day02() Solution {
	games := data02(true)

	total1, total2 := 0, 0
	for _, game := range games {
		// Part 1
		if game.isValid() {
			total1 += game.id
		}

		// Part 2
		total2 += game.power()
	}

	return NewSolution(total1, total2)
}

func data02(full bool) []Game {
	return fn.Map(ReadLines(23, 2, full), func(line string) Game {
		p := fn.CleanSplit(line, ":")
		id := fn.ParseInt(fn.SpaceSplit(p[0])[1])
		game := Game{
			id:    id,
			draws: make([]Int3, 0),
		}
		for _, draw := range fn.CleanSplit(p[1], ";") {
			game.addDraw(draw)
		}
		return game
	})
}

type Game struct {
	id    int
	draws []Int3
}

func (game *Game) addDraw(line string) {
	r, g, b := 0, 0, 0
	for _, part := range fn.CleanSplit(line, ",") {
		p := fn.SpaceSplit(part)
		number := fn.ParseInt(p[0])
		color := p[1]
		switch color {
		case "red":
			r = number
		case "green":
			g = number
		case "blue":
			b = number
		}
	}
	game.draws = append(game.draws, Int3{r, g, b})
}

func (game Game) isValid() bool {
	for _, draw := range game.draws {
		r, g, b := draw.Tuple()
		if r > 12 || g > 13 || b > 14 {
			return false
		}
	}
	return true
}

func (game Game) power() int {
	maxR, maxG, maxB := 0, 0, 0
	for _, draw := range game.draws {
		r, g, b := draw.Tuple()
		maxR = max(maxR, r)
		maxG = max(maxG, g)
		maxB = max(maxB, b)
	}
	return maxR * maxG * maxB
}
