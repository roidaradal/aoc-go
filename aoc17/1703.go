package aoc17

import (
	"math"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day03() Solution {
	x := data03(true)

	// Part 1
	c := spiralCoords(x)
	dist := ManhattanOrigin(c)

	// Part 2
	goal := x
	spiral := map[Coords]int{
		{0, 0}: 1,
	}
	values := []int{0, 1}
	x, value := 2, 0
	for value <= goal {
		curr := spiralCoords(x)
		near := fn.Filter(Surround8(curr), func(c Coords) bool {
			return fn.HasKey(spiral, c)
		})
		value = fn.Sum(fn.Map(near, func(c Coords) int {
			return values[spiral[c]]
		}))
		values = append(values, value)
		spiral[curr] = x
		x += 1
	}

	return NewSolution(dist, value)
}

func data03(full bool) int {
	return fn.ParseInt(ReadFirstLine(17, 3, full))
}

func spiralCoords(x int) Coords {
	layer := spiralLayer(x)
	side, offset := spiralOffset(x, layer)
	if side == 'T' || side == 'L' {
		layer = -layer
	}
	if side == 'T' || side == 'B' {
		return Coords{layer, offset}
	} else {
		return Coords{offset, layer}
	}
}

func spiralLayer(x int) int {
	dims := int(math.Ceil(math.Sqrt(float64(x))))
	if dims%2 == 0 {
		dims += 1
	}
	return (dims - 1) / 2
}

func spiralOffset(x int, layer int) (byte, int) {
	side := "BLTR"
	corners := spiralCorners(layer)
	for i := range len(corners) - 1 {
		c2, c1 := corners[i], corners[i+1]
		if !(c2 >= x && x >= c1) {
			continue
		}

		mid := (c1 + c2) / 2
		offset := mid - x
		if i < 2 {
			offset = -offset
		}
		return side[i], offset
	}
	return 'C', 0
}

func spiralCorners(layer int) []int {
	if layer == 0 {
		return []int{}
	}
	dims := (layer * 2) + 1
	step := dims - 1
	square := int(math.Pow(float64(dims), 2))
	corners := []int{square}
	for i := range 4 {
		corners = append(corners, corners[i]-step)
	}
	return corners
}
