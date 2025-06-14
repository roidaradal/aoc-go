package aoc17

import (
	"fmt"
	"slices"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day08() {
	instructions := data08(true)
	reg := make(map[string]int)
	maxVal := 0
	for _, cmd := range instructions {
		if cmd.isSatisfied(reg[cmd.condReg]) {
			reg[cmd.targetReg] += cmd.inc
			maxVal = max(maxVal, reg[cmd.targetReg])
		}
	}
	fmt.Println(slices.Max(fn.MapValues(reg)))
	fmt.Println(maxVal)
}

func data08(full bool) []Instruction {
	return fn.Map(ReadLines(full), func(line string) Instruction {
		cmd := Instruction{}
		p := fn.CleanSplit(line, " if ")
		head, tail := p[0], p[1]
		h := fn.SpaceSplit(head)
		cmd.targetReg = h[0]
		cmd.inc = fn.ParseInt(h[2])
		if h[1] == "dec" {
			cmd.inc *= -1
		}
		t := fn.SpaceSplit(tail)
		cmd.condReg = t[0]
		cmd.op = t[1]
		cmd.value = fn.ParseInt(t[2])
		return cmd
	})
}

type Instruction struct {
	targetReg string
	condReg   string
	op        string
	inc       int
	value     int
}

func (i Instruction) isSatisfied(value int) bool {
	if i.op == "==" {
		return value == i.value
	} else if i.op == "!=" {
		return value != i.value
	} else if i.op == ">" {
		return value > i.value
	} else if i.op == ">=" {
		return value >= i.value
	} else if i.op == "<" {
		return value < i.value
	} else if i.op == "<=" {
		return value <= i.value
	} else {
		return false
	}
}
