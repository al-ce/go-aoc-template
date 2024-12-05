package day20

import (
    "strings"
    "testing"
)

var lines = strings.Split(`example input`, "\n")

var partOneAnswer = "example answer"
var partTwoAnswer = "example answer"

func TestPartOne(t *testing.T) {
    solution := PartOne(lines)
    if solution != partOneAnswer {
        t.Errorf("Day 20 Part 1: \nexpected %v got %v", partOneAnswer, solution)
    }
}

func TestPartTwo(t *testing.T) {
    solution := PartTwo(lines)
    if solution != partTwoAnswer {
        t.Errorf("Day 20 Part 2: \nexpected %v got %v", partTwoAnswer, solution)
    }
}
