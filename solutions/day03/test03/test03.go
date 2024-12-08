package main

import (
	"fmt"
	"strings"

	day03 "go-aoc-template/solutions/day03"
)

var lines = strings.Split(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`, "\n")

var (
	partOneAnswer = "161"
	partTwoAnswer = "48"
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
	runTest(1, day03.PartOne, partOneAnswer)
	runTest(2, day03.PartTwo, partTwoAnswer)
}
