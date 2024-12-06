package day02

import (
    "strings"
    "testing"
)

var lines = strings.Split(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`, "\n")

var partOneAnswer = "2"
var partTwoAnswer = "4"

func TestPartOne(t *testing.T) {
    solution := PartOne(lines)
    if solution != partOneAnswer {
        t.Errorf("Day 2 Part 1: \nexpected %v got %v", partOneAnswer, solution)
    }
}

func TestPartTwo(t *testing.T) {
    solution := PartTwo(lines)
    if solution != partTwoAnswer {
        t.Errorf("Day 2 Part 2: \nexpected %v got %v", partTwoAnswer, solution)
    }
}
