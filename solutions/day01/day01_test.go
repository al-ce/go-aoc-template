package day01

import (
	"strings"
	"testing"
)

var lines = strings.Split(`3   4
4   3
2   5
1   3
3   9
3   3`, "\n")

var partOneAnswer = "11"
var partTwoAnswer = "31"

func TestPartOne(t *testing.T) {
	solution := PartOne(lines)
	if solution != partOneAnswer {
		t.Errorf("Day 1 Part 1: \nexpected %v got %v", partOneAnswer, solution)
	}
}

func TestPartTwo(t *testing.T) {
	solution := PartTwo(lines)
	if solution != partTwoAnswer {
		t.Errorf("Day 1 Part 2: \nexpected %v got %v", partTwoAnswer, solution)
	}
}
