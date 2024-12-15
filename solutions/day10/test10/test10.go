package main

import (
	"fmt"
	"strings"

	day10 "go-aoc-template/solutions/day10"
)

var lines = strings.Split(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`, "\n")

var (
	partOneAnswer = "36"
	partTwoAnswer = "81"
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
	runTest(1, day10.PartOne, partOneAnswer)
	runTest(2, day10.PartTwo, partTwoAnswer)
}
