package main

import (
	"fmt"
	"time"

	"github.com/roidaradal/aoc-go/aoc18"
)

func main() {
	start := time.Now()

	aoc18.Day02()

	fmt.Printf("\nTime: %v\n", time.Since(start))
}
