package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/aoc-go/aoc15"
	"github.com/roidaradal/aoc-go/aoc16"
	"github.com/roidaradal/aoc-go/aoc17"
	"github.com/roidaradal/aoc-go/aoc18"
	"github.com/roidaradal/aoc-go/aoc19"
	"github.com/roidaradal/fn"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: go run . yydd (test)")
		return
	}

	yydd := args[0]
	year := fn.ParseInt(yydd[:2])
	day := fn.ParseInt(yydd[2:])
	testMode := len(args) >= 2 && args[1] == "test"

	godotenv.Load()
	sol := aoc.GetSolution(year, day)

	start := time.Now()
	ans := solve(year, day)
	if testMode {
		for i := range 2 {
			if ans[i] == sol[i] {
				fmt.Printf("OK%d: %s\n", i+1, sol[i])
			} else {
				fmt.Printf("Part%d: Exp vs Got:\n%s\n%s\n", i+1, sol[i], ans[i])
			}
		}
	} else {
		fmt.Println(ans[0])
		fmt.Println(ans[1])
	}

	fmt.Printf("\nTime: %v\n", time.Since(start))
}

func solve(year, day int) aoc.Solution {
	switch year {
	case 15:
		return solve15(day)
	case 16:
		return solve16(day)
	case 17:
		return solve17(day)
	case 18:
		return solve18(day)
	case 19:
		return solve19(day)
	}
	panic("Invalid year")
}

func solve15(day int) aoc.Solution {
	switch day {
	case 1:
		return aoc15.Day01()
	case 2:
		return aoc15.Day02()
	case 3:
		return aoc15.Day03()
	case 4:
		return aoc15.Day04()
	case 5:
		return aoc15.Day05()
	case 6:
		return aoc15.Day06()
	case 7:
		return aoc15.Day07()
	case 8:
		return aoc15.Day08()
	case 9:
		return aoc15.Day09()
	case 10:
		return aoc15.Day10()
	}
	panic("Invalid day")
}

func solve16(day int) aoc.Solution {
	switch day {
	case 1:
		return aoc16.Day01()
	case 2:
		return aoc16.Day02()
	case 3:
		return aoc16.Day03()
	case 4:
		return aoc16.Day04()
	case 5:
		return aoc16.Day05()
	case 6:
		return aoc16.Day06()
	case 7:
		return aoc16.Day07()
	case 8:
		return aoc16.Day08()
	case 9:
		return aoc16.Day09()
	case 10:
		return aoc16.Day10()
	}
	panic("Invalid day")
}

func solve17(day int) aoc.Solution {
	switch day {
	case 1:
		return aoc17.Day01()
	case 2:
		return aoc17.Day02()
	case 3:
		return aoc17.Day03()
	case 4:
		return aoc17.Day04()
	case 5:
		return aoc17.Day05()
	case 6:
		return aoc17.Day06()
	case 7:
		return aoc17.Day07()
	case 8:
		return aoc17.Day08()
	case 9:
		return aoc17.Day09()
	case 10:
		return aoc17.Day10()
	}
	panic("Invalid day")
}

func solve18(day int) aoc.Solution {
	switch day {
	case 1:
		return aoc18.Day01()
	case 2:
		return aoc18.Day02()
	case 3:
		return aoc18.Day03()
	case 4:
		return aoc18.Day04()
	case 5:
		return aoc18.Day05()
	case 6:
		return aoc18.Day06()
	case 7:
		return aoc18.Day07()
	case 8:
		return aoc18.Day08()
	case 9:
		return aoc18.Day09()
	case 10:
		return aoc18.Day10()
	}
	panic("Invalid day")
}

func solve19(day int) aoc.Solution {
	switch day {
	case 1:
		return aoc19.Day01()
	case 2:
		return aoc19.Day02()
	}
	panic("Invalid day")
}
