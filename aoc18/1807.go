package aoc18

import (
	"slices"
	"strings"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/ds"
)

func Day07() Solution {
	g := data07(true)
	limit := len(g.Vertices())

	// Part 1
	done := make([]string, 0)
	for range limit {
		task := nextTask(g, done)
		done = append(done, task)
	}
	order := strings.Join(done, "")

	// Part 2
	fixed, workers := 60, 5
	timer := 0
	done = make([]string, 0)
	ongoing := ds.NewSet[string]()
	queue := make(map[int]StrInt)
	for i := range workers {
		queue[i] = EmptyTask
	}
	for len(done) < limit {
		candidates := fn.Filter(taskCandidates(g, done), func(task string) bool {
			return !ongoing.Contains(task)
		})
		availableEntries := fn.Filter(fn.MapEntries(queue), func(e fn.Entry[int, StrInt]) bool {
			return e.Value.Str == ""
		})
		available := fn.Map(availableEntries, func(e fn.Entry[int, StrInt]) int {
			return e.Key
		})
		count := min(len(candidates), len(available))
		for i := range count {
			worker, task := available[i], candidates[i]
			queue[worker] = StrInt{Str: task, Int: taskDuration(task) + fixed}
			ongoing.Add(task)
		}

		for worker := range workers {
			work := queue[worker]
			task, left := work.Str, work.Int
			if task == "" {
				continue
			}
			left -= 1
			if left == 0 {
				queue[worker] = EmptyTask
				done = append(done, task)
				ongoing.Delete(task)
			} else {
				queue[worker] = StrInt{Str: task, Int: left}
			}
		}

		timer += 1
	}

	return NewSolution(order, timer)
}

var EmptyTask = StrInt{Str: "", Int: 0}

func data07(full bool) *ds.Graph {
	g := ds.NewGraph()
	for _, line := range ReadLines(18, 7, full) {
		p := fn.SpaceSplit(line)
		v1, v2 := p[1], Last(p, 3)
		g.AddVertex(v1)
		g.AddVertex(v2)
		g.AddEdge(v2, v1)
	}
	g.InitEdges()
	return g
}

func taskCandidates(g *ds.Graph, done []string) []string {
	candidates := make([]string, 0)
	for _, vertex := range g.Vertices() {
		if slices.Contains(done, vertex) {
			continue
		}
		allDone := fn.All(g.Edges(vertex), func(dep string) bool {
			return slices.Contains(done, dep)
		})
		if allDone {
			candidates = append(candidates, vertex)
		}
	}
	slices.Sort(candidates)
	return candidates
}

func nextTask(g *ds.Graph, done []string) string {
	return taskCandidates(g, done)[0]
}

func taskDuration(task string) int {
	return int(task[0] - 64)
}
