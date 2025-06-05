package main

import (
	"fmt"
	"time"

	"github.com/roidaradal/aoc-go/aoc16"
)

func main() {
	start := time.Now()

	aoc16.Day02()

	fmt.Printf("\nTime: %v\n", time.Since(start))
}
