package aoc

type Graph struct {
	Vertices map[string]bool
	Edges    map[Str2]int
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: make(map[string]bool),
		Edges:    make(map[Str2]int),
	}
}

func (g *Graph) AddVertex(v string) {
	g.Vertices[v] = true
}

func (g *Graph) AddEdgeWeight(pair Str2, weight int) {
	g.Edges[pair] = weight
}
