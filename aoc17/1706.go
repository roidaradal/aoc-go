package aoc17

import (
	"fmt"
	"strings"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day06() {
	banks := data06(true)
	numBanks := len(banks)
	done := make(map[string]int)
	state := bankState(banks)
	count := 0
	done[state] = count
	for {
		index := ArgMax(banks)
		take := banks[index]
		banks[index] = 0
		for i := range take {
			idx := (index + i + 1) % numBanks
			banks[idx] += 1
		}

		count += 1
		state = bankState(banks)
		if stateCount, ok := done[state]; ok {
			fmt.Println(count)
			fmt.Println(count - stateCount)
			return
		}
		done[state] = count
	}
}

func data06(full bool) []int {
	return ToIntList(ReadLines(full)[0], " ")
}

func bankState(banks []int) string {
	state := fn.Map(banks, func(bank int) string {
		return fmt.Sprintf("%d", bank)
	})
	return strings.Join(state, ",")
}
