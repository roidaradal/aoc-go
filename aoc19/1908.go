package aoc19

import (
	"fmt"
	"slices"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day08() Solution {
	layers, dims := data08(true)

	// Part 1
	freq := make(map[int]map[rune]int)
	for i, layer := range layers {
		freq[i] = CharFreq(layer, nil)
	}
	entries := fn.Map(fn.MapKeys(freq), func(i int) Int2 {
		return Int2{freq[i]['0'], i}
	})
	min0 := slices.MinFunc(entries, SortInt2)[1]
	f := freq[min0]
	part1 := f['1'] * f['2']

	// Part 2
	h, w := dims[0], dims[1]
	img := make([]rune, 0)
	for i := range h * w {
		for _, layer := range layers {
			pixels := []rune(layer)
			if pixels[i] != '2' {
				img = append(img, pixels[i])
				break
			}
		}
	}

	T := map[rune]rune{
		'1': '#',
		'0': ' ',
	}
	for i := 0; i < len(img); i += w {
		row := fn.Translate(img[i:i+w], T)
		fmt.Println(string(row))
	}

	return NewSolution(part1, "")
}

func data08(full bool) ([]string, Dims2) {
	dims := Dims2{6, 25}
	h, w := dims[0], dims[1]
	layer := h * w
	line := ReadFirstLine(19, 8, full)
	img := make([]string, 0)
	for i := 0; i < len(line); i += layer {
		img = append(img, line[i:i+layer])
	}
	return img, dims
}
