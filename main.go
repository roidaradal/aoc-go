package main

import (
	"fmt"
	"time"

	"github.com/roidaradal/aoc-go/aoc15"
)

func main() {
	start := time.Now()

	aoc15.Day08()

	fmt.Printf("\nTime: %v\n", time.Since(start))
}
