package aoc16

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day04() {
	rooms := data04(true)
	total := 0
	roomID := 0
	for _, room := range rooms {
		// Part 1
		if room.isReal() {
			total += room.id
		}
		// Part 2
		if roomID == 0 && room.decrypt() == "northpole-object-storage" {
			roomID = room.id
		}
	}
	fmt.Println(total)
	fmt.Println(roomID)
}

func data04(full bool) []Room {
	return fn.Map(ReadLines(full), func(line string) Room {
		p := fn.CleanSplit(line, "[")
		head, tail := p[0], p[1]
		h := fn.CleanSplit(head, "-")
		lastIdx := len(h) - 1
		return Room{
			checksum: strings.TrimSuffix(tail, "]"),
			name:     strings.Join(h[:lastIdx], "-"),
			id:       fn.ParseInt(h[lastIdx]),
		}
	})
}

type Room struct {
	checksum string
	name     string
	id       int
}

func (r Room) isReal() bool {
	freq := CharFreq(r.name, []rune{'-'})
	if len(freq) < 5 {
		return false
	}
	pairs := make([]CharInt, 0)
	for k, v := range freq {
		pairs = append(pairs, CharInt{Char: k, Int: v})
	}
	slices.SortFunc(pairs, func(a, b CharInt) int {
		cmp1 := cmp.Compare(b.Int, a.Int)
		if cmp1 != 0 {
			return cmp1
		}
		return cmp.Compare(a.Char, b.Char)
	})
	top5 := fn.Map(pairs[:5], func(pair CharInt) rune {
		return pair.Char
	})
	return string(top5) == r.checksum
}

func (r Room) decrypt() string {
	msg := []rune(r.name)
	for range r.id {
		for i, char := range msg {
			msg[i] = rotateLetter(char)
		}
	}
	return string(msg)
}

func rotateLetter(letter rune) rune {
	if letter == '-' {
		return '-'
	} else if letter == 'z' {
		return 'a'
	} else {
		return letter + 1
	}
}
