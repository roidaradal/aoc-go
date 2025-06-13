package aoc17

import (
	"fmt"
	"sort"
	"strings"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day07() {
	t := data07(true)
	for _, name := range t.nodes {
		if _, ok := t.parentOf[name]; !ok {
			fmt.Println(name)
			break
		}
	}

	weight := make(map[string]int)
	q := t.nodes
	for len(q) > 0 {
		q2 := make([]string, 0)
		for _, node := range q {
			if _, ok := t.children[node]; !ok {
				weight[node] = t.weight[node]
				continue
			}
			computable := fn.All(t.children[node], func(child string) bool {
				_, ok := weight[child]
				return ok
			})
			if computable {
				childWeights := make(map[int]bool)
				totalChild := 0
				for _, child := range t.children[node] {
					w := weight[child]
					childWeights[w] = true
					totalChild += w
				}
				weight[node] = t.weight[node] + totalChild

				if len(childWeights) == 2 {
					weights := fn.MapKeys(childWeights)
					sort.Ints(weights)
					target, heavy := weights[0], weights[1]
					heavyChild := ""
					for _, child := range t.children[node] {
						if weight[child] == heavy {
							heavyChild = child
							break
						}
					}
					targetWeight := t.weight[heavyChild] - (heavy - target)
					fmt.Println(targetWeight)
					return
				}
			} else {
				q2 = append(q2, node)
			}
		}
		q = q2
	}
}

func data07(full bool) Tree {
	t := Tree{
		nodes:    make([]string, 0),
		weight:   make(map[string]int),
		parentOf: make(map[string]string),
		children: make(map[string][]string),
	}
	for _, line := range ReadLines(full) {
		p := fn.CleanSplit(line, "->")
		node := fn.CleanSplit(p[0], "(")
		name := node[0]
		t.nodes = append(t.nodes, name)
		t.weight[name] = fn.ParseInt(strings.TrimSuffix(node[1], ")"))
		if len(p) == 1 {
			continue
		}
		if _, ok := t.children[name]; !ok {
			t.children[name] = make([]string, 0)
		}
		for _, child := range fn.CleanSplit(p[1], ",") {
			t.parentOf[child] = name
			t.children[name] = append(t.children[name], child)
		}
	}
	return t
}

type Tree struct {
	nodes    []string
	weight   map[string]int
	parentOf map[string]string
	children map[string][]string
}
