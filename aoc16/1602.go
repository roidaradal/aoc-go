package aoc16

import (
	"fmt"

	. "github.com/roidaradal/aoc-go/aoc"
)

func Day02() {
	movesList := data02(true)

	grid1 := []string{"123", "456", "789"}
	pad1 := Keypad{
		Grid:  grid1,
		Start: Coords{1, 1},
		InsideBounds: func(c Coords) bool {
			return InsideBounds(c, Dims2{3, 3})
		},
	}
	code := solveCode(pad1, movesList)
	fmt.Println(code)

	grid2 := []string{"00100", "02340", "56789", "0ABC0", "00D00"}
	pad2 := Keypad{
		Grid:  grid2,
		Start: Coords{2, 0},
		InsideBounds: func(c Coords) bool {
			row, col := c[0], c[1]
			return InsideBounds(c, Dims2{5, 5}) && grid2[row][col] != '0'
		},
	}
	code = solveCode(pad2, movesList)
	fmt.Println(code)
}

func data02(full bool) []string {
	return ReadLines(full)
}

type Keypad struct {
	Grid         []string
	Start        Coords
	InsideBounds func(Coords) bool
}

func solveCode(pad Keypad, movesList []string) string {
	T := map[rune]Delta{
		'U': U,
		'D': D,
		'L': L,
		'R': R,
	}
	code := make([]byte, 0)
	curr := pad.Start
	for _, moves := range movesList {
		for _, m := range moves {
			nxt := Move(curr, T[m])
			if pad.InsideBounds(nxt) {
				curr = nxt
			}
		}
		row, col := curr[0], curr[1]
		code = append(code, pad.Grid[row][col])
	}
	return string(code)
}
