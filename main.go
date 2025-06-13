package main

import (
	"fmt"
	"time"

	"github.com/roidaradal/aoc-go/aoc17"
)

func main() {
	start := time.Now()

	aoc17.Day07()

	fmt.Printf("\nTime: %v\n", time.Since(start))
}
