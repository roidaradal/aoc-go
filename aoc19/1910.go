package aoc19

import (
	"cmp"
	"slices"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/ds"
)

type Line = [2]float64

func Day10() Solution {
	asteroids := data10(true)

	// Part 1
	visible := make(map[Coords]map[Coords]bool)
	for _, a1 := range asteroids {
		visible[a1] = make(map[Coords]bool)
		for _, a2 := range asteroids {
			visible[a1][a2] = false
		}
	}
	sameRow := make(map[int]*ds.Set[Coords])
	sameCol := make(map[int]*ds.Set[Coords])
	sameLine := make(map[Line]*ds.Set[Coords])
	for _, combo := range Combinations(asteroids, 2) {
		a1, a2 := combo[0], combo[1]
		y1, x1 := a1[0], a1[1]
		y2, x2 := a2[0], a2[1]
		if y1 == y2 {
			if !fn.HasKey(sameRow, y1) {
				sameRow[y1] = ds.NewSet[Coords]()
			}
			sameRow[y1].Add(a1)
			sameRow[y1].Add(a2)
		} else if x1 == x2 {
			if !fn.HasKey(sameCol, x1) {
				sameCol[x1] = ds.NewSet[Coords]()
			}
			sameCol[x1].Add(a1)
			sameCol[x1].Add(a2)
		} else {
			mb := lineEq(a1, a2)
			if !fn.HasKey(sameLine, mb) {
				sameLine[mb] = ds.NewSet[Coords]()
			}
			sameLine[mb].Add(a1)
			sameLine[mb].Add(a2)
		}
	}

	oneRow := make(map[int][]Coords)
	oneCol := make(map[int][]Coords)
	oneLine := make(map[Line][]Coords)
	for row := range sameRow {
		coords := sameRow[row].Items()
		slices.SortFunc(coords, SortCoordsX)
		oneRow[row] = coords
	}
	for col := range sameCol {
		coords := sameCol[col].Items()
		slices.SortFunc(coords, SortCoordsY)
		oneCol[col] = coords
	}
	for mb := range sameLine {
		coords := sameLine[mb].Items()
		slices.SortFunc(coords, SortCoordsX)
		oneLine[mb] = coords
	}

	lines := make([][]Coords, 0)
	lines = append(lines, fn.MapValues(oneRow)...)
	lines = append(lines, fn.MapValues(oneCol)...)
	lines = append(lines, fn.MapValues(oneLine)...)
	for _, coords := range lines {
		for i := range len(coords) - 1 {
			a1, a2 := coords[i], coords[i+1]
			visible[a1][a2] = true
			visible[a2][a1] = true
		}
	}

	scores := make([]AsteroidScore, len(asteroids))
	for i, a := range asteroids {
		vis := fn.Filter(fn.MapValues(visible[a]), func(ok bool) bool {
			return ok
		})
		scores[i] = AsteroidScore{
			asteroid: a,
			score:    len(vis),
		}
	}
	maxEntry := slices.MaxFunc(scores, sortAsteroidScore)
	maxScore := maxEntry.score

	// Part 2
	station := maxEntry.asteroid
	y, x := station[0], station[1]
	hrzn := oneRow[y]
	vert := oneCol[x]
	diag := make([]LinePoints, 0)
	for mb, pts := range oneLine {
		if slices.Contains(pts, station) {
			diag = append(diag, LinePoints{
				slope:  mb[0],
				points: pts,
			})
		}
	}
	slices.SortFunc(diag, func(a, b LinePoints) int {
		return cmp.Compare(a.slope, b.slope)
	})
	ndiag := fn.Map(fn.Filter(diag, func(lp LinePoints) bool {
		return lp.slope < 0
	}), func(lp LinePoints) []Coords {
		return lp.points
	})
	pdiag := fn.Map(fn.Filter(diag, func(lp LinePoints) bool {
		return lp.slope > 0
	}), func(lp LinePoints) []Coords {
		return lp.points
	})

	goal := 200
	destroyed := make([]Coords, 0, goal)
	quadrant := 1
	for len(destroyed) < goal {
		if quadrant == 1 {
			idx := slices.Index(vert, station)
			if idx > 0 {
				a := vert[idx-1]
				destroyed = UniqueAppend(destroyed, a)
				vert = RemoveIndex(vert, idx-1)
			}
			for p, pts := range ndiag {
				idx := slices.Index(pts, station)
				if idx == len(pts)-1 {
					continue
				}
				a := pts[idx+1]
				destroyed = UniqueAppend(destroyed, a)
				ndiag[p] = RemoveIndex(pts, idx+1)
			}
			quadrant = 2
		} else if quadrant == 2 {
			idx := slices.Index(hrzn, station)
			if idx < len(hrzn)-1 {
				a := hrzn[idx+1]
				destroyed = UniqueAppend(destroyed, a)
				hrzn = RemoveIndex(hrzn, idx+1)
			}
			for p, pts := range pdiag {
				idx := slices.Index(pts, station)
				if idx == len(pts)-1 {
					continue
				}
				a := pts[idx+1]
				destroyed = UniqueAppend(destroyed, a)
				pdiag[p] = RemoveIndex(pts, idx+1)
			}
			quadrant = 3
		} else if quadrant == 3 {
			idx := slices.Index(vert, station)
			if idx < len(vert)-1 {
				a := vert[idx+1]
				destroyed = UniqueAppend(destroyed, a)
				vert = RemoveIndex(vert, idx+1)
			}
			for p, pts := range ndiag {
				idx := slices.Index(pts, station)
				if idx == 0 {
					continue
				}
				a := pts[idx-1]
				destroyed = UniqueAppend(destroyed, a)
				ndiag[p] = RemoveIndex(pts, idx-1)
			}
			quadrant = 4
		} else if quadrant == 4 {
			idx := slices.Index(hrzn, station)
			if idx > 0 {
				a := hrzn[idx-1]
				destroyed = UniqueAppend(destroyed, a)
				hrzn = RemoveIndex(hrzn, idx-1)
			}
			for p, pts := range pdiag {
				idx := slices.Index(pts, station)
				if idx == 0 {
					continue
				}
				a := pts[idx-1]
				destroyed = UniqueAppend(destroyed, a)
				pdiag[p] = RemoveIndex(pts, idx-1)
			}
			quadrant = 1
		}
	}

	lastAsteroid := destroyed[goal-1]
	y, x = lastAsteroid[0], lastAsteroid[1]
	score := (x * 100) + y

	return NewSolution(maxScore, score)
}

func data10(full bool) []Coords {
	asteroids := make([]Coords, 0)
	for row, line := range ReadLines(19, 10, full) {
		for col, char := range []rune(line) {
			if char == '#' {
				asteroids = append(asteroids, Coords{row, col})
			}
		}
	}
	return asteroids
}

func lineEq(a1, a2 Coords) Line {
	asteroids := []Coords{a1, a2}
	slices.SortFunc(asteroids, SortCoordsX)
	a1, a2 = asteroids[0], asteroids[1]

	y1, x1 := a1[0], a1[1]
	y2, x2 := a2[0], a2[1]
	dy := float64(y2 - y1)
	dx := float64(x2 - x1)
	m := dy / dx
	b := float64(y1) - (m * float64(x1))
	return Line{m, b}
}

type AsteroidScore struct {
	asteroid Coords
	score    int
}

func sortAsteroidScore(a, b AsteroidScore) int {
	cmp1 := cmp.Compare(a.score, b.score)
	if cmp1 != 0 {
		return cmp1
	}
	cmp2 := cmp.Compare(a.asteroid[0], b.asteroid[0])
	if cmp2 != 0 {
		return cmp2
	}
	return cmp.Compare(a.asteroid[1], b.asteroid[1])
}

type LinePoints struct {
	slope  float64
	points []Coords
}
