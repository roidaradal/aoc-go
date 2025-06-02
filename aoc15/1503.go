package aoc15

import (
	"fmt"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day03() {
	moves := data03(true)

	// Part 1
	visited := walk(moves)
	fmt.Println(len(visited))

	// Part 2
	limit := len(moves)
	santa := make([]Delta, 0)
	robo := make([]Delta, 0)
	i := 0
	for i < limit {
		santa = append(santa, moves[i])
		robo = append(robo, moves[i+1])
		i += 2
	}
	visited = walk(santa)
	visited2 := walk(robo)
	for c := range visited2 {
		visited[c] = true
	}
	fmt.Println(len(visited))
}

func data03(full bool) []Delta {
	T := map[rune]Delta{
		'>': R,
		'<': L,
		'^': U,
		'v': D,
	}
	line := ReadLines(full)[0]
	return fn.Map([]rune(line), func(x rune) Delta {
		return T[x]
	})
}

func walk(moves []Delta) map[Coords]bool {
	var start Coords = [2]int{0, 0}
	visited := make(map[Coords]bool)
	visited[start] = true
	curr := start
	for _, delta := range moves {
		curr = Move(curr, delta)
		visited[curr] = true
	}
	return visited
}
