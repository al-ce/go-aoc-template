package main

import (
	"fmt"
	"strings"

	day18 "go-aoc-template/solutions/day18"
)

var lines = strings.Split(`example input`, "\n")

var (
	partOneAnswer = "example answer"
	partTwoAnswer = "example answer"
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
	runTest(1, day18.PartOne, partOneAnswer)
	runTest(2, day18.PartTwo, partTwoAnswer)
}
