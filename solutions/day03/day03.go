package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func runProgram(program string, pattern *regexp.Regexp) int {
	matches := pattern.FindAllStringSubmatch(program, -1)
	sum := 0
	for _, match := range matches {
		left, _ := strconv.Atoi(match[1])
		right, _ := strconv.Atoi(match[2])
		sum += left * right
	}
	return sum
}

func PartOne(lines []string) string {
	program := strings.Join(lines[:], "")
	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	sum := runProgram(program, pattern)
	return fmt.Sprintf("%d", sum)
}

func PartTwo(lines []string) string {
	program := strings.Join(lines[:], "")
	// Ignore any instructions between don't() and do()
	pattern := regexp.MustCompile(`don't\(\).*?do\(\)`)
	program = pattern.ReplaceAllString(program, "")
	return PartOne(append([]string{}, program))
}
