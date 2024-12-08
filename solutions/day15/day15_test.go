package day15

import (
	"fmt"
	"strings"
	"testing"
)

var lines = strings.Split(`example input`, "\n")

var (
	partOneAnswer = "example answer"
	partTwoAnswer = "example answer"
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
