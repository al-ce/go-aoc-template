package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func PartOne(lines []string) string {
	program := strings.Join(lines[:], "")
	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := pattern.FindAllStringSubmatch(program, -1)

	sum := 0
	for _, match := range matches {
		left, _ := strconv.Atoi(match[1])
		right, _ := strconv.Atoi(match[2])
		sum += left * right
	}
	return fmt.Sprintf("%d", sum)
}

func PartTwo(lines []string) string {
	return "TODO"
}
