package aoc

import (
	"log"

	"github.com/roidaradal/fn/io"
)

func ReadLines(full bool) []string {
	path := "test.txt"
	if full {
		path = "data.txt"
	}
	lines, err := io.ReadTextLines(path)
	if err != nil {
		log.Fatal("Error:", err)
	}
	return lines
}
