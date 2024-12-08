package day04

import (
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

var partOneAnswer = "18"
var partTwoAnswer = "example answer"

func TestPartOne(t *testing.T) {
    solution := PartOne(lines)
    if solution != partOneAnswer {
        t.Errorf("Day 4 Part 1: \nexpected %v got %v", partOneAnswer, solution)
    }
}

func TestPartTwo(t *testing.T) {
    solution := PartTwo(lines)
    if solution != partTwoAnswer {
        t.Errorf("Day 4 Part 2: \nexpected %v got %v", partTwoAnswer, solution)
    }
}
