package day03

import (
    "strings"
    "testing"
)

var example1 = strings.Split(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`, "\n")

var example2 = strings.Split("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", "\n")

var partOneAnswer = "161"
var partTwoAnswer = "48"

// func TestPartOne(t *testing.T) {
//     solution := PartOne(example1)
//     if solution != partOneAnswer {
//         t.Errorf("Day 3 Part 1: \nexpected %v got %v", partOneAnswer, solution)
//     }
// }

func TestPartTwo(t *testing.T) {
    solution := PartTwo(example2)
    if solution != partTwoAnswer {
        t.Errorf("Day 3 Part 2: \nexpected %v got %v", partTwoAnswer, solution)
    }
}
