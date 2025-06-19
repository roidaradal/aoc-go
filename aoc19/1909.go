package aoc19

import (
	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day09() Solution {
	// Part 1
	numbers := data09(true)
	out1 := runProgram09(numbers, 1)

	// Part 2
	numbers = data09(true)
	out2 := runProgram09(numbers, 2)

	return NewSolution(out1, out2)
}

func data09(full bool) map[int]int {
	numbers := ToIntList(ReadFirstLine(19, 9, full), ",")
	memory := make(map[int]int)
	for i, x := range numbers {
		memory[i] = x
	}
	return memory
}

func runProgram09(numbers map[int]int, start int) int {
	i, rbase := 0, 0
	output := 0
	for {
		head, tail := commandParts(numbers[i])
		cmd := fn.ParseInt(tail)
		if cmd == 99 {
			break
		}

		if cmd == 1 || cmd == 2 || cmd == 7 || cmd == 8 {
			in1, in2, out := numbers[i+1], numbers[i+2], numbers[i+3]
			m := intcodeModes(head, 3)
			m1, m2, m3 := m[0], m[1], m[2]
			a := intcodeParam2(in1, m1, rbase, numbers)
			b := intcodeParam2(in2, m2, rbase, numbers)
			c := intcodeIndex(out, m3, rbase)
			if cmd == 1 {
				numbers[c] = a + b
			} else if cmd == 2 {
				numbers[c] = a * b
			} else if cmd == 7 {
				if a < b {
					numbers[c] = 1
				} else {
					numbers[c] = 0
				}
			} else if cmd == 8 {
				if a == b {
					numbers[c] = 1
				} else {
					numbers[c] = 0
				}
			}
			i += 4
		} else if cmd == 3 {
			m := intcodeModes(head, 1)[0]
			idx := intcodeIndex(numbers[i+1], m, rbase)
			numbers[idx] = start
			i += 2
		} else if cmd == 4 {
			m := intcodeModes(head, 1)[0]
			output = intcodeParam2(numbers[i+1], m, rbase, numbers)
			i += 2
		} else if cmd == 9 {
			m := intcodeModes(head, 1)[0]
			jmp := intcodeParam2(numbers[i+1], m, rbase, numbers)
			rbase += jmp
			i += 2
		} else if cmd == 5 || cmd == 6 {
			p1, p2 := numbers[i+1], numbers[i+2]
			m := intcodeModes(head, 2)
			m1, m2 := m[0], m[1]
			isZero := intcodeParam2(p1, m1, rbase, numbers) == 0
			doJump := isZero // cmd == 6
			if cmd == 5 {
				doJump = !isZero
			}
			if doJump {
				i = intcodeParam2(p2, m2, rbase, numbers)
			} else {
				i += 3
			}
		}
	}
	return output
}
