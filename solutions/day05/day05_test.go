package day05

import (
	"fmt"
	"strings"
	"testing"
)

var lines = strings.Split(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`, "\n")

var (
	partOneAnswer = "143"
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
