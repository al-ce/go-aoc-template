package day02

import (
	"fmt"

	h "go-aoc-template/internal/helpers"
)

func getDirection(report []int) int {
	i, j := 0, len(report)-1
	if i == j {
		i = 1
		j -= 1
	}
	return h.GetSignOfDiff(report[i], report[j])
}

func isSafe(this, next, direction int) bool {
	thisDirection := h.GetSignOfDiff(this, next)
	thisDiff := h.GetAbsDiff(this, next)
	return thisDirection == direction && thisDiff <= 3
}

// Check that report is safe (constant direction of change, change <= 3)
func checkReport(report []int, direction int) int {
	if direction == 0 {
		return -1
	}
	for i := 0; i <= len(report)-2; i++ {
		this, next := report[i], report[i+1]
		if !isSafe(this, next, direction) {
			return -1
		}
	}
	return 0
}

func PartOne(lines []string) string {
	safeCount := len(lines)
	for _, reportStr := range lines {
		report := h.ParseIntString(reportStr, " ")
		direction := getDirection(report)
		safeCount += checkReport(report, direction)
	}
	return fmt.Sprintf("%d", safeCount)
}

func PartTwo(lines []string) string {
	safeCount := 0
	for _, reportStr := range lines {
		report := h.ParseIntString(reportStr, " ")
		direction := getDirection(report)

		// Check all subsets, "removing" one at a time
		// e.g.: 9 7 6 2 1
		// i = 0: _ 7 6 2 1; i = 1: 9 _ 6 2 1 etc.
		// Break as soon as we find a safe subset

		for i := 1; i <= len(report); i++ {
			start, rest := report[:i-1], report[i:]
			subset := []int{}
			subset = append(subset, start...)
			subset = append(subset, rest...)

			if checkReport(subset, direction) == 0 {
				safeCount += 1
				break
			}
		}
	}
	return fmt.Sprintf("%d", safeCount)
}
