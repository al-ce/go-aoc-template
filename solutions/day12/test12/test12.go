package main

import (
	"fmt"
	"strings"

	day12 "go-aoc-template/solutions/day12"
)

// Expects 236 part 2
// var lines = strings.Split(`EEEEE
// EXXXX
// EEEEE
// EXXXX
// EEEEE`, "\n")


// Expects 80 part 2
// var lines = strings.Split(`AAAA
// BBCD
// BBCC
// EEEC`, "\n")

// Expects 436 part 2
var lines = strings.Split(`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`, "\n")

var (
	partOneAnswer = "140"
	partTwoAnswer = "236"
)

type SolutionFunc func([]string) string

func runTest(part int, solution SolutionFunc, expected string) {
	fmt.Printf("Part %d: ", part)
	result := solution(lines)
	if result != expected {
		fmt.Printf("\033[31m%v\033[0m (expected \033[32m%v\033[0m)\n", result, expected)
	} else {
		fmt.Printf("\033[32m%v\033[0m\n", result)
	}
}

func main() {
	runTest(1, day12.PartOne, partOneAnswer)
	runTest(2, day12.PartTwo, partTwoAnswer)
}
