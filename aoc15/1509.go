package aoc15

import (
	"fmt"
	"math"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day09() {
	g := data09(true)
	minDistance := math.MaxInt
	maxDistance := 0
	vertices := fn.MapKeys(g.Vertices)
	for _, path := range Permutations(vertices, len(vertices)) {
		distance := computeDistance(path, g)
		minDistance = min(minDistance, distance)
		maxDistance = max(maxDistance, distance)
	}
	fmt.Println(minDistance)
	fmt.Println(maxDistance)
}

func data09(full bool) *Graph {
	g := NewGraph()
	for _, line := range ReadLines(full) {
		p := fn.CleanSplit(line, "=")
		v := fn.CleanSplit(p[0], "to")
		w := fn.ParseInt(p[1])
		v1, v2 := v[0], v[1]
		g.AddVertex(v1)
		g.AddVertex(v2)
		g.AddEdgeWeight(Str2{v1, v2}, w)
		g.AddEdgeWeight(Str2{v2, v1}, w)
	}
	return g
}

func computeDistance(path []string, g *Graph) int {
	total := 0
	for i := 1; i < len(path); i++ {
		pair := Str2{path[i-1], path[i]}
		total += g.Edges[pair]
	}
	return total
}
