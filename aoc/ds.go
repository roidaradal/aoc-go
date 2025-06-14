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

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	last := len(s.items) - 1
	top := s.items[last]
	s.items = s.items[:last]
	return top
}

func (s Stack[T]) Top() T {
	last := len(s.items) - 1
	return s.items[last]
}

func (s Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
