package main

import (
	"fmt"
	"strings"

	day02 "go-aoc-template/solutions/day02"
)

var lines = strings.Split(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`, "\n")

var (
	partOneAnswer = "2"
	partTwoAnswer = "4"
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
	runTest(1, day02.PartOne, partOneAnswer)
	runTest(2, day02.PartTwo, partTwoAnswer)
}
