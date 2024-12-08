package day02

import (
	"fmt"
	"strings"
	"testing"
)

var lines = strings.Split(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`, "\n")

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
