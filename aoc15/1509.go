package aoc15

import (
	"math"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/ds"
)

func Day09() Solution {
	g := data09(true)
	minDistance := math.MaxInt
	maxDistance := 0
	vertices := g.Vertices()
	for _, path := range Permutations(vertices, len(vertices)) {
		distance := computeDistance(path, g)
		// Part 1
		minDistance = min(minDistance, distance)
		// Part 2
		maxDistance = max(maxDistance, distance)
	}
	return NewSolution(minDistance, maxDistance)
}

func data09(full bool) *ds.Graph {
	g := ds.NewGraph()
	for _, line := range ReadLines(15, 9, full) {
		p := fn.CleanSplit(line, "=")
		v := fn.CleanSplit(p[0], "to")
		w := fn.ParseInt(p[1])
		v1, v2 := v[0], v[1]
		g.AddVertex(v1)
		g.AddVertex(v2)
		g.AddEdgeWeight(ds.VertexPair{v1, v2}, w)
		g.AddEdgeWeight(ds.VertexPair{v2, v1}, w)
	}
	return g
}

func computeDistance(path []string, g *ds.Graph) int {
	total := 0
	for i := 1; i < len(path); i++ {
		pair := ds.VertexPair{path[i-1], path[i]}
		total += g.EdgeWeight(pair)
	}
	return total
}
