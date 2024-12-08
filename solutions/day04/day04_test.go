package day04

import (
	"fmt"
	"strings"
	"testing"
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

func runTest(_ *testing.T, part int, solution SolutionFunc, expected string) {
	fmt.Printf("Part %d: ", part)
	result := solution(lines)
	if result != expected {
		fmt.Printf("\033[31m%v\033[0m (expected \033[32m%v\033[0m)\n", result, expected)
	} else {
		fmt.Printf("\033[32m%v\033[0m\n", result)
	}
}

func TestPartOne(t *testing.T) {
	runTest(t, 1, PartOne, partOneAnswer)
}

func TestPartTwo(t *testing.T) {
	runTest(t, 2, PartTwo, partTwoAnswer)
}
