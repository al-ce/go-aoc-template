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

mkdir -p "solutions/day$(printf "%02d" "${day}")/test$(printf "%02d" "${day}")" && \
# Create day{N}_test.go
cat << EOF > "solutions/day$(printf "%02d" "${day}")/test$(printf "%02d" "${day}")/test$(printf "%02d" "${day}").go"
package main

import (
	"fmt"
	"strings"

	day$(printf "%02d" "${day}") "go-aoc-template/solutions/day$(printf "%02d" "${day}")"
)

var lines = strings.Split(\`example input\`, "\n")

var (
	partOneAnswer = "example answer"
	partTwoAnswer = "example answer"
)

type SolutionFunc func([]string) string

func runTest(part int, solution SolutionFunc, expected string) {
	fmt.Printf("Part %d: ", part)
	result := solution(lines)
	if result != expected {
		fmt.Printf("\033[31m%v\033[0m (expected \033[32m%v\033[0m)\n", result, expected)
	} else {
		fmt.Printf("\033[32m%v\033[0m\n", result)
	}
}

func main() {
	runTest(1, day$(printf "%02d" "${day}").PartOne, partOneAnswer)
	runTest(2, day$(printf "%02d" "${day}").PartTwo, partTwoAnswer)
}
EOF
