package main

import (
	"fmt"
	"strings"

	day06 "go-aoc-template/solutions/day06"
)

var lines = strings.Split(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`, "\n")

var (
	partOneAnswer = "18"
	partTwoAnswer = "9"
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
	runTest(1, day06.PartOne, partOneAnswer)
	runTest(2, day06.PartTwo, partTwoAnswer)
}
