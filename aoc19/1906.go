package aoc19

import (
	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/ds"
)

func Day06() Solution {
	t := data06(true)

	// Part 1
	total := bfsTraversal(t, "COM", nil)

	// Part 2
	start := t.parent["YOU"]
	goal := t.parent["SAN"]
	depth := bfsTraversal(t, start, &goal)

	return NewSolution(total, depth)
}

func data06(full bool) *Tree {
	t := &Tree{
		nodes:  ds.NewSet[string](),
		edges:  make(map[string][]string),
		parent: make(map[string]string),
	}
	for _, line := range ReadLines(19, 6, full) {
		p := fn.CleanSplit(line, ")")
		node, child := p[0], p[1]
		t.nodes.Add(node)
		t.nodes.Add(child)
		t.edges[node] = append(t.edges[node], child)
		t.edges[child] = append(t.edges[child], node)
		t.parent[child] = node
	}
	return t
}

type Tree struct {
	nodes  *ds.Set[string]
	edges  map[string][]string
	parent map[string]string
}

func bfsTraversal(t *Tree, start string, goal *string) int {
	visited := make(map[string]int)
	q := ds.NewQueue[StrInt]()
	q.Enqueue(StrInt{Str: start, Int: 0})
	for q.Len() > 0 {
		item := q.Dequeue()
		node, depth := item.Str, item.Int
		if fn.HasKey(visited, node) {
			continue
		}
		visited[node] = depth
		if goal != nil && node == *goal {
			return depth
		}
		for _, node2 := range t.edges[node] {
			if !fn.HasKey(visited, node2) {
				q.Enqueue(StrInt{Str: node2, Int: depth + 1})
			}
		}
	}

	return fn.Sum(fn.MapValues(visited))
}
