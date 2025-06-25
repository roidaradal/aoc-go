package aoc20

import (
	"strings"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/ds"
)

type Hierarchy = map[string][]StrInt

func Day07() Solution {
	h := data07(true)

	// Part 1
	parents := make(map[string][]string)
	for parentColor, bagCounts := range h {
		for _, bag := range bagCounts {
			color := bag.Str
			parents[color] = append(parents[color], parentColor)
		}
	}
	valid := ds.NewSet[string]()
	q := ds.QueueFrom(parents["shiny gold"])
	for q.Len() > 0 {
		color := q.Dequeue()
		valid.Add(color)
		for _, nxtColor := range parents[color] {
			q.Enqueue(nxtColor)
		}
	}
	count1 := valid.Len()

	// Part 2
	count2 := countInside("shiny gold", h)

	return NewSolution(count1, count2)
}

func data07(full bool) Hierarchy {
	h := make(Hierarchy)
	for _, line := range ReadLines(20, 7, full) {
		p := fn.CleanSplit(line, "contain")
		head, tail := p[0], p[1]
		if tail == "no other bags." {
			continue
		}
		p = fn.SpaceSplit(head)
		p = p[:len(p)-1] // remove 'bags'
		color := strings.Join(p, " ")
		bags := fn.CleanSplit(tail, ",")
		h[color] = fn.Map(bags, bagCount)
	}
	return h
}

func bagCount(text string) StrInt {
	p := fn.SpaceSplit(text)
	color := strings.Join(p[1:len(p)-1], " ")
	count := fn.ParseInt(p[0])
	return StrInt{Str: color, Int: count}
}

func countInside(color string, h Hierarchy) int {
	total := 0
	if fn.HasKey(h, color) {
		for _, bag := range h[color] {
			color2, count := bag.Tuple()
			total += count + (count * countInside(color2, h))
		}
	}
	return total
}
