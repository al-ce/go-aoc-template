#!/usr/bin/env bash

day=$1
project_name="go-aoc-template"

# Create day{N}.go
mkdir -p "solutions/day$(printf "%02d" "${day}")" && \
cat << EOF > "solutions/day$(printf "%02d" "${day}")/day$(printf "%02d" "${day}").go"
package day$(printf "%02d" "${day}")

// import (
//     h "${project_name}/internal/helpers"
// )

func PartOne(lines []string) string {
    return "TODO"
}

func PartTwo(lines []string) string {
    return "TODO"
}
EOF

# Create day{N}_test.go
cat << EOF > "solutions/day$(printf "%02d" "${day}")/day$(printf "%02d" "${day}")_test.go"
package day$(printf "%02d" "${day}")

import (
    "strings"
    "testing"
)

var lines = strings.Split(\`example input\`, "\n")

var partOneAnswer = "example answer"
var partTwoAnswer = "example answer"

func TestPartOne(t *testing.T) {
    solution := PartOne(lines)
    if solution != partOneAnswer {
        t.Errorf("Day ${day} Part 1: \nexpected %v got %v", partOneAnswer, solution)
    }
}

func TestPartTwo(t *testing.T) {
    solution := PartTwo(lines)
    if solution != partTwoAnswer {
        t.Errorf("Day ${day} Part 2: \nexpected %v got %v", partTwoAnswer, solution)
    }
}
EOF
