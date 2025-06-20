package aoc20

import (
	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day02() Solution {
	passwords := data02(true)
	count1, count2 := 0, 0
	for _, p := range passwords {
		// Part 1
		if p.isValid() {
			count1 += 1
		}

		// Part 2
		if p.isValid2() {
			count2 += 1
		}
	}
	return NewSolution(count1, count2)
}

func data02(full bool) []Password {
	return fn.Map(ReadLines(20, 2, full), func(line string) Password {
		password := Password{}
		p := fn.CleanSplit(line, ":")
		h := fn.SpaceSplit(p[0])
		nums := ToIntList(h[0], "-")
		password.num1 = nums[0]
		password.num2 = nums[1]
		password.char = []rune(h[1])[0]
		password.text = p[1]
		return password
	})
}

type Password struct {
	num1 int
	num2 int
	char rune
	text string
}

func (p Password) isValid() bool {
	freq := CharFreq(p.text, nil)
	freqChar := freq[p.char]
	return p.num1 <= freqChar && freqChar <= p.num2
}

func (p Password) isValid2() bool {
	count := 0
	letters := []rune(p.text)
	for _, idx := range []int{p.num1 - 1, p.num2 - 1} {
		if letters[idx] == p.char {
			count += 1
		}
	}
	return count == 1
}
