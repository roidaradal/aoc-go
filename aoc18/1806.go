package aoc18

import (
	"slices"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/ds"
)

func Day06() Solution {
	points := data06(true)
	rows := slices.Max(fn.Map(points, GetCoordsY)) + 1
	cols := slices.Max(fn.Map(points, GetCoordsX)) + 1

	// Part 1
	count := make(map[int]int, len(points))
	for i := range len(points) {
		count[i] = 1
	}
	edge := ds.NewSet[int]()
	for row := range rows {
		for col := range cols {
			pt := Coords{row, col}
			if slices.Contains(points, pt) {
				continue
			}
			idx := findClosest(pt, points)
			if idx != nil {
				count[*idx] += 1
				if row == 0 || col == 0 || row == rows-1 || col == cols-1 {
					edge.Add(*idx)
				}
			}
		}
	}
	entries := make([]Int2, 0)
	for i, x := range count {
		if !edge.Contains(i) {
			entries = append(entries, Int2{x, i})
		}
	}
	maxArea := slices.MaxFunc(entries, SortInt2)[0]

	// Part 2
	goal := 10_000
	valid := 0
	for row := range rows {
		for col := range cols {
			c := Coords{row, col}
			total := fn.Sum(fn.Map(points, func(pt Coords) int {
				return Manhattan(c, pt)
			}))
			if total < goal {
				valid += 1
			}
		}
	}

	return NewSolution(maxArea, valid)
}

func data06(full bool) []Coords {
	return fn.Map(ReadLines(18, 6, full), func(line string) Coords {
		c := ToInt2(line, ",")
		return Coords{c[1], c[0]}
	})
}

func findClosest(c Coords, points []Coords) *int {
	dist := make(map[int]int)
	for i, pt := range points {
		dist[i] = Manhattan(c, pt)
	}
	entries := make([]Int2, 0, len(dist))
	for i, d := range dist {
		entries = append(entries, Int2{d, i})
	}
	minDist := slices.MinFunc(entries, SortInt2)[0]
	minEntries := fn.Filter(entries, func(e Int2) bool {
		return e[0] == minDist
	})
	if len(minEntries) != 1 {
		return nil
	}
	return &minEntries[0][1]
}
